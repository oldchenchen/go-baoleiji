package router

import (
	. "ZgoMast/application/Utils"
	"github.com/gin-gonic/gin"
)
import . "ZgoMast/application/Api"

func InitUserRouter(r *gin.RouterGroup) {
	r.POST("/user", UserCreate)
	Register(r, []string{"GET", "POST", "PUT"}, "/user/authenticate", UserAuthenticate)
}
