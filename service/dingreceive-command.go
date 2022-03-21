package service

import (
	"zx/global"
)

func Hello1(args []string) []byte {
	// do what you want to
	global.ZX_LOG.Info("Hello 接口1")
	for _, arg := range args {
		global.ZX_LOG.Info(arg)
	}
	return nil
}

func Hello2(args []string) []byte {
	// do what you want to
	global.ZX_LOG.Info("Hello 接口2")
	for _, arg := range args {
		global.ZX_LOG.Info(arg)
	}
	return nil
}
