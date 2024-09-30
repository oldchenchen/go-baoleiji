package Init

import "github.com/gin-gonic/gin"
import "ZgoMast/application/Middleware"
import . "ZgoMast/application/router"

func InitRouter() *gin.Engine {

	Router := gin.Default()
	//使用日志中间件
	Router.Use(Middleware.GinLogger())
	TestGroup := Router.Group("/test")
	TestGroup.GET("/", func(c *gin.Context) {
		c.String(200, "ddd")
	})
	APIRouter := Router.Group("/api")
	InitUserRouter(APIRouter)

	APIRouter := Router.Group("/api")
	InitUserRouter()

	return Router
}
