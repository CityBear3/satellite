package archive

import (
	"strings"

	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type ContentType struct {
	value string
}

func NewContentType(value string) (ContentType, error) {
	switch value {
	case "image/jpeg", "image/png":
		return ContentType{value: value}, nil
	default:
		return ContentType{}, apperrs.InvalidFileError
	}
}

func (c ContentType) Value() string {
	return c.value
}

func (c ContentType) GetExt() string {
	return strings.TrimPrefix(c.value, "image/")
}
