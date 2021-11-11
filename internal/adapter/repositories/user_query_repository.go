package repositories

import (
	"database/sql"
	"file_manager/internal/common/log"
	"file_manager/internal/errors"
	"file_manager/internal/models"
)

type UserQueryRepository struct {
	*baseRepository
}

func NewUserQueryRepository(baseRepository *baseRepository) *UserQueryRepository {
	return &UserQueryRepository{
		baseRepository: baseRepository,
	}
}

func (q *UserQueryRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	stmt, err := q.db.Prepare(`SELECT ID, FULL_NAME, USER_NAME, PASSWORD FROM USERS_CTRL WHERE USER_NAME = ?`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = stmt.Close(); err != nil {
			log.Fatalf("cannot close stmt,err: [%v]", err)
		}
	}()
	err = stmt.QueryRow(username).Scan(&user.Id, &user.FullName, &user.Username, &user.Password)
	switch {
	case err == sql.ErrNoRows:
		return nil, errors.ErrEntityNotFound
	case err != nil:
		return nil, err
	}
	return user, nil
}
