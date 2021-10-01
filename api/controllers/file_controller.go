package controllers

import (
	"file_manager/api/mappers"
	"file_manager/bootstrap"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	BaseController
}

func (o *FileController) Display(c *gin.Context) {
	files, err := bootstrap.FileService.GetFile(c)
	if err != nil {
		o.Error(c, err.GetHttpCode(), err.GetMessage())
	}
	data := mappers.ConvertDirEntitiesToResource(files)
	o.Success(c, data)

}

func (o *FileController) UploadFile(c *gin.Context) {
	err := bootstrap.FileService.UploadFile(c)
	if err != nil {
		o.Error(c, err.GetHttpCode(), err.GetMessage())
	}
	o.Success(c, "uploadOk")
}

func (o *FileController) DeleteFile(c *gin.Context) {
	err := bootstrap.FileService.DeleteFile(c)
	if err != nil {
		o.Error(c, err.GetHttpCode(), err.GetMessage())
	}
	o.Success(c, "deleteOk")
}
