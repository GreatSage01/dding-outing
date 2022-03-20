package router

import (
	"zx/router/dingtalk"
)

type RouterGroup struct {
	DingtalkRouter dingtalk.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
