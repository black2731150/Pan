package dao

import (
	"errors"
	"pan/common"

	"gorm.io/gorm"
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
func (u *User) UserNameLogin() bool {
	user := new(User)
	db := common.GetGormDB()
	if err := db.First(&user, "user_name=? and password=?", u.UserName, u.Password).Error; err != nil {
		return false
	} else {
		return true
	}
}

//去查找那个用户名存不存在，如果存在
func (u *User) HaveTheUserName() bool {
	db := common.GetGormDB()
	if err := db.First(NewUser(), "user_name=?", u.UserName).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		} else {
			return true
		}
	}

	return false
}

//检测用户邮箱是否已经注册
func (u *User) HaveTheEmail() bool {
	db := common.GetGormDB()
	if err := db.First(NewUser(), "email=?", u.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRegistered) {
			return false
		}
	} else {
		return true
	}

	return false
}
