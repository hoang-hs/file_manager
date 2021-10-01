package services

import (
	"file_manager/internal/entities"
	"file_manager/internal/enums"
	"file_manager/internal/helpers"
	"file_manager/internal/models"
	"file_manager/internal/repositories"
	"log"
	"net/http"
)

type RegisterService struct {
	userRepository repositories.UserRepository
}

func NewRegisterService(u repositories.UserRepository) *RegisterService {
	return &RegisterService{
		userRepository: u,
	}
}

func (r *RegisterService) SignUp(registerPack *entities.RegisterPackage) (*models.User, enums.Error) {
	user, err := r.validate(registerPack)
	if err != nil {
		log.Printf("error when validate, err:[%v]", err.Error())
		return nil, err
	}

	modelUser, newErr := r.userRepository.Insert(user)
	if newErr != nil {
		//log.Printf("error when insert, err: [%v]", err.Error())
		return nil, enums.ErrSystemError
	}
	return modelUser, nil
}

func (r *RegisterService) validate(registerPack *entities.RegisterPackage) (*models.User, enums.Error) {
	user, _ := r.userRepository.FindByUsername(registerPack.Username)
	if user != nil {
		log.Printf("username has exist")
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
