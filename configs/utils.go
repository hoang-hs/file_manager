package configs

import (
	"log"
	"os"
	"strconv"
)

func getString(key string) string {
	return os.Getenv(key)
}
func GetInt(key string, defaultValue int) int {
	strValue := os.Getenv(key)
	if strValue == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue
	}
	return value
}

func getStringD(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getIntD(key string, defaultValue int) int {
	value := os.Getenv(key)
	valueInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("format env var wrong")
	}
	if value == "" {
		return defaultValue
	}
	return valueInt
}
