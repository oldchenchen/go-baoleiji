package Middleware

import (
	"ZgoMast/application/Constants"
	. "ZgoMast/application/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// JWTAuthorization 验证token
func JWTAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"code":    Constants.CodeAuthticateFail,
				"message": Constants.TokenIsInvalid,
			})
			c.Abort()
			return
		}

		zap.S().Debugf("get token: %#v", token)

		j := NewJWT()
		// parseToken 解析token包含的信息
		claims, err := j.GetPayloadByToken(token)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    Constants.CodeAuthticateFail,
				"message": err.Error(),
			})
			c.Abort()
			return
		}

		// 存储认证的载荷信息和token，保留至后面使用。
		c.Set("claims", claims)
		c.Set("access_token", token)
		c.Next()
	}
}
