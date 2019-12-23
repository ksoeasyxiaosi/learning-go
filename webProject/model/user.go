package model

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/jinzhu/gorm"
)

// User 用户模型
type User struct {
	gorm.Model
	UserName string
	Password string
	Nickname string
	Status   string
	Avatar   string `gorm:"size:1000"`
}

func (u User) ValidateInsert() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.UserName, validation.Required),
		validation.Field(&u.Password, validation.Required))
}
