package helpers

import (
	"file_manager/configs"
)

func GetPrivateKey() ([]byte, error) {
	privateKeyBytes := []byte(configs.Get().SecretKey)
	return privateKeyBytes, nil
}
