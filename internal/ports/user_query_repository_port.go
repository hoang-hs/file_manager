package ports

import (
	"file_manager/internal/adapter/decorators"
	"file_manager/internal/models"
)

type UserQueryRepositoryPort interface {
	FindByUsername(username string) (*models.User, error)
}

func InitUserQueryRepositoryPort(userQueryRepository *decorators.UserRepositoryDecorator) UserQueryRepositoryPort {
	return userQueryRepository
}
