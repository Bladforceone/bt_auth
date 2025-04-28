package repository

import (
	"bt_auth/internal/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.UserInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, id int64, user *model.UserInfo) error
	Delete(ctx context.Context, id int64) error
}
