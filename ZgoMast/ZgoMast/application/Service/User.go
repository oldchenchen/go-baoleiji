package Service

import (
	"ZgoMast/application/Constants"
	. "ZgoMast/application/Model"
	. "ZgoMast/application/utils"
	"ZgoMast/application/validator"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
用户登录认证
*/

func UserLogin(ctx *gin.Context) (user User, err error) {
	user = User{}
	if err = ctx.ShouldBindJSON(&user); err != nil {
		return user, err
	}
	// 校验输入字段
	if err = validator.UserValidator(&user); err != nil {
		return user, err
	}
	user.GetOneByAccount(user.Username)
	if user.ID < 1 {
		return user, errors.New(Constants.NoSuchUser)
	}
	ret := CheckPassword(user.HashPassword, user.Password)
	if !ret {
		err = errors.New(Constants.PasswordError)
		return
	}

	return
}

/**
创建用户
*/

func CreateUser(ctx *gin.Context) (uint, error) {
	user := User{}
	var err error
	if err = ctx.ShouldBindJSON(&user); err != nil {
		return 0, err
	}

	// 校验输入字段
	if err = validator.UserValidator(&user); err != nil {
		fmt.Println("err:::", err)
		return 0, err
	}

	user.HashPassword, err = MakeHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	return user.Insert()
}

/**
获取指定ID的用户
*/

func GetUserById(id uint) (user User) {
	user = User{}
	user.GetOneByid(int(id))
	return user
}

/**
根据账户信息(用户名、手机、邮箱)获取用户
*/

func GetUserByAccount(account string) (user User) {
	user = User{}
	user.GetOneByAccount(account)
	return user
}

/**
获取所有用户
*/

func GetAllUser() []User {
	user := User{}
	return user.GetALL()
}

/**
更新密码
*/

func ChangeUserPassword(user User, RawPassword string) {
	password, _ := MakeHashPassword(RawPassword)
	user.Chpasswd(password)
}
