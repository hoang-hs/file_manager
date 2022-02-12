package registry

import (
	"file_manager/src/api/services"
	"file_manager/src/core/usecases"
)

func NewAuthService(authServiceImpl *usecases.AuthService) services.AuthService {
	return authServiceImpl
}
