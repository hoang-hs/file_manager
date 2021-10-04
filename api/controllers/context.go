package controllers

import (
	"database/sql"
	"file_manager/internal/services"
)

type ApplicationContext struct {
	DB              *sql.DB
	FileService     *services.FileService
	AuthService     *services.AuthService
	RegisterService *services.RegisterService
}
