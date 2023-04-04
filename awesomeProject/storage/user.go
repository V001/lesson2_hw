package storage

import (
	"awesomeProject/model"
	"fmt"
)

type IUserRepository interface {
	GetUser(ID uint) (model.User, error)
	DeleteUser(ID uint) error
	CreateUser(item model.User) error
}

type UserRepositry struct {
	DB []model.User
}

func NewUserRepositry() *UserRepositry {
	return &UserRepositry{}
}

func (r *UserRepositry) GetUser(ID uint) (model.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositry) DeleteUser(ID uint) error {
	//TODO implement me
	panic("implement me")
}

func (r *UserRepositry) CreateUser(item model.User) error {
	r.DB = append(r.DB, item)
	fmt.Println(r.DB)
	return nil
}
