package service

import (
	"goli/api"
	"goli/model"
)

// Register 用户注册
func Register(params api.RegisterParams) (model.User, error) {
	user := model.User{
		UserName:       params.Username,
		PasswordDigest: params.Password,
	}

	// 创建用户
	if err := model.DB.Create(user).Error; err != nil {
		return model.User{}, err
	}

	return user, nil
}
