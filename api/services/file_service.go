package services

import (
	"file_manager/internal/entities"
	"file_manager/internal/errors"
	"file_manager/internal/services"
	"github.com/gin-gonic/gin"
)

type IFileService interface {
	GetFile(path string) (*entities.Files, errors.Error)
	UploadFile(c *gin.Context, path string) errors.Error
	DeleteFile(path string) errors.Error
}

func InitFileService(fileServiceImpl *services.FileService) IFileService {
	return fileServiceImpl
}
