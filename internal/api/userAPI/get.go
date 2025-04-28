package userAPI

import (
	"bt_auth/internal/converter"
	desc "bt_auth/pkg/user_v1"
	"context"
)

func (s *Server) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return converter.ToGetResponseFromService(user), nil
}
