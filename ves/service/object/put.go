package objectservice

import (
	"github.com/HyperService-Consortium/go-ves/ves/control"
	"github.com/HyperService-Consortium/go-ves/ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) fillPutFields(
	c controller.MContext, object *model.Object,
	req *control.PutObjectRequest) (fields []string) {
	return nil
}
