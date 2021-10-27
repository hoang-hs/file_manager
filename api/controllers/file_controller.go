package controllers

import (
	"file_manager/api/mappers"
	"file_manager/api/resources"
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
	path := o.GetPath(c)
	if len(path) == 0 {
		return
	}
	files, err := o.AppContext.FileService.GetFile(path)
	if err != nil {
		o.ErrorData(c, err)
		return
	}
	data := mappers.ConvertDirEntitiesToResource(files)
	o.Success(c, data)

}

func (o *FileController) UploadFile(c *gin.Context) {
	path := o.GetPath(c)
	if len(path) == 0 {
		return
	}
	err := o.AppContext.FileService.UploadFile(c, path)
	if err != nil {
		o.ErrorData(c, err)
		return
	}
	o.Success(c, resources.NewMessageResource("upload file successfully "))
}

func (o *FileController) DeleteFile(c *gin.Context) {
	path := o.GetPath(c)
	if len(path) == 0 {
		return
	}
	err := o.AppContext.FileService.DeleteFile(path)
	if err != nil {
		o.ErrorData(c, err)
		return
	}
	o.Success(c, resources.NewMessageResource("delete file successfully "))

}
