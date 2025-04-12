package env

import (
	"os"
	"strconv"
)

func GetString(key string, fallback string) string {
	val, err := os.LookupEnv(key)
	if err {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, err := os.LookupEnv(key)
	if err {
		return fallback
	}

	conv, newErr := strconv.Atoi(val)
	if newErr != nil {
		return fallback
	}

	return conv
}
