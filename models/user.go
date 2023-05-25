package models

import (
	"errors"
	"pan/dao"

	"gorm.io/gorm"
)

//用户信息
type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Phonenum  string `json:"phonenum"`
	Filespace string `json:"filespace"`
}

func NewUser() *User {
	return &User{}
}

//注册新用户到数据库
func (u *User) RegisterNewUser() {
	dao.GetGormDB().Create(u)
}

//去查找用户和密码在不在数据库，在就返回true，否则false
func (u *User) UserNameLogin() bool {
	user := new(User)
	db := dao.GetGormDB()
	if err := db.First(&user, "user_name=? and password=?", u.UserName, u.Password).Error; err != nil {
		return false
	} else {
		return true
	}
}

//去查找那个用户名存不存在，如果存在
func (u *User) HaveTheUserName() bool {
	// fmt.Println("In function username:", u.UserName)
	db := dao.GetGormDB()
	newUser := NewUser()
	if u.UserName == "" {
		return true
	}
	if err := db.Table("users").First(newUser, "user_name=?", u.UserName).Error; err != nil {
		// fmt.Println(err)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// fmt.Println("用户名:", newUser.UserName, "没有被注册")
			return false
		}
	}

	// fmt.Println("用户名:", newUser.UserName, "已经被注册")
	return true
}

//检测用户邮箱是否已经注册
func (u *User) HaveTheEmail() bool {
	db := dao.GetGormDB()
	if err := db.Table("users").First(NewUser(), "email=?", u.Email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
	}

	return true
}

func (u *User) GetUsernameFormEmail() error {
	db := dao.GetGormDB()
	user := NewUser()
	if err := db.Table("users").First(user, "email=?", u.Email).Error; err != nil {
		return err
	}
	u.UserName = user.UserName
	return nil
}

func (u *User) GetEmailFromUserName() error {
	db := dao.GetGormDB()
	user := NewUser()
	if err := db.Table("users").First(user, "user_name", u.UserName).Error; err != nil {
		return err
	}
	u.Email = user.Email
	return nil
}

func (u *User) GetEmailAndUsernameFromID() error {
	db := dao.GetGormDB()
	user := NewUser()
	if err := db.Table("users").First(user, "id=?", u.ID).Error; err != nil {
		return err
	}

	u.Email = user.Email
	u.UserName = user.UserName
	return nil
}

func (u *User) GetUserIDFromUsername() error {
	db := dao.GetGormDB()
	user := NewUser()
	if err := db.Table("users").First(user, "user_name", u.UserName).Error; err != nil {
		return err
	}
	u.ID = user.ID
	return nil
}

func (u *User) GetUserIDFromEmail() error {
	db := dao.GetGormDB()
	user := NewUser()
	if err := db.Table("users").First(user, "email=?", u.Email).Error; err != nil {
		return err
	}
	u.ID = user.ID
	return nil
}

func (u *User) UpdatePassword(newpassword string) error {
	db := dao.GetGormDB()
	user := NewUser()
	result := db.Table("users").First(user, "email=?", u.Email)
	if result.Error != nil {
		return result.Error
	}

	user.Password = newpassword
	result = db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
