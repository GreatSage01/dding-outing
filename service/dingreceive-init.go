package service

import (
	"zx/global"
	"zx/service/dingtalk"
)

func DingreceiveInit() {
	dingtalk.RegisterCommand("hello1", Hello1, 2, true)
	dingtalk.RegisterCommand("hello2", Hello2, 2, true)
	global.ZX_LOG.Info("命令注册成功")
}
