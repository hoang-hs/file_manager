package ports

import (
	"file_manager/src/core/entities"
)

type UserCommandRepositoryPort interface {
	Insert(user *entities.User) error
}
