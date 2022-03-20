package service

import (
	"zx/service/outgoing"
)

type ServiceGroup struct {
	OutgoingServiceGroup outgoing.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
