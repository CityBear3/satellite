package validations

import (
	"github.com/CityBear3/satellite/pb/image/v1"
	"github.com/CityBear3/satellite/pkg/apperrs"
	"net/http"
)

// ValidateUploadImageStream validate upload image stream
func ValidateUploadImageStream(meta *image.Meta, data []byte) error {
	if meta == nil {
		return apperrs.NewError(apperrs.BadRequest, apperrs.InvalidMetaInfoMsg)
	}
	if len(data) == 0 {
		return apperrs.NewError(apperrs.BadRequest, apperrs.InvalidDataMsg)
	}
	contentType := http.DetectContentType(data)
	if !isAllowedContentType(contentType) {
		return apperrs.NewError(apperrs.BadRequest, apperrs.InvalidFileExtMsg)
	}

	return nil
}

// check file extension is allowed
func isAllowedContentType(contentType string) bool {
	allowedContentTypes := []string{"image/jpeg", "image/png"}
	for _, a := range allowedContentTypes {
		if contentType == a {
			return true
		}
	}

	return false
}
