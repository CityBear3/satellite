package authentication

import (
	"sync"

	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type Secret interface {
	checkReadOnce() error
	Value() (string, error)
}

type HashedSecret struct {
	value    string
	once     sync.Once
	consumed bool
}

func NewHashedSecret(value string) (*HashedSecret, error) {
	if len(value) != 60 || ("$2a$10" != value[0:6]) {
		return nil, apperrs.UnexpectedError
	}

	return &HashedSecret{
		value: value,
	}, nil
}

func (s *HashedSecret) Value() (string, error) {
	if err := s.checkReadOnce(); err != nil {
		return "", err
	}

	return s.value, nil
}

func (s *HashedSecret) checkReadOnce() error {
	if s.consumed {
		return apperrs.UnexpectedError
	}

	s.once.Do(func() {
		s.consumed = true
	})

	return nil
}

type RawSecret struct {
	value    string
	once     sync.Once
	consumed bool
}

func NewRawSecret(value string) *RawSecret {
	return &RawSecret{
		value: value,
	}
}

func (s *RawSecret) Value() (string, error) {
	if err := s.checkReadOnce(); err != nil {
		return "", err
	}

	return s.value, nil
}

func (s *RawSecret) checkReadOnce() error {
	if s.consumed {
		return apperrs.UnexpectedError
	}

	s.once.Do(func() {
		s.consumed = true
	})

	return nil
}
