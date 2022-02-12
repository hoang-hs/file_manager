package repositories

import (
	"file_manager/src/adapter/database/mappers"
	"file_manager/src/common/log"
	"file_manager/src/core/entities"
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

func (u *UserCommandRepository) Insert(user *entities.User) error {
	userModel := mappers.ConvertUserEntityToModel(user)
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	userModel.Id = uuid.New().String()

	stmt, newRrr := tx.Prepare(`INSERT INTO USERS_CTRL (ID, FULL_NAME, USER_NAME, PASSWORD)
								VALUES(?, ?, ?, ?)`)
	if newRrr != nil {
		_ = tx.Rollback()
		return err
	}
	defer func() {
		if err := stmt.Close(); err != nil {
			log.Fatalf("cannot close stmt, err:[%v]", err)
		}
	}()
	_, err = stmt.Exec(userModel.Id, userModel.FullName, userModel.Username, userModel.Password)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	err = tx.Commit()
	return nil
}
