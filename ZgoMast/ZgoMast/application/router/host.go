package router

import (
	"ZgoMast/application/Middleware"
	"ZgoMast/application/utils"
	"github.com/gin-gonic/gin"
)

func InitHostRouter(Router *gin.RouterGroup) {
	/**
	  主机相关的路由组
	*/
	HostRouter := Router.Group("host")

	{
		// 主机类别-添加
		utils.Register(HostRouter, []string{"POST"}, "category", Middleware.JWTAuthorization(), api.HostCategoryCreate)
	}
}
