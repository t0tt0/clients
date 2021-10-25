package base_service

import (
	ginhelper "github.com/HyperService-Consortium/go-ves/lib/backend/gin-helper"
	"github.com/HyperService-Consortium/go-ves/lib/backend/serial"
	types2 "github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

type ListableObjectToolLite interface {
	CreateFilter() interface{}
	ProcessListResults(c controller.MContext, r interface{}) interface{}
}

type ListService struct {
	tool       ListableObjectToolLite
	filterFunc FilterFunc
}

func NewListService(tool ListableObjectToolLite, filterFunc FilterFunc) ListService {
	return ListService{
		tool:       tool,
		filterFunc: filterFunc,
	}
}

type FilterFunc = func(f interface{}) (interface{}, error)

type ListReply struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
}

func (srv *ListService) List(c controller.MContext) {
	var f = srv.tool.CreateFilter()
	if !ginhelper.BindRequest(c, f) {
		return
	}
	result, err := srv.filterFunc(f)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, serial.ErrorSerializer{
			Code: types2.CodeSelectError,
			Err:  err.Error(),
		})
		return
	}

	result = srv.tool.ProcessListResults(c, result)
	if !c.IsAborted() {
		c.JSON(http.StatusOK, result)
	}
	return
}
