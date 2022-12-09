package service

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"to-do-list/model"
	"to-do-list/pkg/e"
	"to-do-list/pkg/util"
	"to-do-list/serializer"
)

// UserService 用户注册服务
type UserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,min=3,max=15" example:"FanOne"`
	Password string `form:"password" json:"password" binding:"required,min=5,max=16" example:"FanOne666"`
}

// /注册服务
func (service *UserService) Register() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64
	//验证数据库里是否纯在这个用户 ,模糊搜索db.where().First().count()
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户 (db.create())
	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		fmt.Printf(err.Error())
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

// Login 用户登陆函数
func (service *UserService) Login() serializer.Response {
	var user model.User
	code := e.SUCCESS
	//搜索 db.where().first()
	if err := model.DB.Where("user_name=?", service.UserName).First(&user).Error; err != nil {
		//如果查询不到，返回相应的错误
		if gorm.IsRecordNotFoundError(err) {
			util.LogrusObj.Info(err)
			code = e.ErrorNotExistUser
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
			}
		}
		util.LogrusObj.Info(err)
		code = e.ErrorDatabase
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	if !user.CheckPassword(service.Password) {
		code = e.ErrorNotCompare
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	token, err := util.GenerateToken(user.ID, service.UserName, 0)
	if err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorAuthToken
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.TokenData{User: serializer.BuildUser(user), Token: token},
		Msg:    e.GetMsg(code),
	}
}

// / todo test
func (service *UserService) RegisterTest() *serializer.Response {
	code := e.SUCCESS
	var user model.User
	var count int64
	//验证数据库里是否纯在这个用户 ,模糊搜索db.where().First().count()
	//err := model.DB.Model(&model.User{}).Select("user_name").Find(&user, 1).Error
	model.DB.Model(&model.User{}).Where("user_name=?", service.UserName).First(&user).Count(&count)
	//表单验证
	if count == 1 {
		code = e.ErrorExistUser
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	user.UserName = service.UserName
	//加密密码
	if err := user.SetPassword(service.Password); err != nil {
		util.LogrusObj.Info(err)
		code = e.ErrorFailEncryption
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	//创建用户 (db.create())
	if err := model.DB.Create(&user).Error; err != nil {
		util.LogrusObj.Info(err)
		fmt.Printf(err.Error())
		code = e.ErrorDatabase
		return &serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
		}
	}
	return &serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//结构体标签有： form(form 表单标签) , json(json 标签), binding(表单验证标签), example(样例标签),
//gorm (GORM标签), bson(MongoDB标签)，orm(Beego标签)， xorm(XORM标签)，ini(ini标签)
//学习地址：https://juejin.cn/post/7005465902804123679
