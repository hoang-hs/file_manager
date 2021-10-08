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

func NewAppLiCationContext(db *sql.DB, fileService *services.FileService,
	authService *services.AuthService, registerService *services.RegisterService) *ApplicationContext {
	return &ApplicationContext{
		DB:              db,
		FileService:     fileService,
		AuthService:     authService,
		RegisterService: registerService,
	}
}
