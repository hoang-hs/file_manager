package configs

import (
	"os"
	"strconv"
	"time"
)

func getString(key string) string {
	return os.Getenv(key)
}

func getInt(key string) int {
	value := os.Getenv(key)
	valueInt, _ := strconv.Atoi(value)
	return valueInt
}
func getStringD(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getTimeDurationWithDefault(key string, defaultValue int) time.Duration {
	strValue := os.Getenv(key)
	if len(strValue) == 0 {
		return time.Duration(defaultValue) * time.Minute
	}
	numericValue, _ := strconv.Atoi(strValue)
	return time.Duration(numericValue) * time.Minute
}
