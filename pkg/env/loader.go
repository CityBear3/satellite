package env

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// ValueNotFoundError error message for value not found
const ValueNotFoundError = "%s is not found. this environment value must be set."

// GetStrEnv get value
func GetStrEnv(key string, defaultValue string) string {
	v, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}

	return v
}

func GetRequiredStrEnv(key string) (string, error) {
	v, exist := os.LookupEnv(key)
	if !exist {
		return "", errors.New(fmt.Sprintf(ValueNotFoundError, key))
	}

	return v, nil
}

// GetIntEnv get int value
func GetIntEnv(key string, defaultValue int) (int, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue, nil
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func GetRequiredIntEnv(key string) (int, error) {
	v, exist := os.LookupEnv(key)
	if !exist {
		return 0, errors.New(fmt.Sprintf(ValueNotFoundError, key))
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func GetBoolEnv(key string, defaultValue bool) (bool, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue, nil
	}

	boolValue, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}

	return boolValue, nil
}

func GetRequiredBoolEnv(key string) (bool, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		return false, errors.New(fmt.Sprintf(ValueNotFoundError, key))
	}

	boolValue, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}

	return boolValue, nil
}
