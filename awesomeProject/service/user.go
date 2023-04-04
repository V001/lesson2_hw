package service

import (
	"awesomeProject/model"
	"awesomeProject/storage"
	"fmt"
)

type UserService struct {
	repo *storage.Storage
}

func NewUserService(repo *storage.Storage) *UserService {
	return &UserService{repo: repo}
}

type IUserService interface {
	Create(user model.User) error
	Get(ID uint) (model.User, error)
	Update(user model.User) error
	Delete(ID string) error
}

func (s UserService) Create(user model.User) error {
	err := s.repo.User.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s UserService) Get(ID uint) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Update(book model.User) error {
	//TODO implement me
	panic("implement me")
}

func (s UserService) Delete(ID string) error {
	fmt.Println("Hellow world from delete")
	return nil
}
