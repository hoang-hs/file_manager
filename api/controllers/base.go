package controllers

import (
	"file_manager/api/resources"
	"file_manager/configs"
	"file_manager/internal/common/log"
	"file_manager/internal/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct {
	AppContext *ApplicationContext
}

func (b *BaseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
	c.Abort()
}

func (b *BaseController) Error(c *gin.Context, httpCode int, message string) {
	c.JSON(httpCode, resources.NewMessageResource(message))
	c.Abort()
}

func (b *BaseController) ErrorData(c *gin.Context, data errors.Error) {
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

func (b *BaseController) GetQuery(c *gin.Context) string {
	var query struct {
		Path string `form:"path"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		log.Errorf("bind query fail, err %s", err)
		b.DefaultBadRequest(c)
		return ""
	}
	if len(query.Path) == 0 {
		log.Errorf("path is nil")
		b.DefaultBadRequest(c)
		return ""
	}
	query.Path = configs.Get().Root + query.Path
	return query.Path
}
