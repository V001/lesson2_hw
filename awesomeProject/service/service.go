package service

import (
	"awesomeProject/storage"
	"fmt"
)

type Service struct {
	storage *storage.Storage
	User    IUserService
}

func NewService(storage *storage.Storage) (*Service, error) {
	if storage == nil {
		return nil, fmt.Errorf("storage is empty")
	}
	var service Service
	service.User = NewUserService(storage)

	return &service, nil
}
