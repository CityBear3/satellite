package client

import "github.com/CityBear3/satellite/internal/pkg/apperrs"

type Name string

func NewName(value string) (Name, error) {
	if len(value) > 200 {
		return "", apperrs.UnexpectedError
	}

	return Name(value), nil
}

func (c Name) String() string {
	return string(c)
}
