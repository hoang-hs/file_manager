package controllers

import (
	"file_manager/src/api/mappers"
	"file_manager/src/api/request"
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RegisterController struct {
	*baseController
	registerService services.RegisterService
}

func NewRegisterController(baseController *baseController, registerService services.RegisterService) *RegisterController {
	return &RegisterController{
		baseController:  baseController,
		registerService: registerService,
	}
}

func (r *RegisterController) SignUp(c *gin.Context) {
	registerRequest := request.RegisterRequest{}
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		log.Errorf("bind json fail, err:[%v]", err)
		r.BadRequest(c, err.Error())
		return
	}
	err := validator.New().Struct(registerRequest)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("query invalid, err: [%v]", err)
		}
		r.DefaultBadRequest(c)
		return
	}

	user, newErr := r.registerService.SignUp(&registerRequest)
	if newErr != nil {
		r.ErrorData(c, newErr)
		return
	}
	resUser := mappers.ConvertUserEntityToResource(user)
	r.Success(c, resUser)
}
