package configs

import (
	"github.com/joho/godotenv"
	"log"
)

func LoadEnv(pathConfig string) {
	if err := godotenv.Load(pathConfig); err != nil {
		log.Fatalf("Can not load env")
	}
}
