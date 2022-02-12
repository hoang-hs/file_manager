package registry

import (
	"file_manager/src/api/services"
	"file_manager/src/core/usecases"
)

func NewRegisterService(registerServiceImpl *usecases.RegisterService) services.RegisterService {
	return registerServiceImpl
}
