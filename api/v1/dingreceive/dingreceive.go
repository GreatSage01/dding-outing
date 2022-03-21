package dingreceive

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zx/model"
	"zx/model/common/response"
)

type DingreceiveApi struct {
}

func (r *DingreceiveApi) Init(c *gin.Context) {
	var obj model.OutGoingModel
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	if msg := DingreceiveService.ExecDingCommand(obj); msg != nil {
		//发送执行结果给发起钉钉

		response.OkWithDetailed(string(msg), "创建成功", c)
	}
}
