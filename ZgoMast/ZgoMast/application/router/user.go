package router

import (
	"ZgoMast/application/Middleware"
	. "ZgoMast/application/utils"
	"github.com/gin-gonic/gin"
)
import . "ZgoMast/application/Api"

func InitUserRouter(r *gin.RouterGroup) {
	r.POST("/user", Middleware.JWTAuthorization(), UserCreate)
	Register(r, []string{"GET", "POST", "PUT"}, "/user/authenticate", UserAuthenticate)
}
