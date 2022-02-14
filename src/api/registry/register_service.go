package registry

import (
	"file_manager/src/api/services"
	"file_manager/src/core/usecases"
)

func NewRegisterService(registerServiceImpl *usecases.RegisterUseCase) services.RegisterService {
	return registerServiceImpl
}
