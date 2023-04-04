package handlers

import (
	"awesomeProject/service"
	"awesomeProject/storage"
)

type Handlers struct {
	User *UserHandler
}

func NewHandlers(storage *storage.Storage, service *service.Service) *Handlers {
	return &Handlers{
		User: NewUserHandler(storage, service),
	}
}
