package controllers

import (
	"context"
	"file_manager/api/resources"
	"file_manager/internal/enums"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
	ctx *context.Context
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (b *BaseController) SetContext(ctx *context.Context) {
	b.ctx = ctx
}

func (b *BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	c.Abort()
}

func (b *BaseController) Error(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, resources.NewErrorResource(message))
	c.Abort()
}

func (b *BaseController) ErrorData(c *gin.Context, data enums.Error) {
	httpCode := data.GetHttpCode()
	if httpCode <= 0 {
		httpCode = http.StatusBadRequest
	}
	c.JSON(httpCode, data)
}

func (b *BaseController) BadRequest(c *gin.Context, message string) {
	b.Error(c, http.StatusBadRequest, message)
}

func (b *BaseController) DefaultBadRequest(c *gin.Context) {
	b.BadRequest(c, "Invalid request")
}

func (b *BaseController) InternetServerError(c *gin.Context) {
	b.Error(c, http.StatusInternalServerError, "System error")
}

func (b *BaseController) Unauthorized(c *gin.Context) {
	b.Error(c, http.StatusUnauthorized, "Unauthorized")
}
