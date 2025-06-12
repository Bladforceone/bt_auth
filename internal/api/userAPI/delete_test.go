package userAPI

import (
	"bt_auth/internal/service"
	"bt_auth/internal/service/mocks"
	desc "bt_auth/pkg/user_v1"
	"context"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDelete(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mt minimock.Tester) service.UserService
	type args struct {
		ctx context.Context
		req *desc.DeleteRequest
	}
	var (
		ctx = context.Background()
		mt  = minimock.Tester(t)
		id  = int64(1)
		req = &desc.DeleteRequest{
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
				mock.DeleteMock.Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userServiceMock := tt.userServiceMock(mt)
			api := NewServer(userServiceMock)

			_, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
