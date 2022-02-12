package mappers

import (
	"file_manager/src/api/resources"
	"file_manager/src/core/entities"
)

func ConvertUserEntityToResource(user *entities.User) *resources.User {
	return &resources.User{
		Id:       user.Id,
		FullName: user.FullName,
		Username: user.Username,
	}
}
