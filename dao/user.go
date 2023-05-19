package dao

import (
	"pan/common"
)

//用户信息
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

//注册新用户到数据库
func (u *User) RegisterNewUser() {
	common.GetGormDB().Create(u)
}

//去查找用户和密码在不在数据库，在就返回true，否则false
func (u *User) HaveUser() bool {
	user := new(User)
	db := common.GetGormDB()
	if err := db.First(&user, "user_name=? and password=?", u.UserName, u.Password).Error; err != nil {
		return false
	} else {
		return true
	}
}
