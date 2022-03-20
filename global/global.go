package global

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"zx/config"
)

var (
	//配置文件
	ZX_CONFIG config.Server
	ZX_VP     *viper.Viper
	//日志
	ZX_LOG *zap.Logger
)
