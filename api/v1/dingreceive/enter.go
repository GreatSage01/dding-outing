package dingreceive

import "zx/service"

type ApiGroup struct {
	DingreceiveApi
}

var DingreceiveService = service.ServiceGroupApp.DingreceiveServiceGroup
