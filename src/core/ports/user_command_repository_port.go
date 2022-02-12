package ports

import (
	"file_manager/src/adapter/database/repositories"
	"file_manager/src/core/entities"
)

type UserCommandRepositoryPort interface {
	Insert(user *entities.User) error
}

func NewUserCommandRepositoryPort(userCommandRepository *repositories.UserCommandRepository) UserCommandRepositoryPort {
	return userCommandRepository
}
