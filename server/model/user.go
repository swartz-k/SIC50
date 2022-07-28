package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type userModel struct {
}

var UserModel = userModel{}

type User struct {
	*gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *userModel) getBaseModel() *gorm.DB {
	return db.Model(&User{}).Preload(clause.Associations)
}

func (u *userModel) GetByMail(mail string) (*User, error) {
	var user User
	err := u.getBaseModel().Where("email = ?", mail).First(&user).Error
	return &user, err
}
