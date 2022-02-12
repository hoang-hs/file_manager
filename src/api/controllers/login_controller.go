package controllers

import (
	"file_manager/src/api/mappers"
	"file_manager/src/api/request"
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"file_manager/src/configs"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type LoginController struct {
	*baseController
	authService services.AuthService
}

func NewLoginController(baseController *baseController, authService services.AuthService) *LoginController {
	return &LoginController{
		baseController: baseController,
		authService:    authService,
	}
}

func (l *LoginController) Login(c *gin.Context) {
	authRequest := request.AuthRequest{}
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		log.Errorf("bind json fail, err:[%v]", err)
		l.BadRequest(c, err.Error())
		return
	}
	err := validator.New().Struct(authRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("query invalid, err: [%v]", err)
		}
		l.DefaultBadRequest(c)
		return
	}

	authentication, newErr := l.authService.Authenticate(&authRequest)
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
