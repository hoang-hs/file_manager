package helpers

import (
	configs2 "file_manager/src/configs"
)

func GetPrivateKey() ([]byte, error) {
	privateKeyBytes := []byte(configs2.Get().SecretKey)
	return privateKeyBytes, nil
}
