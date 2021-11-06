package ports

import (
	"file_manager/internal/adapter/repositories"
	"file_manager/internal/models"
)

type UserCommandRepositoryPort interface {
	Insert(user *models.User) (*models.User, error)
}

func InitUserCommandRepositoryPort(userCommandRepository *repositories.UserCommandRepository) UserCommandRepositoryPort {
	return userCommandRepository
}
