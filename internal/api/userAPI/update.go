package userAPI

import (
	"bt_auth/internal/converter"
	desc "bt_auth/pkg/user_v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Server) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := s.userService.Update(ctx, request.GetId(), converter.ToUpdateRequestFromProto(request))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
