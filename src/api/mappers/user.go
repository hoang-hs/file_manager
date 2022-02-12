package mappers

import (
	"file_manager/src/adapter/database/models"
	"file_manager/src/api/resources"
)

func ConvertUserModelToResource(user *models.User) *resources.User {
	return &resources.User{
		Id:       user.Id,
		FullName: user.FullName,
		Username: user.Username,
	}
}
