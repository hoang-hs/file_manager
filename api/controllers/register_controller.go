package controllers

import (
	"file_manager/api/mappers"
	"file_manager/internal/common/log"
	"file_manager/internal/entities"
	"file_manager/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterController struct {
	*baseController
	registerService *services.RegisterService
}

func NewRegisterController(baseController *baseController, registerService *services.RegisterService) *RegisterController {
	return &RegisterController{
		baseController:  baseController,
		registerService: registerService,
	}
}

func (o *RegisterController) SignUp(c *gin.Context) {
	registerPack := entities.RegisterPackage{}
	if err := c.ShouldBindJSON(&registerPack); err != nil {
		log.Errorf("bind json fail, err:[%v]", err)
		o.BadRequest(c, err.Error())
		return
	}
	validate := validator.New()
	err := validate.Struct(registerPack)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("query invalid, err: [%v]", err)
		}
		o.DefaultBadRequest(c)
		return
	}

	userModel, newErr := o.registerService.SignUp(&registerPack)
	if newErr != nil {
		o.ErrorData(c, newErr)
		return
	}
	resUser := mappers.ConvertUserModelToResource(userModel)
	o.Success(c, resUser)
}
