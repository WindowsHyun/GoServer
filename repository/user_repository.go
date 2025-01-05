package repository

import "GoServer/model"

type UserRepository interface {
	GetUserByID(int) (*model.User, error)
	SaveUser(*model.User) error
}
