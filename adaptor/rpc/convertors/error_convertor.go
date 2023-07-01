package convertors

import (
	"github.com/CityBear3/satellite/pkg/apperrs"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ConvertError(err error) error {
	switch err := err.(type) {
	case *apperrs.Error:
		return convertAppErrToRPCError(err)
	default:
		return status.Error(codes.Unknown, "Unknown error occurred. Please call administrator if this error is not fixed.")
	}
}

func convertAppErrToRPCError(err *apperrs.Error) error {
	switch err.Code {
	case apperrs.NotFound:
		return status.Error(codes.NotFound, err.Msg)
	case apperrs.BadRequest:
		return status.Error(codes.InvalidArgument, err.Msg)
	default:
		return status.Error(codes.Internal, "Unexpected error occurred. Please call administrator if this error is not fixed.")
	}
}
