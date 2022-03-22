package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zx/global"
	"zx/model/common/response"
	"zx/service"
	"zx/service/dingtalk"
)

func (api *ApiGroup) Init(c *gin.Context) {
	var msg dingtalk.OutGoingModel
	if err := c.ShouldBindJSON(&msg); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	if res := dingtalk.ExecDingCommand(msg); res == nil {
		//发送执行结果给发起钉钉
		global.ZX_LOG.Info(string(res))
		global.ZX_LOG.Info("发送消息到钉钉")
		service.Dingsend(msg)
		response.OkWithDetailed(string(res), "创建成功", c)
	}
}
