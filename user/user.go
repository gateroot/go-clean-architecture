package user

import (
	"awesomeProject1/model"
	"awesomeProject1/trace"
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
	tracer   trace.Tracer
}

func NewUserUsecase(userRepo repository.UserRepository, tracer trace.Tracer) UserUsecase {
	return userUsecase{userRepo, tracer}
}

func (uu userUsecase) Get(userID int) (*model.User, error) {
	uu.tracer.Trace("Get User")
	return uu.userRepo.Get(userID)
}

func (uu userUsecase) Add(user *model.User) (*model.User, error) {
	uu.tracer.Trace("Add User")
	return uu.userRepo.Add(user)
}

func (uu userUsecase) Edit(user *model.User) error {
	uu.tracer.Trace("Edit User")
	return uu.userRepo.Edit(user)
}

func (uu userUsecase) Delete(userID int) error {
	uu.tracer.Trace("Delete User")
	return uu.userRepo.Delete(userID)
}
