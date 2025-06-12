package userAPI

import (
	model "bt_auth/internal/model"
	"bt_auth/internal/service"
	"bt_auth/internal/service/mocks"
	desc "bt_auth/pkg/user_v1"
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGet(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mt minimock.Tester) service.UserService
	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}
	var (
		ctx = context.Background()
		mt  = minimock.Tester(t)
		id  = int64(1)
		req = &desc.GetRequest{
			Id: id,
		}
	)
	tests := []struct {
		name            string
		args            args
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "success",
			args: args{
				ctx: ctx,
				req: req,
			},
			err: nil,
			userServiceMock: func(mt minimock.Tester) service.UserService {
				mock := mocks.NewUserServiceMock(mt)
				mock.GetMock.Return(&model.User{
					ID: id,
					Info: &model.UserInfo{
						Name:     "",
						Email:    "",
						Password: "",
						Role:     "",
					},
				}, nil)
				return mock
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mt)
			api := NewServer(userServiceMock)

			resp, err := api.Get(tt.args.ctx, tt.args.req)
			if tt.err != nil {
				require.Equal(t, tt.err, err)
			} else {
				require.NoError(t, err)
				require.NotNil(t, resp)
				require.Equal(t, id, resp.GetId())
			}
		})
	}
}
