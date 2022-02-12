package mappers

import (
	"file_manager/src/api/resources"
	"file_manager/src/core/entities"
)

func ConvertAuthenticationEntityToResource(authentication *entities.Authentication) *resources.Authentication {
	userRes := ConvertUserEntityToResource(authentication.User)
	return &resources.Authentication{
		AccessToken: authentication.AccessToken,
		User:        userRes,
	}
}
