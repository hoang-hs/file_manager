package usecases

import (
	"file_manager/src/api/request"
	"file_manager/src/common/log"
	"file_manager/src/core/entities"
	"file_manager/src/core/errors"
	"file_manager/src/core/ports"
	"file_manager/src/helpers"
	"net/http"
)

type RegisterService struct {
	userQueryRepositoryPort   ports.UserQueryRepositoryPort
	userCommandRepositoryPort ports.UserCommandRepositoryPort
}

func NewRegisterService(
	userQueryRepositoryPort ports.UserQueryRepositoryPort,
	userCommandRepositoryPort ports.UserCommandRepositoryPort,
) *RegisterService {
	return &RegisterService{
		userQueryRepositoryPort:   userQueryRepositoryPort,
		userCommandRepositoryPort: userCommandRepositoryPort,
	}
}

func (r *RegisterService) SignUp(registerRequest *request.RegisterRequest) (*entities.User, errors.Error) {
	user, err := r.validate(registerRequest)
	if err != nil {
		log.Errorf("user has exist, err:[%v]", err)
		return nil, err
	}

	newErr := r.userCommandRepositoryPort.Insert(user)
	if newErr != nil {
		log.Errorf("error when insert, err: [%v]", err)
		return nil, errors.ErrSystemError
	}
	return user, nil
}

func (r *RegisterService) validate(registerRequest *request.RegisterRequest) (*entities.User, errors.Error) {
	user, _ := r.userQueryRepositoryPort.FindByUsername(registerRequest.Username)
	if user != nil {
		return nil, errors.NewCustomHttpError(http.StatusConflict, "This username has exist")
	}
	password, err := helpers.HashPassword(registerRequest.Password)
	if err != nil {
		return nil, errors.ErrSystemError
	}

	user = &entities.User{
		FullName: registerRequest.FullName,
		Username: registerRequest.Username,
		Password: password,
	}
	return user, nil
}
