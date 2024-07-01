package dao

import (
	"api/models"
	"database/sql"
)

type UsersDAO struct {
	db *sql.DB
}

func CreateUserDao(db *sql.DB) *UsersDAO {
	return &UsersDAO{db}
}

func (u UsersDAO) Create(user models.User) (uint64, error) {
	return 0, nil
}
