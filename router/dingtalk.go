package router

import (
	v1 "zx/api/v1"

	"github.com/gin-gonic/gin"
)

func (r *RouterGroup) InitOutgoingRouter(Router *gin.RouterGroup) {
	dingreceiveRouter := Router.Group("dingtalk")
	var api = v1.ApiGroupApp
	{
		dingreceiveRouter.POST("", api.Init)
	}

}
