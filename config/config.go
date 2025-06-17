package config

import (
	"os"
	"strconv"
)

func GetConfig(key string) string {
	return os.Getenv(key)
}

func GetIntConfig(key string) int {
	value := os.Getenv(key)
	if value == "" {
		return 0
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return intValue
}

func GetConfigOrDefault(key, defaultValue string) string {
	res := os.Getenv(key)
	if res == "" {
		return defaultValue
	}
	return res
}

func GetIntConfigOrDefault(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}