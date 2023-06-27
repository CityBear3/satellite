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
func GetStrEnv(key string, defaultValue string, required bool) (string, error) {
	v, exist := os.LookupEnv(key)
	if !exist {
		if required {
			return "", errors.New(fmt.Sprintf(ValueNotFoundError, key))
		}
		return defaultValue, nil
	}

	return v, nil
}

// GetIntEnv get int value
func GetIntEnv(key string, defaultValue int, required bool) (int, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if required {
			return 0, errors.New(fmt.Sprintf(ValueNotFoundError, key))
		}
		return defaultValue, nil
	}

	n, err := strconv.Atoi(v)
	if err != nil {
		return 0, err
	}

	return n, nil
}

func GetBoolEnv(key string, defaultValue bool, required bool) (bool, error) {
	v, exists := os.LookupEnv(key)
	if !exists {
		if required {
			return false, errors.New(fmt.Sprintf(ValueNotFoundError, key))
		}
		return defaultValue, nil
	}

	boolValue, err := strconv.ParseBool(v)
	if err != nil {
		return false, err
	}

	return boolValue, nil
}
