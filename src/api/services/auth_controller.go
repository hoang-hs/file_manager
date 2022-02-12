package services

import (
	request2 "file_manager/src/api/request"
	errors2 "file_manager/src/core/errors"
)

type AuthService interface {
	Authenticate(authPackage request2.AuthPackage) (*request2.Authentication, errors2.Error)
}
