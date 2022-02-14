package usecases

import (
	"file_manager/src/common/log"
	"file_manager/src/core/entities"
	"file_manager/src/core/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type FileUseCase struct {
}

func NewFileUseCase() *FileUseCase {
	return &FileUseCase{}
}

func (f *FileUseCase) GetFile(path string) (*entities.Files, errors.Error) {
	files, err := os.Open(path)
	if err != nil {
		log.Errorf("cannot open dir,err: [%v]", err)
		return nil, errors.NewCustomHttpError(http.StatusBadRequest, "dir invalid")
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

func (f *FileUseCase) UploadFile(c *gin.Context, path string) errors.Error {
	file, err := c.FormFile("file")
	if err != nil {
		log.Errorf("cannot get file from form file, err: [%v]", err)
		return errors.NewCustomHttpError(http.StatusBadRequest, "form file invalid")
	}
	path = path + "/"
	if err = c.SaveUploadedFile(file, path+file.Filename); err != nil {
		log.Errorf("cannot save file, err: [%v]", err)
		return errors.NewCustomHttpError(http.StatusBadRequest, "query wrong")
	}
	return nil
}

func (f *FileUseCase) DeleteFile(path string) errors.Error {
	err := os.RemoveAll(path)
	if err != nil {
		log.Errorf("cannot remove file, err: [%v]", err)
		return errors.NewCustomHttpError(http.StatusBadRequest, "query wrong")
	}
	return nil
}
