package userAPI

import (
	"bt_auth/internal/converter"
	desc "bt_auth/pkg/user_v1"
	"context"
	"errors"
	"log"
)

func (s *Server) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	err := request.Validate()
	if err != nil {
		log.Printf("Validate error: %v", err)
		return nil, err
	}
	info := request.GetInfo()
	if info == nil {
		return nil, errors.New("info is nil")
	}
	id, err := s.userService.Create(ctx, converter.ToUserInfoFromProto(request.GetInfo()))
	if err != nil {
		return nil, err
	}
	return &desc.CreateResponse{Id: id}, nil
}
