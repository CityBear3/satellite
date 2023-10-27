package convertors

import (
	"errors"

	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertError(logger *zap.Logger, err error) error {
	var appError *apperrs.Error
	var code codes.Code
	if !errors.As(err, &appError) {
		goto Internal
	}

	switch appError.Code {
	case apperrs.NotFound:
		code = codes.NotFound
	case apperrs.BadRequest:
		code = codes.InvalidArgument
	case apperrs.Unauthenticated:
		code = codes.Unauthenticated
	default:
		goto Internal
	}
	return status.Error(code, appError.Msg)

Internal:
	logger.Error(err.Error())
	return status.Error(codes.Internal, apperrs.UnexpectedError.Msg)
}
