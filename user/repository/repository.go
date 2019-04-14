package repository

import "awesomeProject1/model"

type UserRepository interface {
	Get(userID int) (*model.User, error)
	Add(user *model.User) (*model.User, error)
	Edit(user *model.User) error
	Delete(userID int) error
}
