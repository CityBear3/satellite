package primitive

import "github.com/CityBear3/satellite/internal/pkg/apperrs"

type ClientName string

func NewClientName(value string) (ClientName, error) {
	if len(value) > 200 {
		return "", apperrs.UnexpectedError
	}

	return ClientName(value), nil
}

func (c ClientName) String() string {
	return string(c)
}
