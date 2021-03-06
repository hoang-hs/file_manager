package usecases

import (
	"file_manager/src/api/request"
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"file_manager/src/configs"
	"file_manager/src/core/entities"
	"file_manager/src/core/errors"
	"file_manager/src/core/ports"
	"file_manager/src/helpers"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type AuthUseCase struct {
	userQueryRepositoryPort ports.UserQueryRepositoryPort
}

func NewAuthUseCase(userQueryRepositoryPort ports.UserQueryRepositoryPort) services.AuthService {
	return &AuthUseCase{
		userQueryRepositoryPort: userQueryRepositoryPort,
	}
}

func (a *AuthUseCase) Authenticate(authRequest *request.AuthRequest) (*entities.Authentication, errors.Error) {
	username := authRequest.Username
	user, err := a.userQueryRepositoryPort.FindByUsername(username)
	if err == errors.ErrEntityNotFound {
		log.Errorf("Can not find user with username: %s", username)
		return nil, errors.ErrUnAuthenticated
	}
	if err != nil {
		log.Errorf("Can not find user, err:[%s]", err)
		return nil, errors.ErrSystemError
	}
	if !a.validatePassword(user.Password, authRequest.Password) {
		log.Errorf("Fail when validate password for username: %s", user.Username)
		return nil, errors.ErrUnAuthenticated
	}

	tokenInfo := entities.AccessTokenInfo{
		UserId:          user.Id,
		ExpiredDuration: configs.Get().ExpiredDuration,
	}
	token, err := a.generateToken(tokenInfo)
	if err != nil {
		log.Errorf("Can not generate token: %s", err)
		return nil, errors.ErrSystemError
	}

	return &entities.Authentication{
		AccessToken: token,
		User:        user,
	}, nil
}

func (a *AuthUseCase) validatePassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func (a *AuthUseCase) generateToken(tokenInfo entities.AccessTokenInfo) (string, error) {
	privateKey, err := helpers.GetPrivateKey()
	if err != nil {
		log.Errorf("parse private key error, err:[%v]", err)
		return "", err
	}
	claims := &entities.Claims{
		Id: tokenInfo.UserId,
		StandardClaims: jwt.StandardClaims{
			Id:        tokenInfo.UserId,
			ExpiresAt: time.Now().Add(tokenInfo.ExpiredDuration).Unix(),
		},
	}
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := aToken.SignedString(privateKey)
	if err != nil {
		log.Errorf("Error when generate token: %s", err)
		return "", err
	}
	return token, nil
}
