package helper

import (
	"github.com/CityBear3/satellite/internal/domain/primitive/authentication"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSecret() (*authentication.HashedSecrets, error) {
	secret := "test-1234"

	value, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	newSecret, err := authentication.NewHashedSecrets(string(value))
	if err != nil {
		return nil, err
	}

	return newSecret, nil
}
