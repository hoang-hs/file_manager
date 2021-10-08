package controllers

import (
	"file_manager/api/mappers"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	BaseController
}

func NewFileController(appContext *ApplicationContext) *FileController {
	return &FileController{
		BaseController{
			AppContext: appContext,
		},
	}
}

func (o *FileController) Display(c *gin.Context) {
	files, err := o.AppContext.FileService.GetFile(c)
	if err != nil {
		o.Error(c, err.GetHttpCode(), err.GetMessage())
	}
	data := mappers.ConvertDirEntitiesToResource(files)
	o.Success(c, data)

}

func (o *FileController) UploadFile(c *gin.Context) {
	err := o.AppContext.FileService.UploadFile(c)
	if err != nil {
		o.Error(c, err.GetHttpCode(), err.GetMessage())
	}
	o.Success(c, "upload successfully")
}

func (o *FileController) DeleteFile(c *gin.Context) {
	err := o.AppContext.FileService.DeleteFile(c)
	if err != nil {
		o.Error(c, err.GetHttpCode(), err.GetMessage())
	}
	o.Success(c, "delete successfully")
}
