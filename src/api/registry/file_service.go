package registry

import (
	"file_manager/src/api/services"
	"file_manager/src/core/usecases"
)

func NewFileService(fileServiceImpl *usecases.FileUseCase) services.FileService {
	return fileServiceImpl
}
