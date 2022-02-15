package ports

import (
	"file_manager/src/core/entities"
)

type UserQueryRepositoryPort interface {
	FindByUsername(username string) (*entities.User, error)
}
