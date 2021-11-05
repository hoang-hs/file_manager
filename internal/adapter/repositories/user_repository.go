package repositories

import (
	"database/sql"
	log2 "file_manager/internal/common/log"
	"file_manager/internal/enums"
	"file_manager/internal/models"
	"github.com/google/uuid"
)

type UserRepository struct {
	baseRepository
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		*NewBaseRepository(db),
	}
}

func (u *UserRepository) FindByUsername(username string) (*models.User, error) {
	user := &models.User{}
	stmt, err := u.db.Prepare(`SELECT ID, FULL_NAME, USER_NAME, PASSWORD FROM USERS_CTRL WHERE USER_NAME = ?`)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = stmt.Close(); err != nil {
			log2.Fatalf("cannot close stmt,err: [%v]", err)
		}
	}()

	err = stmt.QueryRow(username).Scan(&user.Id, &user.FullName, &user.Username, &user.Password)
	switch {
	case err == sql.ErrNoRows:
		return nil, enums.ErrEntityNotFound
	case err != nil:
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) Insert(user *models.User) (*models.User, error) {
	tx, err := u.db.Begin()
	if err != nil {
		log2.Errorf("u.db cannot begin, err:[%v]", err)
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
			log2.Fatalf("cannot close stmt, err:[%v]", err)
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