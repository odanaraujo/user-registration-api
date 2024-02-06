package userservice

import (
	"context"
	"github.com/odanaraujo/user-api/internal/dto"
	"github.com/odanaraujo/user-api/internal/repository/userepository"
)

func NewUserService(repo userepository.UserRepository) UserService {
	return &service{repo: repo}
}

type service struct {
	repo userepository.UserRepository
}

type UserService interface {
	CreateUser(ctx context.Context, dto dto.CreateUserDto) error
}
