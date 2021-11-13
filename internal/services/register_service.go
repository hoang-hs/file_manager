package services

import (
	"file_manager/internal/common/log"
	"file_manager/internal/entities"
	"file_manager/internal/errors"
	"file_manager/internal/helpers"
	"file_manager/internal/models"
	"file_manager/internal/ports"
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

func (r *RegisterService) SignUp(registerPack *entities.RegisterPackage) (*models.User, errors.Error) {
	user, err := r.validate(registerPack)
	if err != nil {
		log.Errorf("user has exist, err:[%v]", err)
		return nil, err
	}

	modelUser, newErr := r.userCommandRepositoryPort.Insert(user)
	if newErr != nil {
		log.Errorf("error when insert, err: [%v]", err)
		return nil, errors.ErrSystemError
	}
	return modelUser, nil
}

func (r *RegisterService) validate(registerPack *entities.RegisterPackage) (*models.User, errors.Error) {
	user, _ := r.userQueryRepositoryPort.FindByUsername(registerPack.Username)
	if user != nil {
		return nil, errors.NewCustomHttpError(http.StatusConflict, "This username has exist")
	}
	password, err := helpers.HashPassword(registerPack.Password)
	if err != nil {
		return nil, errors.ErrSystemError
	}

	user = &models.User{
		FullName: registerPack.FullName,
		Username: registerPack.Username,
		Password: password,
	}
	return user, nil
}
