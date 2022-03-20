package v1

import (
	"zx/api/v1/outgoing"
)

type ApiGroup struct {
	OutgoingApiGroup outgoing.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
