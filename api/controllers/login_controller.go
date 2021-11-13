package controllers

import (
	"file_manager/api/mappers"
	"file_manager/configs"
	"file_manager/internal/common/log"
	"file_manager/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type LoginController struct {
	BaseController
}

func NewLoginController(appContext *ApplicationContext) *LoginController {
	return &LoginController{
		BaseController{
			AppContext: appContext,
		},
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

	authentication, newErr := o.AppContext.AuthService.Authenticate(authPackage)
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
