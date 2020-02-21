package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/go-ves/lib/backend/gin-helper"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

type CObjectToolLite interface {
	FObject
	CObject
}

type CServiceInterface interface {
	Post(c controller.MContext)
}

type CService struct {
	Tool       CObjectToolLite
	CreateFunc func(object interface{}) (aff int64, err error)
	k          string
}

func NewCService(tool CObjectToolLite, cf func(object interface{}) (aff int64, err error), k string) CService {
	return CService{
		Tool:       tool,
		CreateFunc: cf,
		k:          k,
	}
}

func (srv CService) Post(c controller.MContext) {
	var obj = srv.Tool.SerializePost(c)
	if c.IsAborted() {
		return
	}
	aff, err := srv.CreateFunc(obj)
	if ginhelper.CreateObj(c, aff, err) {
		c.JSON(http.StatusOK, srv.Tool.ResponsePost(obj))
	}
}
