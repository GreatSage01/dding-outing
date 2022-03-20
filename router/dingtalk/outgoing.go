package dingtalk

import (
	v1 "zx/api/v1"

	"github.com/gin-gonic/gin"
)

type DingtalkRouter struct {
}

func (out *DingtalkRouter) InitOutgoingRouter(Router *gin.RouterGroup) {
	outgoingRouter := Router.Group("outgoing")
	var outgoingApi = v1.ApiGroupApp.OutgoingApiGroup.OutgoingApi
	{
		outgoingRouter.POST("test", outgoingApi.Init)
	}

}
