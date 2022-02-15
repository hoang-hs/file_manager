package usecases

import (
	"file_manager/src/api/request"
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"file_manager/src/core/entities"
	"file_manager/src/core/errors"
	"file_manager/src/core/ports"
	"file_manager/src/helpers"
	"net/http"
)

type RegisterUseCase struct {
	userQueryRepositoryPort   ports.UserQueryRepositoryPort
	userCommandRepositoryPort ports.UserCommandRepositoryPort
}

func NewRegisterService(
	userQueryRepositoryPort ports.UserQueryRepositoryPort,
	userCommandRepositoryPort ports.UserCommandRepositoryPort,
) services.RegisterService {
	return &RegisterUseCase{
		userQueryRepositoryPort:   userQueryRepositoryPort,
		userCommandRepositoryPort: userCommandRepositoryPort,
	}
}

func (r *RegisterUseCase) SignUp(registerRequest *request.RegisterRequest) (*entities.User, errors.Error) {
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

func (r *RegisterUseCase) validate(registerRequest *request.RegisterRequest) (*entities.User, errors.Error) {
	user, err := r.userQueryRepositoryPort.FindByUsername(registerRequest.Username)
	if err != nil && err != errors.ErrEntityNotFound {
		log.Errorf("Can not find user, err:[%s]", err)
		return nil, errors.ErrSystemError
	}
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
