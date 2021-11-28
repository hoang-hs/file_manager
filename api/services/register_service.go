package services

import (
	"file_manager/internal/entities"
	"file_manager/internal/errors"
	"file_manager/internal/models"
	"file_manager/internal/services"
)

type IRegisterService interface {
	SignUp(registerPack *entities.RegisterPackage) (*models.User, errors.Error)
}

func InitRegisterService(registerServiceImpl *services.RegisterService) IRegisterService {
	return registerServiceImpl
}
