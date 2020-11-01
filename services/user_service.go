package services

import (
	"golang.org/x/crypto/bcrypt"
	"spike/datamodels"
	"spike/repositories"
)

type IUserService interface {
	IsPwdSuccess(userName string, password string) (user *datamodels.User, isOk bool)
	AddUser(user *datamodels.User) (userId int64, err error)
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository}
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func (u *UserService) IsPwdSuccess(userName string, password string) (user *datamodels.User, isOk bool) {
	var err error
	user, err = u.UserRepository.Select(userName)
	if err != nil {
		return
	}
	isOk, _ = validatePassword(password, user.HashPassword)
	if !isOk {
		return &datamodels.User{}, false
	}
	return
}

func (u *UserService) AddUser(user *datamodels.User) (userId int64, err error) {
	pwd, err := GeneratePassword(user.HashPassword)
	if err != nil {
		return
	}
	user.HashPassword = string(pwd)
	return u.UserRepository.Insert(user)
}

func validatePassword(password string, hashPassword string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password)); err != nil {
		return false, err
	}
	return true, nil
}

func GeneratePassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}
