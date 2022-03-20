package initialize

import (
	// _ "zx/docs"
	"zx/global"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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
	// systemRouter := router.RouterGroupApp.System
	// exampleRouter := router.RouterGroupApp.Example
	// autocodeRouter := router.RouterGroupApp.Autocode
	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	// {
	// 	systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	// 	systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	// }
	// PrivateGroup := Router.Group("")
	// PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	// {
	// 	systemRouter.InitApiRouter(PrivateGroup)                    // 注册功能api路由
	// 	systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
	// 	systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
	// 	systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
	// 	systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
	// 	systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
	// 	systemRouter.InitAutoCodeRouter(PrivateGroup)               // 创建自动化代码
	// 	systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
	// 	systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
	// 	systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
	// 	systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
	// 	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	// 	exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
	// 	exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由

	// 	// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.
	// 	autocodeRouter.InitSysAutoCodeExampleRouter(PrivateGroup)
	// 	autocodeRouter.InitProjectinfosRouter(PrivateGroup)
	// 	autocodeRouter.InitK8sinfosRouter(PrivateGroup)
	// 	autocodeRouter.InitGroupinfosRouter(PrivateGroup)
	// 	autocodeRouter.InitRelationsRouter(PrivateGroup)
	// 	// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
	// }

	// InstallPlugin(PublicGroup, PrivateGroup) // 安装插件

	global.ZX_LOG.Info("router register success")
	return Router
}