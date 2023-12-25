package authentication

import (
	"sync"

	"github.com/CityBear3/satellite/internal/pkg/apperrs"
)

type Secrets interface {
	checkReadOnce() error
	Value() (string, error)
}

type HashedSecrets struct {
	value    string
	once     sync.Once
	consumed bool
}

func NewHashedSecrets(value string) (*HashedSecrets, error) {
	if len(value) != 60 || ("$2a$10" != value[0:6]) {
		return nil, apperrs.UnexpectedError
	}

	return &HashedSecrets{
		value: value,
	}, nil
}

func (s *HashedSecrets) Value() (string, error) {
	if err := s.checkReadOnce(); err != nil {
		return "", err
	}

	return s.value, nil
}

func (s *HashedSecrets) checkReadOnce() error {
	if s.consumed {
		return apperrs.UnexpectedError
	}

	s.once.Do(func() {
		s.consumed = true
	})

	return nil
}

type RawSecrets struct {
	value    string
	once     sync.Once
	consumed bool
}

func NewRawSecrets(value string) *RawSecrets {
	return &RawSecrets{
		value: value,
	}
}

func (s *RawSecrets) Value() (string, error) {
	if err := s.checkReadOnce(); err != nil {
		return "", err
	}

	return s.value, nil
}

func (s *RawSecrets) checkReadOnce() error {
	if s.consumed {
		return apperrs.UnexpectedError
	}

	s.once.Do(func() {
		s.consumed = true
	})

	return nil
}
