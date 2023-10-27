package convertors

import (
	"database/sql"
	"testing"

	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestConvertError(t *testing.T) {
	logger := zap.NewExample()

	tests := []struct {
		name     string
		inputErr error
		expected error
	}{
		{
			name:     "TestNotFound",
			inputErr: apperrs.NotFoundArchiveError,
			expected: status.Error(codes.NotFound, apperrs.NotFoundArchiveError.Msg),
		},
		{
			name:     "TestBadRequest",
			inputErr: apperrs.InvalidFileError,
			expected: status.Error(codes.InvalidArgument, apperrs.InvalidFileError.Msg),
		},
		{
			name:     "TestUnauthenticatedError",
			inputErr: apperrs.UnauthenticatedError,
			expected: status.Error(codes.Unauthenticated, apperrs.UnauthenticatedError.Msg),
		},
		{
			name:     "TestUnknownError",
			inputErr: sql.ErrTxDone,
			expected: status.Error(codes.Internal, apperrs.UnexpectedError.Msg),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if result := ConvertError(logger, tt.inputErr); result.Error() != tt.expected.Error() {
				t.Errorf("ConvertError() error = %v, expected %v", result, tt.expected)
			}
		})
	}
}
