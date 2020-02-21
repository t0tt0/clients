package base_service

import (
	"github.com/Myriad-Dreamin/dorm"
	dorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/dorm-crud-dao"
	ginhelper "github.com/HyperService-Consortium/go-ves/lib/backend/gin-helper"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
	"net/http"
)

type UObjectToolLite interface {
	FObject
	UObject
}

type UServiceInterface interface {
	Put(c controller.MContext)
}

type UService struct {
	Tool UObjectToolLite
	k    string
	UpdateFields func(obj dorm_crud_dao.ORMObject, fields []string) (int64, error)
}

func NewUService(tool UObjectToolLite, uf func(obj dorm_crud_dao.ORMObject, fields []string) (int64, error), k string) UService {
	return UService{
		Tool: tool,
		UpdateFields: uf,
		k:    k,
	}
}

func (srv UService) Put(c controller.MContext) {
	var req = srv.Tool.GetPutRequest()
	id, ok := ginhelper.ParseUintAndBind(c, srv.k, req)
	if !ok {
		return
	}

	object, err := srv.Tool.GetEntity(id)
	if ginhelper.MaybeSelectError(c, object, err) {
		return
	}

	fields := srv.Tool.FillPutFields(c, object, req)
	if c.IsAborted() {
		return
	}

	_, err = srv.UpdateFields(object.(dorm.ORMObject), fields)
	if ginhelper.UpdateFields(c, err) {
		c.JSON(http.StatusOK, &ginhelper.ResponseOK)
	}
}
