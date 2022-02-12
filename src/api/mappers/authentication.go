package mappers

import (
	"file_manager/src/api/request"
	"file_manager/src/api/resources"
)

func ConvertAuthenticationEntityToResource(authentication *request.Authentication) *resources.Authentication {
	userRes := ConvertUserModelToResource(authentication.User)
	return &resources.Authentication{
		AccessToken: authentication.AccessToken,
		User:        userRes,
	}
}
