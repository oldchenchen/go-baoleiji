package Api

import (
	"ZgoMast/application/Constants"
	. "ZgoMast/application/Service"
	. "ZgoMast/application/Utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

/*
*
用户认证登陆
*/
func UserAuthenticate(ctx *gin.Context) {
	user, err := UserLogin(ctx)
	if err != nil || user.ID < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    Constants.CodeNoSuchUser,
			"message": err.Error(),
		})
		return
	}

	cnewJwt := NewJWT()
	publicClaims := PublicClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
	}

	token, err := cnewJwt.AccessToken(publicClaims)
	if err != nil {
		panic(err.Error())
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    Constants.CodeSuccess,
		"message": Constants.Success,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

/**
创建用户
*/

func UserCreate(ctx *gin.Context) {

	id, err := CreateUser(ctx)
	if err != nil || id < 1 {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    Constants.CodeCreateUserFail,
			"message": Constants.CreateUserFail,
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"code":    Constants.CodeSuccess,
		"message": Constants.CreateUserSuccess,
		"data":    id,
	})
}
