package converter

import (
	"bt_auth/internal/model"
	desc "bt_auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToUserFromService(user *model.User) *desc.User {
	var updatedAt *timestamppb.Timestamp
	if user.UpdatedAt.Valid {
		updatedAt = timestamppb.New(user.UpdatedAt.Time)
	}

	return &desc.User{
		Id:        user.ID,
		Info:      ToUserInfoFromService(user.Info),
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUserInfoFromService(info *model.UserInfo) *desc.UserInfo {
	return &desc.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     desc.Role(desc.Role_value[info.Role]),
	}
}

func ToUserInfoFromProto(info *desc.UserInfo) *model.UserInfo {
	return &model.UserInfo{
		Name:     info.Name,
		Email:    info.Email,
		Password: info.Password,
		Role:     info.Role.String(),
	}
}

func ToGetResponseFromProto(response *model.User) *desc.GetResponse {
	var updatedAt *timestamppb.Timestamp
	if response.UpdatedAt.Valid {
		updatedAt = timestamppb.New(response.UpdatedAt.Time)
	}

	return &desc.GetResponse{
		Id:        response.ID,
		Name:      response.Info.Name,
		Email:     response.Info.Email,
		Role:      desc.Role(desc.Role_value[response.Info.Role]),
		CreatedAt: timestamppb.New(response.CreatedAt),
		UpdatedAt: updatedAt,
	}
}

func ToUpdateRequestFromService(request *desc.UpdateRequest) *model.UserInfo {
	return &model.UserInfo{
		Name:  request.Name.Value,
		Email: request.Email.Value,
	}
}
