package driver

import (
	"fmt"
	"github.com/CityBear3/satellite/adaptor/rpc"
	pbImage "github.com/CityBear3/satellite/pb/image/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	port      int
	isDevelop bool
}

func NewServer(port int, isDevelop bool) *Server {
	return &Server{
		port:      port,
		isDevelop: isDevelop,
	}
}

func (s *Server) Serve() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	if s.isDevelop {
		reflection.Register(server)
	}

	imageService := rpc.NewImageService()
	pbImage.RegisterImageServiceServer(server, imageService)

	log.Println("Start gRPC server")
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
