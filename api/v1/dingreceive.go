package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zx/model/common/response"
	"zx/service/dingtalk"
)

func (api *ApiGroup) Init(c *gin.Context) {
	var obj dingtalk.OutGoingModel
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	if msg := dingtalk.ExecDingCommand(obj); msg != nil {
		//发送执行结果给发起钉钉

		response.OkWithDetailed(string(msg), "创建成功", c)
	}
}
