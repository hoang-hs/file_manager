package mappers

import (
	"file_manager/src/adapter/database/models"
	"file_manager/src/core/entities"
)

func ConvertUserModelToEntity(user *models.User) *entities.User {
	return &entities.User{
		Id:       user.Id,
		FullName: user.FullName,
		Username: user.Username,
		Password: user.Password,
	}
}

func ConvertUserEntityToModel(user *entities.User) *models.User {
	return &models.User{
		FullName: user.FullName,
		Username: user.Username,
		Password: user.Password,
	}
}
