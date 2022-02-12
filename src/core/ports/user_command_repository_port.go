package ports

import (
	"file_manager/src/adapter/database/models"
	"file_manager/src/adapter/database/repositories"
)

type UserCommandRepositoryPort interface {
	Insert(user *models.User) (*models.User, error)
}

func NewUserCommandRepositoryPort(userCommandRepository *repositories.UserCommandRepository) UserCommandRepositoryPort {
	return userCommandRepository
}
