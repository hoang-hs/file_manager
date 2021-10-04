package services

import (
	"file_manager/internal/entities"
	"file_manager/internal/enums"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) GetFile(c *gin.Context) (*entities.Files, enums.Error) {
	root := os.Getenv("ROOT")
	str := root + c.Query("path")
	files, err := os.Open(str)
	if err != nil {
		log.Printf("cannot open dir,err: [%v]", err.Error())
		return nil, enums.NewCustomHttpError(http.StatusBadRequest, "dir invalid")
	}
	fileInfo, err := files.Readdir(-1)
	defer func() {
		err = files.Close()
		if err != nil {
			log.Printf("cannot close file,err: [%v]", err.Error())
		}
	}()
	var data []entities.File
	for _, file := range fileInfo {
		switch mode := file.Mode(); {
		case mode.IsDir():
			data = append(data, *entities.NewFile("folder", file.Name(), file.Size()))
		case mode.IsRegular():
			data = append(data, *entities.NewFile("file", file.Name(), file.Size()))
		}
	}
	listDir := entities.NewFiles(data)
	return listDir, nil
}

func (f *FileService) UploadFile(c *gin.Context) enums.Error {
	root := os.Getenv("ROOT")
	file, err := c.FormFile("file")
	if err != nil {
		log.Printf("cannot get file from formfile, err: [%v]", err.Error())
		return enums.NewCustomHttpError(http.StatusBadRequest, "form file invalid")
	}
	path := root + c.Query("path") + "/"
	if err := c.SaveUploadedFile(file, path+file.Filename); err != nil {
		log.Printf("cannot save file, err: [%v]", err.Error())
		return enums.NewCustomHttpError(http.StatusBadRequest, "query wrong")
	}
	return nil
}

func (f *FileService) DeleteFile(c *gin.Context) enums.Error {
	root := os.Getenv("ROOT")
	str := root + c.Query("path")
	err := os.RemoveAll(str)
	if err != nil {
		log.Printf("cannot remove file, err: [%v]", err.Error())
		return enums.NewCustomHttpError(http.StatusBadRequest, "query wrong")
	}
	return nil
}
