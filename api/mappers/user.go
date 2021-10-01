package mappers

import (
	"file_manager/api/resources"
	"file_manager/internal/models"
)

func ConvertUserModelToResource(user *models.User) *resources.User {
	return &resources.User{
		Id:       user.Id,
		FullName: user.FullName,
		Username: user.Username,
	}
}
