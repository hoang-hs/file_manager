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
	userRepositoryPort ports.UserRepositoryPort
}

func NewRegisterService(userRepositoryPort ports.UserRepositoryPort) *RegisterService {
	return &RegisterService{
		userRepositoryPort: userRepositoryPort,
	}
}

func (r *RegisterService) SignUp(registerPack *entities.RegisterPackage) (*models.User, enums.Error) {
	user, err := r.validate(registerPack)
	if err != nil {
		log2.Errorf("error when validate, err:[%v]", err)
		return nil, err
	}

	modelUser, newErr := r.userRepositoryPort.Insert(user)
	if newErr != nil {
		log2.Errorf("error when insert, err: [%v]", err)
		return nil, enums.ErrSystemError
	}
	return modelUser, nil
}

func (r *RegisterService) validate(registerPack *entities.RegisterPackage) (*models.User, enums.Error) {
	user, _ := r.userRepositoryPort.FindByUsername(registerPack.Username)
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
