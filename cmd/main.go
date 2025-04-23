package main

import (
	desc "bt_auth/pkg/auth_v1"
	"context"
	"github.com/brianvoe/gofakeit/v6"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"net"
)

type server struct {
	desc.UnimplementedAuthV1Server
}

func (s *server) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Create: %v", request)

	return &desc.CreateResponse{Id: gofakeit.Int64()}, nil
}

func (s *server) Get(ctx context.Context, request *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Get: %v", request)

	return &desc.GetResponse{
		Info: &desc.User{
			Id: request.Id,
			Info: &desc.UserInfo{
				Name:            "Вадим",
				Email:           gofakeit.Email(),
				Password:        gofakeit.Password(true, true, true, true, true, 10),
				PasswordConfirm: gofakeit.Password(true, true, true, true, true, 10),
				Role:            desc.Role_user,
			},
			CreatedAt: timestamppb.New(gofakeit.Date()),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func (s *server) Update(ctx context.Context, request *desc.UpdateRequest) (*emptypb.Empty, error) {
	log.Printf("Update: %v", request)
	return &emptypb.Empty{}, nil
}

func (s *server) Delete(ctx context.Context, request *desc.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Delete: %v", request)
	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterAuthV1Server(s, &server{})

	s.Serve(lis)
}
