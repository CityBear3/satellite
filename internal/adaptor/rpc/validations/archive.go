package validations

import (
	"github.com/CityBear3/satellite/internal/pkg/apperrs"
	"github.com/CityBear3/satellite/pb/archive/v1"
)

// max data size: 5MB
const maxDataSize = 5e6

// ValidateCreateArchive validate upload image stream
func ValidateCreateArchive(meta *archive.Meta, data []byte) error {
	if meta == nil {
		return apperrs.InvalidMetaInfoError
	}

	if len(data) == 0 || len(data) >= maxDataSize {
		return apperrs.InvalidFileSizeError
	}

	return nil
}
