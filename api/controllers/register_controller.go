package controllers

import (
	"file_manager/api/mappers"
	"file_manager/api/services"
	"file_manager/internal/common/log"
	"file_manager/internal/entities"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterController struct {
	*baseController
	registerService services.IRegisterService
}

func NewRegisterController(baseController *baseController, registerService services.IRegisterService) *RegisterController {
	return &RegisterController{
		baseController:  baseController,
		registerService: registerService,
	}
}

func (r *RegisterController) SignUp(c *gin.Context) {
	registerPack := entities.RegisterPackage{}
	if err := c.ShouldBindJSON(&registerPack); err != nil {
		log.Errorf("bind json fail, err:[%v]", err)
		r.BadRequest(c, err.Error())
		return
	}
	err := validator.New().Struct(registerPack)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("query invalid, err: [%v]", err)
		}
		r.DefaultBadRequest(c)
		return
	}

	userModel, newErr := r.registerService.SignUp(&registerPack)
	if newErr != nil {
		r.ErrorData(c, newErr)
		return
	}
	resUser := mappers.ConvertUserModelToResource(userModel)
	r.Success(c, resUser)
}
