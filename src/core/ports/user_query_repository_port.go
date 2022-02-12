package ports

import (
	"file_manager/src/adapter/decorators"
	"file_manager/src/core/entities"
)

type UserQueryRepositoryPort interface {
	FindByUsername(username string) (*entities.User, error)
}

func NewUserQueryRepositoryPort(userQueryRepository *decorators.UserRepositoryDecorator) UserQueryRepositoryPort {
	return userQueryRepository
}
