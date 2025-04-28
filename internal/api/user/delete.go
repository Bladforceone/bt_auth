package user

import (
	desc "bt_auth/pkg/user_v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := s.userService.Delete(ctx, request.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
