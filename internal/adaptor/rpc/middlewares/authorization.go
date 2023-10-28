package middlewares

import (
	"context"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/repository"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AuthorizationInterceptor struct {
	logger           *zap.Logger
	deviceRepository repository.IDeviceRepository
	clientRepository repository.IClientRepository
}

func NewAuthorizationInterceptor(
	logger *zap.Logger,
	deviceRepository repository.IDeviceRepository,
	clientRepository repository.IClientRepository,
) AuthorizationInterceptor {
	return AuthorizationInterceptor{
		logger:           logger,
		deviceRepository: deviceRepository,
		clientRepository: clientRepository,
	}
}

func (i AuthorizationInterceptor) Authorization() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		ctx, err = i.authorization(ctx, info.FullMethod)
		if err != nil {
			return nil, convertors.ConvertError(i.logger, err)
		}

		return handler(ctx, req)
	}
}

func (i AuthorizationInterceptor) AuthorizationStream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx, err := i.authorization(ss.Context(), info.FullMethod)
		if err != nil {
			return convertors.ConvertError(i.logger, err)
		}

		wrappedServerStream := middleware.WrapServerStream(ss)
		wrappedServerStream.WrappedContext = ctx

		return handler(srv, wrappedServerStream)
	}
}

func (i AuthorizationInterceptor) authorization(ctx context.Context, method string) (context.Context, error) {
	switch method {
	// device
	case "/satellite.archive.v1.ArchiveService/CreateArchive", "/satellite.event.v1.ArchiveEventService/ReceiveEvent":
		id, err := getID(ctx)
		if err != nil {
			return nil, apperrs.UnauthenticatedError
		}

		device, err := i.deviceRepository.GetDevice(ctx, id)
		if err != nil {
			return nil, apperrs.ForbiddenError
		}

		return context.WithValue(ctx, "device", device), nil
	// client
	case "/satellite.archive.v1.ArchiveService/GetArchive", "/satellite.event.v1.ArchiveEventService/PublishEvent":
		id, err := getID(ctx)
		if err != nil {
			return nil, apperrs.UnauthenticatedError
		}

		client, err := i.clientRepository.GetClient(ctx, id)
		if err != nil {
			return nil, apperrs.ForbiddenError
		}

		return context.WithValue(ctx, "client", client), nil
	default:
		return nil, convertors.ConvertError(i.logger, apperrs.UnexpectedError)
	}
}

func getID(ctx context.Context) (primitive.ID, error) {
	id, ok := ctx.Value("id").(string)
	if !ok {
		return primitive.ID{}, apperrs.UnexpectedError
	}

	parseID, err := primitive.ParseID(id)
	if err != nil {
		return primitive.ID{}, err
	}

	return parseID, nil
}
