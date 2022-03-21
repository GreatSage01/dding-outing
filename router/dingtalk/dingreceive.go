package dingtalk

import (
	v1 "zx/api/v1"

	"github.com/gin-gonic/gin"
)

type DingtalkRouter struct {
}

func (r *DingtalkRouter) InitOutgoingRouter(Router *gin.RouterGroup) {
	dingreceiveRouter := Router.Group("dingreceive")
	var dingreceiveApi = v1.ApiGroupApp.DingreceiveApiGroup
	{
		dingreceiveRouter.POST("test", dingreceiveApi.Init)
	}

}
