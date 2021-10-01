package entities

import "file_manager/internal/models"

type Authentication struct {
	AccessToken string       `json:"access_token"`
	User        *models.User `json:"user"`
}

type AuthPackage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
