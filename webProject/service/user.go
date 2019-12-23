package service

import (
	"github.com/pkg/errors"
	"webProject/model"
	"webProject/repository"
	"webProject/utils"
)

type User struct {
}

func (u User) Insert(userModel *model.User) (*model.User, error) {

	err := userModel.ValidateInsert()
	if err != nil {
		return nil, errors.Wrap(err, utils.Trans("Validate.User.InsertDefeat"))
	}

	userRep := repository.User{}
	userRep.Insert(userModel)

	return userModel, nil
}
