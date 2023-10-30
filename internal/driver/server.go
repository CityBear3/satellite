package driver

import (
	"database/sql"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/CityBear3/satellite/internal/adaptor/event/rabbitmq"
	"github.com/CityBear3/satellite/internal/adaptor/repository/mysql"
	"github.com/CityBear3/satellite/internal/adaptor/rpc"
	"github.com/CityBear3/satellite/internal/adaptor/rpc/middlewares"
	"github.com/CityBear3/satellite/internal/usecase/interactor"
	"github.com/CityBear3/satellite/pb/archive/v1"
	"github.com/CityBear3/satellite/pb/authentication/v1"
	"github.com/CityBear3/satellite/pb/event/v1"
	grpcLog "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	cfg Config
}

func NewServer(cfg Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

// Serve start server
func (s *Server) Serve() error {
	serverCfg := s.cfg.ServerConfig
	dbCfg := s.cfg.DBConfig

	// server
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverCfg.Host, serverCfg.Port))
	if err != nil {
		return err
	}

	// zap
	logger := zap.NewExample()
	if err != nil {
		return err
	}
	logOption := []grpcLog.Option{
		grpcLog.WithLogOnEvents(grpcLog.StartCall, grpcLog.FinishCall),
	}

	// db
	db, err := CreateDB(dbCfg)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	// repository
	txManager := mysql.NewTxManger(db)
	archiveRepository := mysql.NewArchiveRepository(db)
	eventRepository := mysql.NewEventRepository(db)
	deviceRepository := mysql.NewDeviceRepository(db)
	clientRepository := mysql.NewClientRepository(db)

	// event handler
	rqConf := s.cfg.RabbitMQConfig
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", rqConf.User, rqConf.Password, rqConf.Host, rqConf.Port))
	if err != nil {
		return err
	}

	eventHandler := rabbitmq.NewEventHandler(logger, conn)

	// interactor
	archiveInteractor := interactor.NewArchiveInteractor(archiveRepository, eventRepository, txManager)
	eventInteractor := interactor.NewEventInteractor(eventRepository, eventHandler, txManager)
	authenticationInteractor := interactor.NewAuthenticationInteractor(clientRepository, deviceRepository)

	// rpc service
	archiveRPCService := rpc.NewArchiveRPCService(logger, archiveInteractor)
	eventRPCService := rpc.NewEventRPCService(logger, eventInteractor)
	authenticationRPCService := rpc.NewAuthenticationRPCService(logger, authenticationInteractor, s.cfg.AuthConfig.HMACSecret)

	// interceptor
	authenticationInterceptor := middlewares.NewAuthenticationInterceptor(logger, s.cfg.AuthConfig.HMACSecret, deviceRepository)
	authorizationInterceptor := middlewares.NewAuthorizationInterceptor(logger, deviceRepository, clientRepository)
	loggingInterceptor := middlewares.NewLoggingInterceptor(logger)

	serverParameters := keepalive.ServerParameters{
		Time:    10 * time.Second,
		Timeout: 5 * time.Second,
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcLog.UnaryServerInterceptor(loggingInterceptor.Logger(), logOption...),
			authenticationInterceptor.Authentication(),
			authorizationInterceptor.Authorization(),
		),
		grpc.ChainStreamInterceptor(
			grpcLog.StreamServerInterceptor(loggingInterceptor.Logger(), logOption...),
			authenticationInterceptor.AuthenticationStream(),
			authorizationInterceptor.AuthorizationStream(),
		),
		grpc.KeepaliveParams(serverParameters),
	)

	archive.RegisterArchiveServiceServer(server, archiveRPCService)
	event.RegisterArchiveEventServiceServer(server, eventRPCService)
	authentication.RegisterAuthenticationServiceServer(server, authenticationRPCService)

	if serverCfg.IsDevelop {
		reflection.Register(server)
	}

	logger.Info("Start gRPC server")
	if err = server.Serve(listener); err != nil {
		return err
	}

	waitSIGINT()
	server.GracefulStop()

	return nil
}

func waitSIGINT() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	signal.Stop(quit)
	close(quit)
}
