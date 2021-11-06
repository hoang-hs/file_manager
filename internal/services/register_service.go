package services

import (
	log2 "file_manager/internal/common/log"
	"file_manager/internal/entities"
	"file_manager/internal/enums"
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

func (r *RegisterService) SignUp(registerPack *entities.RegisterPackage) (*models.User, enums.Error) {
	user, err := r.validate(registerPack)
	if err != nil {
		log2.Errorf("error when validate, err:[%v]", err)
		return nil, err
	}

	modelUser, newErr := r.userCommandRepositoryPort.Insert(user)
	if newErr != nil {
		log2.Errorf("error when insert, err: [%v]", err)
		return nil, enums.ErrSystemError
	}
	return modelUser, nil
}

func (r *RegisterService) validate(registerPack *entities.RegisterPackage) (*models.User, enums.Error) {
	user, _ := r.userQueryRepositoryPort.FindByUsername(registerPack.Username)
	if user != nil {
		log2.Info("username has exist")
		return nil, enums.NewCustomHttpError(http.StatusConflict, "This username has exist")
	}
	password, err := helpers.HashPassword(registerPack.Password)
	if err != nil {
		return nil, enums.ErrSystemError
	}

	user = &models.User{
		FullName: registerPack.FullName,
		Username: registerPack.Username,
		Password: password,
	}
	return user, nil
}
