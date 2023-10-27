package middlewares

import (
	"context"
	"time"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/domain/primitive"
	"github.com/CityBear3/satellite/internal/domain/repository"
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/golang-jwt/jwt"
	middleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type AuthenticationInterceptor struct {
	logger           *zap.Logger
	HMACSecret       string
	deviceRepository repository.IDeviceRepository
}

func NewAuthenticationInterceptor(logger *zap.Logger, hmacSecret string, deviceRepository repository.IDeviceRepository) AuthenticationInterceptor {
	return AuthenticationInterceptor{
		logger:           logger,
		HMACSecret:       hmacSecret,
		deviceRepository: deviceRepository,
	}
}

func (i AuthenticationInterceptor) AuthFunc(ctx context.Context) (context.Context, error) {
	tokenString, err := grpcAuth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, convertors.ConvertError(i.logger, apperrs.UnauthenticatedError)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, convertors.ConvertError(i.logger, apperrs.UnauthenticatedError)
		}

		return []byte(i.HMACSecret), nil
	})

	if err != nil {
		return nil, convertors.ConvertError(i.logger, apperrs.UnauthenticatedError)
	}

	var id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok = claims["sub"].(string)
		if !ok {
			return nil, convertors.ConvertError(i.logger, apperrs.UnexpectedError)
		}

		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, convertors.ConvertError(i.logger, apperrs.UnauthenticatedError)
		}

	} else {
		return nil, convertors.ConvertError(i.logger, apperrs.UnauthenticatedError)
	}

	ctx = context.WithValue(ctx, "id", id)
	return ctx, nil
}

func (i AuthenticationInterceptor) Authorization() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		switch info.FullMethod {
		default:
			return nil, convertors.ConvertError(i.logger, apperrs.UnexpectedError)
		}
		//return handler(ctx, req)
	}
}

func (i AuthenticationInterceptor) AuthorizationStream() grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		ctx := ss.Context()
		switch info.FullMethod {
		case "/satellite.archive.v1.ArchiveService/CreateArchive", "/satellite.event.v1.EventService/ReceiveEvent":
			id, err := getID(ctx)
			if err != nil {
				return convertors.ConvertError(i.logger, err)
			}

			device, err := i.deviceRepository.GetDevice(ctx, id)
			if err != nil {
				return convertors.ConvertError(i.logger, err)
			}

			ctx = context.WithValue(ctx, "device", device)
			wrappedServerStream := middleware.WrapServerStream(ss)
			wrappedServerStream.WrappedContext = ctx

			return handler(srv, wrappedServerStream)
		case "/satellite.archive.v1.ArchiveService/GetArchive":

		default:
			return convertors.ConvertError(i.logger, apperrs.UnexpectedError)
		}
		return handler(srv, ss)
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
