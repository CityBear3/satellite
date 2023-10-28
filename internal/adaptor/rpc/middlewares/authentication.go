package middlewares

import (
	"context"
	"time"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
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

func (i AuthenticationInterceptor) Authentication() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		ctx, err = i.authentication(ctx, info.FullMethod)
		if err != nil {
			return nil, convertors.ConvertError(i.logger, err)
		}

		return handler(ctx, req)
	}
}

func (i AuthenticationInterceptor) AuthenticationStream() grpc.StreamServerInterceptor {
	return func(srv any, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := i.authentication(ss.Context(), info.FullMethod)
		if err != nil {
			return convertors.ConvertError(i.logger, err)
		}

		wrappedServerStream := middleware.WrapServerStream(ss)
		wrappedServerStream.WrappedContext = ctx

		return handler(srv, wrappedServerStream)
	}
}

// TODO: some method skip this.
func (i AuthenticationInterceptor) authentication(ctx context.Context, method string) (context.Context, error) {
	tokenString, err := grpcAuth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, apperrs.UnauthenticatedError
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperrs.UnauthenticatedError
		}

		return []byte(i.HMACSecret), nil
	})

	if err != nil {
		return nil, apperrs.UnauthenticatedError
	}

	var id string
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, ok = claims["sub"].(string)
		if !ok {
			return nil, apperrs.UnexpectedError
		}

		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return nil, apperrs.UnauthenticatedError
		}

	} else {
		return nil, apperrs.UnauthenticatedError
	}

	ctx = context.WithValue(ctx, "id", id)
	return ctx, nil
}
