package services

import (
	"file_manager/src/core/entities"
	"file_manager/src/core/errors"
	"github.com/gin-gonic/gin"
)

type FileService interface {
	GetFile(path string) (*entities.Files, errors.Error)
	UploadFile(c *gin.Context, path string) errors.Error
	DeleteFile(path string) errors.Error
}
