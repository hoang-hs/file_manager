package mappers

import (
	"file_manager/api/resources"
	"file_manager/internal/entities"
)

func ConvertDirEntitiesToResource(files *entities.Files) *resources.Files {
	var data []resources.File
	for _, file := range files.ListFile {
		data = append(data, resources.File{
			Type: file.Type,
			Name: file.Name,
			Size: file.Size,
		})
	}
	return &resources.Files{
		ListDir: data,
	}
}
