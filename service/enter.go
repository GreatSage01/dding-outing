package service

import (
	"zx/service/dingreceive"
)

type ServiceGroup struct {
	DingreceiveServiceGroup dingreceive.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
