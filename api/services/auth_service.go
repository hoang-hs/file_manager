package services

import (
	"file_manager/internal/entities"
	"file_manager/internal/errors"
	"file_manager/internal/services"
)

type IAuthService interface {
	Authenticate(authPackage entities.AuthPackage) (*entities.Authentication, errors.Error)
}

func InitAuthService(authServiceImpl *services.AuthService) IAuthService {
	return authServiceImpl
}
