package dingreceive

import (
	"zx/global"
)

func DingreceiveInit() {
	RegisterCommand("hello1", Hello1, 2, true)
	RegisterCommand("hello2", Hello2, 2, true)
	global.ZX_LOG.Info("命令注册成功")
}
