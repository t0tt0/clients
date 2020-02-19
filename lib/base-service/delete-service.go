package base_service

import (
	ginhelper "github.com/Myriad-Dreamin/go-ves/lib/gin-helper"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

type DObjectToolLite interface {
	FObject
	DObject
}

type DServiceInterface interface {
	Delete(c controller.MContext)
}

type DService struct {
	Tool       DObjectToolLite
	k          string
	DeleteFunc func(object interface{}) (aff int64, err error)
}

func NewDService(tool DObjectToolLite, df func(object interface{}) (aff int64, err error), k string) DService {
	return DService{
		Tool:       tool,
		DeleteFunc: df,
		k:          k,
	}
}

func (srv DService) Delete(c controller.MContext) {
	id, ok := ginhelper.ParseUint(c, srv.k)
	if !ok {
		return
	}
	obj, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, obj, err) {
		return
	}
	if !srv.Tool.DeleteHook(c, obj) {
		return
	}

	aff, err := srv.DeleteFunc(obj)
	if ginhelper.DeleteObj(c, aff, err) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	} else {
	}
}
