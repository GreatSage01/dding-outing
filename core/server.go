package core

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"zx/global"
	"zx/initialize"
	"zx/service/outgoing"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {

	//outgoing 注册命令
	outgoing.OutgoingInit()
	//加载路由
	Router := initialize.Routers()

	address := fmt.Sprintf(":%d", global.ZX_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	global.ZX_LOG.Info("server run success on ", zap.String("address", address))
	global.ZX_LOG.Error(s.ListenAndServe().Error())
}
