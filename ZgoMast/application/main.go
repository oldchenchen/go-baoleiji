package main

import (
	. "ZgoMast/application/Config"
	"ZgoMast/application/Init"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"os"
	"path/filepath"
)

func main() {
	//1.网络路由
	router := Init.InitRouter()
	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello")
	})
	//2.全局配置文件
	pwd, _ := os.Getwd()
	configPath := filepath.Join(pwd, "application", "Config.json")
	if err := Initconfig(configPath); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}
	gin.SetMode(gin.DebugMode)
	// 3. 日志初始化
	if err := Init.InitLogger(Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}
	zap.S().Debugf("调试信息:%d", Conf.Port)
	// 数据库连接
	Init.InitDB(Conf.DatabaseConfig)
	//监听配置
	if err := router.Run(fmt.Sprintf(":%d", Conf.Port)); nil != err {
		zap.S().Panic("服务端启动失败：", err.Error())
	}
}
