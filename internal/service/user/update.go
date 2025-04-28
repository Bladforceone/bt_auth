package user

import (
	"bt_auth/internal/model"
	"context"
)

func (s *serv) Update(ctx context.Context, id int64, user *model.UserInfo) error {
	err := s.userRepository.Update(ctx, id, user)
	if err != nil {
		return err
	}

	return nil
}
