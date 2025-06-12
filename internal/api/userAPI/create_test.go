package userAPI

import (
	"bt_auth/internal/model"
	"bt_auth/internal/service"
	"bt_auth/internal/service/mocks"
	desc "bt_auth/pkg/user_v1"
	"context"
	"github.com/brianvoe/gofakeit"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

func TestCreate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mt minimock.Tester) service.UserService

	type args struct {
		ctx context.Context
		req *desc.CreateRequest
	}

	var (
		ctx   = context.Background()
		mt    = minimock.Tester(t)
		id    = gofakeit.Int64()
		pass  = gofakeit.Password(true, true, true, true, true, 10)
		name  = gofakeit.Name()
		email = gofakeit.Email()
		role  = desc.Role(gofakeit.Number(0, 1))
		info  = &model.UserInfo{
			Name:     name,
			Email:    email,
			Password: pass,
			Role:     role.String(),
		}

		req = &desc.CreateRequest{
			Info: &desc.UserInfo{
				Name:            name,
				Email:           email,
				Password:        pass,
				PasswordConfirm: pass,
				Role:            role,
			},
		}

		res = &desc.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name            string
		args            args
		want            *desc.CreateResponse
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mt minimock.Tester) service.UserService {
				mock := mocks.NewUserServiceMock(mt)
				mock.CreateMock.Expect(ctx, info).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  status.Error(codes.Internal, "service error"),
			userServiceMock: func(mt minimock.Tester) service.UserService {
				mock := mocks.NewUserServiceMock(mt)
				mock.CreateMock.Expect(ctx, info).Return(0, status.Error(codes.Internal, "service error"))
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userServiceMock := tt.userServiceMock(mt)
			api := NewServer(userServiceMock)

			resGet, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resGet)
		})
	}
}
