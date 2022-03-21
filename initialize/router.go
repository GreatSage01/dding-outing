package initialize

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"zx/global"
	"zx/router"
)

// 初始化总路由

func Routers() *gin.Engine {

	var Router = gin.Default()

	global.ZX_LOG.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	global.ZX_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.ZX_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例

	appRouter := router.RouterGroupApp
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	PrivateGroup := Router.Group("")
	{
		appRouter.InitOutgoingRouter(PrivateGroup)
	}

	global.ZX_LOG.Info("router register success")
	return Router
}
