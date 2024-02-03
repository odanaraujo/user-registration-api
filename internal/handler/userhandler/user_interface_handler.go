package userhandler

import (
	"github.com/odanaraujo/user-api/internal/service/userservice"
	"net/http"
)

func NewUserHandler(service userservice.UserService) UserHandler {
	return &handler{service: service}
}

type handler struct {
	service userservice.UserService
}

type UserHandler interface {
	CreateUser(w http.ResponseWriter, r *http.Request)
}
