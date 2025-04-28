package userAPI

import (
	"bt_auth/internal/service"
	desc "bt_auth/pkg/user_v1"
)

type Server struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func NewServer(userService service.UserService) *Server {
	return &Server{
		userService: userService,
	}
}
