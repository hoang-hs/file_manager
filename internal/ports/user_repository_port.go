package ports

import (
	"file_manager/internal/adapter/repositories"
	"file_manager/internal/models"
)

type UserRepositoryPort interface {
	FindByUsername(username string) (*models.User, error)
	Insert(user *models.User) (*models.User, error)
}

func InitUserRepositoryPort(userRepository *repositories.UserRepository) UserRepositoryPort {
	return userRepository
}
