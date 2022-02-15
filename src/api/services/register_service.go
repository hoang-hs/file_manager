package services

import (
	"file_manager/src/api/request"
	"file_manager/src/core/entities"
	"file_manager/src/core/errors"
)

type RegisterService interface {
	SignUp(registerPack *request.RegisterRequest) (*entities.User, errors.Error)
}
