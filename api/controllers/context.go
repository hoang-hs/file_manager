package controllers

import (
	"file_manager/internal/services"
)

type ApplicationContext struct {
	FileService     *services.FileService
	AuthService     *services.AuthService
	RegisterService *services.RegisterService
}

func NewAppLiCationContext(fileService *services.FileService,
	authService *services.AuthService,
	registerService *services.RegisterService,
) *ApplicationContext {
	return &ApplicationContext{
		FileService:     fileService,
		AuthService:     authService,
		RegisterService: registerService,
	}
}
