package mysql

import (
	"GoServer/model"
	"database/sql"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) GetUserByID(id int) (*model.User, error) {
	return &model.User{}, nil
}

func (r *MySQLUserRepository) SaveUser(user *model.User) error {
	return nil
}
