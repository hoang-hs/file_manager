package controllers

import (
	"file_manager/src/api/resources"
	"file_manager/src/core/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type baseController struct {
}

func NewBaseController() *baseController {
	return &baseController{}
}

func (b *baseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (b *baseController) Error(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, resources.NewMessageResource(message))
	c.Abort()
}

func (b *baseController) ErrorData(c *gin.Context, data errors.Error) {
	httpCode := data.GetHttpCode()
	if httpCode <= 0 {
		httpCode = http.StatusBadRequest
	}
	c.JSON(httpCode, data)
}

func (b *baseController) BadRequest(c *gin.Context, message string) {
	b.Error(c, http.StatusBadRequest, message)
}

func (b *baseController) DefaultBadRequest(c *gin.Context) {
	b.BadRequest(c, "Invalid request")
}

func (b *baseController) InternetServerError(c *gin.Context) {
	b.Error(c, http.StatusInternalServerError, "System error")
}

func (b *baseController) Unauthorized(c *gin.Context) {
	b.Error(c, http.StatusUnauthorized, "Unauthorized")
}
