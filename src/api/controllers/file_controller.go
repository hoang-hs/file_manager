package controllers

import (
	"file_manager/src/api/mappers"
	"file_manager/src/api/resources"
	"file_manager/src/api/services"
	"file_manager/src/common/log"
	"file_manager/src/configs"
	"github.com/gin-gonic/gin"
)

type FileController struct {
	*baseController
	fileService services.FileService
}

func NewFileController(baseController *baseController, fileService services.FileService) *FileController {
	return &FileController{
		baseController: baseController,
		fileService:    fileService,
	}
}

func (f *FileController) Display(c *gin.Context) {
	path := f.GetQuery(c)
	if len(path) == 0 {
		return
	}
	files, err := f.fileService.GetFile(path)
	if err != nil {
		f.ErrorData(c, err)
		return
	}
	data := mappers.ConvertDirEntitiesToResource(files)
	f.Success(c, data)

}

func (f *FileController) UploadFile(c *gin.Context) {
	path := f.GetQuery(c)
	if len(path) == 0 {
		return
	}
	err := f.fileService.UploadFile(c, path)
	if err != nil {
		f.ErrorData(c, err)
		return
	}
	f.Success(c, resources.NewMessageResource("upload file successfully "))
}

func (f *FileController) DeleteFile(c *gin.Context) {
	path := f.GetQuery(c)
	if len(path) == 0 {
		return
	}
	err := f.fileService.DeleteFile(path)
	if err != nil {
		f.ErrorData(c, err)
		return
	}
	f.Success(c, resources.NewMessageResource("delete file successfully "))
}

func (f *FileController) GetQuery(c *gin.Context) string {
	var query struct {
		Path string `form:"path"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		log.Errorf("bind query fail, err %s", err)
		f.DefaultBadRequest(c)
		return ""
	}
	if len(query.Path) == 0 {
		log.Errorf("path is nil")
		f.DefaultBadRequest(c)
		return ""
	}
	query.Path = configs.Get().Root + query.Path
	return query.Path
}
