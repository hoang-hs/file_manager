package services

import (
	"file_manager/internal/entities"
	"file_manager/internal/enums"
	"file_manager/internal/helpers"
	"file_manager/internal/models"
	"file_manager/internal/repositories"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService struct {
	expiredDuration time.Duration
	userRepository  repositories.UserRepository
}

func NewAuthService(u repositories.UserRepository, expiredDuration time.Duration) *AuthService {
	return &AuthService{
		expiredDuration: expiredDuration,
		userRepository:  u,
	}
}

func (auth *AuthService) Authenticate(authPackage entities.AuthPackage) (*entities.Authentication, enums.Error) {
	username := authPackage.Username
	user := &models.User{}
	var err error
	user, err = auth.userRepository.FindByUsername(username)
	if err == enums.ErrEntityNotFound {
		log.Printf("Can not find user with username: %s", username)
		return nil, *enums.ErrUnAuthenticated
	}
	if err != nil {
		log.Printf("Error when query to database: %s", err.Error())
		return nil, *enums.ErrSystemError
	}
	if !auth.validatePassword(user.Password, authPackage.Password) {
		log.Printf("Fail when validate password for username: %s", user.Username)
		return nil, *enums.ErrUnAuthenticated
	}

	tokenInfo := entities.AccessTokenInfo{
		UserId:          user.Id,
		ExpiredDuration: auth.expiredDuration,
	}
	token, err := auth.generateToken(tokenInfo)
	if err != nil {
		log.Printf("Can not generate token: %s", err)
		return nil, *enums.ErrSystemError
	}
	return &entities.Authentication{
		AccessToken: token,
		User:        user,
	}, nil
}

func (auth *AuthService) validatePassword(hashPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}

func (auth *AuthService) generateToken(tokenInfo entities.AccessTokenInfo) (string, error) {
	privateKey, err := helpers.GetPrivateKey()
	if err != nil {
		log.Printf("parse private key error, err:[%v]", err.Error())
		return "", err
	}
	claims := &entities.Claims{
		Username: tokenInfo.UserId,
		StandardClaims: jwt.StandardClaims{
			Id:        tokenInfo.UserId,
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
		},
	}
	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := aToken.SignedString(privateKey)
	if err != nil {
		log.Printf("Error when generate token: %s", err)
		return "", err
	}
	return token, nil
}
