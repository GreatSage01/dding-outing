package v1

import (
	"zx/api/v1/dingreceive"
)

type ApiGroup struct {
	DingreceiveApiGroup dingreceive.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
