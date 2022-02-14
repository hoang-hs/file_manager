package registry

import (
	"file_manager/src/api/services"
	"file_manager/src/core/usecases"
)

func NewAuthService(authServiceImpl *usecases.AuthUseCase) services.AuthService {
	return authServiceImpl
}
