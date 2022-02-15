package services

import (
	"file_manager/src/api/request"
	"file_manager/src/core/entities"
	"file_manager/src/core/errors"
)

type AuthService interface {
	Authenticate(authPackage *request.AuthRequest) (*entities.Authentication, errors.Error)
}
