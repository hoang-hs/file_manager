package controllers

import (
	"file_manager/api/mappers"
	"file_manager/internal/entities"
	"github.com/gin-gonic/gin"
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
		o.BadRequest(c, err.Error())
		return
	}

	authentication, err := o.AppContext.AuthService.Authenticate(authPackage)
	if err != nil {
		o.ErrorData(c, err)
		return
	}

	http.SetCookie(c.Writer, &http.Cookie{
		Name:    "token",
		Value:   authentication.AccessToken,
		Expires: time.Now().Add(time.Minute * 15),
	})

	resAuth := mappers.ConvertAuthenticationEntityToResource(authentication)
	o.Success(c, resAuth)
}
