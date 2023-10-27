package primitive

import "github.com/CityBear3/satellite/internal/pkg/apperrs"

type DeviceName string

func NewDeviceName(value string) (DeviceName, error) {
	if len(value) > 200 {
		return "", apperrs.UnexpectedError
	}

	return DeviceName(value), nil
}

func (d DeviceName) String() string {
	return string(d)
}
