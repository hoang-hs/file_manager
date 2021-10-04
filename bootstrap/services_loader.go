package bootstrap

import (
	"database/sql"
	"file_manager/api/controllers"
	"file_manager/configs"
	"file_manager/internal/repositories"
	"file_manager/internal/services"
	"time"
)

func LoadServices(dbConnection *sql.DB) *controllers.ApplicationContext {
	userRepository := repositories.NewUserRepository(dbConnection)
	authService := services.NewAuthService(*userRepository, time.Duration(configs.Get().ExpiredDuration))
	registerService := services.NewRegisterService(*userRepository)
	return &controllers.ApplicationContext{
		DB:              dbConnection,
		FileService:     services.NewFileService(),
		AuthService:     authService,
		RegisterService: registerService,
	}
}
