package configs

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv(pathConfig string) {
	myDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(myDir)
	if err := godotenv.Load(pathConfig); err != nil {
		log.Fatalf("Can not load env")
	}
}
