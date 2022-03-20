package outgoing

import "zx/service"

type ApiGroup struct {
	OutgoingApi
}

var OutgoingService = service.ServiceGroupApp.OutgoingServiceGroup.OutgoingService
