package dao

import (
	"pan/common"
)

type User struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phonenum string `json:"phonenum"`
}

func NewUser() *User {
	return &User{}
}

func (u *User) RegisterNewUser() {
	common.GetGormDB().Create(u)
}

func (u *User) HaveUser() bool {
	user := new(User)
	db := common.GetGormDB()
	if err := db.First(&user, "user_name=? and password=?", u.UserName, u.Password).Error; err != nil {
		return false
	} else {
		return true
	}
}
