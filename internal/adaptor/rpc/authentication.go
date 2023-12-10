package rpc

import (
	"context"
	"time"

	"github.com/CityBear3/satellite/internal/adaptor/rpc/convertors"
	"github.com/CityBear3/satellite/internal/usecase"
	"github.com/CityBear3/satellite/pb/authentication/v1"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

type AuthenticationRPCService struct {
	logger                   *zap.Logger
	authenticationInteractor usecase.AuthenticationUseCase
	hmacSecret               string
	authPb.UnimplementedAuthenticationServiceServer
}

func NewAuthenticationRPCService(
	logger *zap.Logger,
	authenticationInteractor usecase.AuthenticationUseCase,
	hmacSecret string,
) AuthenticationRPCService {
	return AuthenticationRPCService{
		logger:                   logger,
		authenticationInteractor: authenticationInteractor,
		hmacSecret:               hmacSecret,
	}
}

func (s AuthenticationRPCService) handleError(err error) error {
	return convertors.ConvertError(s.logger, err)
}

func (s AuthenticationRPCService) AuthenticateClient(ctx context.Context, request *authPb.AuthenticateRequest) (*authPb.AuthenticateResponse, error) {
	authenticateRequest, err := convertors.ConvertToAuthenticationRequest(request)
	if err != nil {
		return nil, s.handleError(err)
	}

	if err = s.authenticationInteractor.AuthenticateClient(ctx, authenticateRequest); err != nil {
		return nil, s.handleError(err)
	}

	token, err := createToken(request.Id, s.hmacSecret)
	if err != nil {
		return nil, s.handleError(err)
	}

	return &authPb.AuthenticateResponse{
		Token: token,
	}, nil
}

func (s AuthenticationRPCService) AuthenticateDevice(ctx context.Context, request *authPb.AuthenticateRequest) (*authPb.AuthenticateResponse, error) {
	authenticateRequest, err := convertors.ConvertToAuthenticationRequest(request)
	if err != nil {
		return nil, s.handleError(err)
	}

	if err = s.authenticationInteractor.AuthenticateDevice(ctx, authenticateRequest); err != nil {
		return nil, s.handleError(err)
	}

	token, err := createToken(request.Id, s.hmacSecret)
	if err != nil {
		return nil, s.handleError(err)
	}

	return &authPb.AuthenticateResponse{
		Token: token,
	}, nil
}

func createToken(id string, hmacSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24 * 180).Unix(),
	})

	signedToken, err := token.SignedString([]byte(hmacSecret))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
