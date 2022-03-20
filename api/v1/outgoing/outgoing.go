package outgoing

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"zx/model"
	"zx/model/common/response"
)

type OutgoingApi struct {
}

func (o *OutgoingApi) Init(c *gin.Context) {
	var obj model.OutGoingModel
	if err := c.ShouldBindJSON(&obj); err != nil {
		c.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()})
		return
	}

	if msg := OutgoingService.ExecDingCommand(obj); msg != nil {
		response.OkWithDetailed(string(msg), "创建成功", c)
	}
}
