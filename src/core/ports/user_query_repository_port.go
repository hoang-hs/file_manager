package ports

import (
	"file_manager/src/adapter/database/models"
	"file_manager/src/adapter/decorators"
)

type UserQueryRepositoryPort interface {
	FindByUsername(username string) (*models.User, error)
}

func NewUserQueryRepositoryPort(userQueryRepository *decorators.UserRepositoryDecorator) UserQueryRepositoryPort {
	return userQueryRepository
}
