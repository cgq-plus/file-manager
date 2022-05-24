package initialize

import (
	"file-manager/controller"
	"file-manager/global"
	"file-manager/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.CONFIG.Application.RunMode)
	engine := gin.New()
	// 跨域，如需跨域可以打开下面的注释
	engine.Use(middleware.Cors()) // 直接放行全部跨域请求
	// 健康监测
	engine.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "pong",
		})
		return
	})
	//文件列表
	engine.POST("/file/list", controller.List)
	//下载文件
	engine.POST("/file/down", controller.Download)
	//删除文件
	engine.POST("/file/delete", controller.Delete)
	return engine
}
