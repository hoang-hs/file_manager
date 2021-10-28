package services

import (
	"file_manager/internal/entities"
	"file_manager/internal/enums"
	"file_manager/internal/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) GetFile(path string) (*entities.Files, enums.Error) {
	files, err := os.Open(path)
	if err != nil {
		log.Errorf("cannot open dir,err: [%v]", err)
		return nil, enums.NewCustomHttpError(http.StatusBadRequest, "dir invalid")
	}
	fileInfo, err := files.Readdir(-1)
	defer func() {
		err = files.Close()
		if err != nil {
			log.Errorf("cannot close file,err: [%v]", err)
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

func (f *FileService) UploadFile(c *gin.Context, path string) enums.Error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("cannot get file from form file, err: [%v]", err)
		return enums.NewCustomHttpError(http.StatusBadRequest, "form file invalid")
	}
	path = path + "/"
	if err := c.SaveUploadedFile(file, path+file.Filename); err != nil {
		log.Errorf("cannot save file, err: [%v]", err)
		return enums.NewCustomHttpError(http.StatusBadRequest, "query wrong")
	}
	return nil
}

func (f *FileService) DeleteFile(path string) enums.Error {
	err := os.RemoveAll(path)
	if err != nil {
		log.Errorf("cannot remove file, err: [%v]", err)
		return enums.NewCustomHttpError(http.StatusBadRequest, "query wrong")
	}
	return nil
}
