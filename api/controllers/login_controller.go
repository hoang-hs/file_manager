package controllers

import (
	"file_manager/api/mappers"
	"file_manager/configs"
	"file_manager/internal/common/log"
	"file_manager/internal/entities"
	"file_manager/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type LoginController struct {
	*baseController
	authService *services.AuthService
}

func NewLoginController(baseController *baseController, authService *services.AuthService) *LoginController {
	return &LoginController{
		baseController: baseController,
		authService:    authService,
	}
}

func (o *LoginController) Login(c *gin.Context) {
	authPackage := entities.AuthPackage{}
	if err := c.ShouldBindJSON(&authPackage); err != nil {
		log.Errorf("bind json fail, err:[%v]", err)
		o.BadRequest(c, err.Error())
		return
	}
	validate := validator.New()
	err := validate.Struct(authPackage)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("query invalid, err: [%v]", err)
		}
		o.DefaultBadRequest(c)
		return
	}

	authentication, newErr := o.authService.Authenticate(authPackage)
	if newErr != nil {
		o.ErrorData(c, newErr)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   authentication.AccessToken,
		Expires: time.Now().Add(time.Minute * configs.Get().ExpiredDuration),
	})

	resAuth := mappers.ConvertAuthenticationEntityToResource(authentication)
	o.Success(c, resAuth)
}
