package user

import (
	"awesomeProject1/model"
	"awesomeProject1/user/repository"
)

type UserUsecase interface {
	Get(userID int) (*model.User, error)
	Add(user *model.User) (*model.User, error)
	Edit(user *model.User) error
	Delete(userID int) error
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(userRepo repository.UserRepository) UserUsecase {
	return userUsecase{userRepo}
}

func (uu userUsecase) Get(userID int) (*model.User, error) {
	return uu.userRepo.Get(userID)
}

func (uu userUsecase) Add(user *model.User) (*model.User, error) {
	return uu.userRepo.Add(user)
}

func (uu userUsecase) Edit(user *model.User) error {
	return uu.userRepo.Edit(user)
}

func (uu userUsecase) Delete(userID int) error {
	return uu.userRepo.Delete(userID)
}





