package controllers

import (
	"file_manager/api/mappers"
	"file_manager/bootstrap"
	"file_manager/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RegisterController struct {
	BaseController
}

func (o *RegisterController) SignUp(c *gin.Context) {
	registerPack := entities.RegisterPackage{}
	if err := c.ShouldBindJSON(&registerPack); err != nil {
		c.JSON(http.StatusBadRequest, "check request")
		return
	}
	userModel, err := bootstrap.RegisterService.SignUp(&registerPack)
	if err != nil {
		o.Error(c, err.GetHttpCode(), err.GetMessage())
		return
	}

	resUser := mappers.ConvertUserModelToResource(userModel)
	o.Success(c, resUser)
}