package storage

import (
	"awesomeProject/config"
	"context"
)

type Storage struct {
	User IUserRepository
}

func NewStorage(ctx context.Context, cfg *config.Config) (*Storage, error) {
	var storage Storage
	storage.User = NewUserRepositry()
	return &storage, nil
}
