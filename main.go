package main

import (
	"zx/core"
	"zx/global"
)

// 程序初始入口
func main() {
	global.ZX_VP = core.Viper() // 初始化Viper
	global.ZX_LOG = core.Zap()  // 初始化zap日志库
	//主程序
	core.RunWindowsServer()
}
