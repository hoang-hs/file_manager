package services

import (
	"file_manager/src/adapter/database/models"
	"file_manager/src/api/request"
	"file_manager/src/core/errors"
)

type RegisterService interface {
	SignUp(registerPack *request.RegisterPackage) (*models.User, errors.Error)
}
