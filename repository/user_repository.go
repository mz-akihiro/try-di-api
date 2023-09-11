package repository

import (
	"database/sql"
)

type IUserRepository interface {
	GetUserByEmail(storedPassword *string, userId *int, email string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserByEmail(storedPassword *string, userId *int, email string) error {
	if err := ur.db.QueryRow("SELECT id, password FROM user WHERE email = ?", email).Scan(userId, storedPassword); err != nil {
		return err
	}
	return nil
}
