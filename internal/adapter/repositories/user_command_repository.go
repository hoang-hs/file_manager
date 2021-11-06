package repositories

import (
	"file_manager/internal/common/log"
	"file_manager/internal/models"
	"github.com/google/uuid"
)

type UserCommandRepository struct {
	*baseRepository
}

func NewUserCommandRepository(baseRepository *baseRepository) *UserCommandRepository {
	return &UserCommandRepository{
		baseRepository: baseRepository,
	}
}

func (c *UserCommandRepository) Insert(user *models.User) (*models.User, error) {
	tx, err := c.db.Begin()
	if err != nil {
		log.Errorf("u.db cannot begin, err:[%v]", err)
		return nil, err
	}
	user.Id = uuid.New().String()

	stmt, newRrr := tx.Prepare(`INSERT INTO USERS_CTRL (ID, FULL_NAME, USER_NAME, PASSWORD)
								VALUES(?, ?, ?, ?)`)
	if newRrr != nil {
		_ = tx.Rollback()
		return nil, err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Fatalf("cannot close stmt, err:[%v]", err)
		}
	}()
	_, err = stmt.Exec(user.Id, user.FullName, user.Username, user.Password)
	if err != nil {
		_ = tx.Rollback()
		return nil, err
	}
	err = tx.Commit()
	return user, nil
}
