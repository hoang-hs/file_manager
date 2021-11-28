package controllers

import (
	"file_manager/api/mappers"
	"file_manager/api/services"
	"file_manager/configs"
	"file_manager/internal/common/log"
	"file_manager/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type LoginController struct {
	*baseController
	authService services.IAuthService
}

func NewLoginController(baseController *baseController, authService services.IAuthService) *LoginController {
	return &LoginController{
		baseController: baseController,
		authService:    authService,
	}
}

func (l *LoginController) Login(c *gin.Context) {
	authPackage := entities.AuthPackage{}
	if err := c.ShouldBindJSON(&authPackage); err != nil {
		log.Errorf("bind json fail, err:[%v]", err)
		l.BadRequest(c, err.Error())
		return
	}
	err := validator.New().Struct(authPackage)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("query invalid, err: [%v]", err)
		}
		l.DefaultBadRequest(c)
		return
	}

	authentication, newErr := l.authService.Authenticate(authPackage)
	if newErr != nil {
		l.ErrorData(c, newErr)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   authentication.AccessToken,
		Expires: time.Now().Add(time.Minute * configs.Get().ExpiredDuration),
	})
	c.SetCookie("access_token", authentication.AccessToken, 15, "/", "0.0.0.0", false, true)
	resAuth := mappers.ConvertAuthenticationEntityToResource(authentication)
	l.Success(c, resAuth)
}
