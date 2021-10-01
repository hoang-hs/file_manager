package bootstrap

import (
	"database/sql"
	"file_manager/configs"
	"file_manager/internal/repositories"
	"file_manager/internal/services"
	"time"
)

var AuthService *services.AuthService
var RegisterService *services.RegisterService
var FileService *services.FileService
var TokenService *services.TokenService

func LoadServices(dbConnection *sql.DB) {
	userRepository := repositories.NewUserRepository(dbConnection)
	authService := services.NewAuthService(*userRepository, time.Duration(configs.Get().ExpiredDuration))
	registerService := services.NewRegisterService(*userRepository)
	publish(authService, registerService)
}

func publish(auth *services.AuthService, registerService *services.RegisterService) {
	AuthService = auth
	RegisterService = registerService
}
