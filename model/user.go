package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

//操作orm 必须先定义结构体,orm 是对结构体进行增删改查操作

// User 用户模型
type User struct {
	gorm.Model
	UserName       string `gorm:"unique"` //用户名唯一，不能重复
	PasswordDigest string //存储密文
}

const (
	PassWordCost = 12 //密码加密难度
)

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}

//orm 连接路径 => dsn := "root:root123@(localhost:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
