package objectservice

import (
	"github.com/HyperService-Consortium/go-ves/central-ves/control"
	"github.com/HyperService-Consortium/go-ves/central-ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) fillPutFields(
	c controller.MContext, object *model.Object,
	req *control.PutObjectRequest) (fields []string) {
	return nil
}
