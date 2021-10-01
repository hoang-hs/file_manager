package mappers

import (
	"file_manager/api/resources"
	"file_manager/internal/entities"
)

func ConvertAuthenticationEntityToResource(a *entities.Authentication) *resources.Authentication {
	userRes := ConvertUserModelToResource(a.User)
	return &resources.Authentication{
		AccessToken: a.AccessToken,
		User:        userRes,
	}
}
