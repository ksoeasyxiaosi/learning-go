package repository

import (
	"webProject/model"
)

type User struct {
}

func (u User) Insert(userModel *model.User) (*model.User, error) {

	if err := model.DB.Create(&userModel).Error; err != nil {
		return nil, err
	}

	return userModel, nil
}
