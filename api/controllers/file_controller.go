package controllers

import (
	"file_manager/api/mappers"
	"file_manager/api/resources"
	"file_manager/internal/services"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	*baseController
	fileService *services.FileService
}

func NewFileController(baseController *baseController, fileService *services.FileService) *FileController {
	return &FileController{
		baseController: baseController,
		fileService:    fileService,
	}
}

func (o *FileController) Display(c *gin.Context) {
	path := o.GetQuery(c)
	if len(path) == 0 {
		return
	}
	files, err := o.fileService.GetFile(path)
	if err != nil {
		o.ErrorData(c, err)
		return
	}
	data := mappers.ConvertDirEntitiesToResource(files)
	o.Success(c, data)

}

func (o *FileController) UploadFile(c *gin.Context) {
	path := o.GetQuery(c)
	if len(path) == 0 {
		return
	}
	err := o.fileService.UploadFile(c, path)
	if err != nil {
		o.ErrorData(c, err)
		return
	}
	o.Success(c, resources.NewMessageResource("upload file successfully "))
}

func (o *FileController) DeleteFile(c *gin.Context) {
	path := o.GetQuery(c)
	if len(path) == 0 {
		return
	}
	err := o.fileService.DeleteFile(path)
	if err != nil {
		o.ErrorData(c, err)
		return
	}
	o.Success(c, resources.NewMessageResource("delete file successfully "))

}
