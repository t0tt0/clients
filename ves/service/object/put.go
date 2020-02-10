package objectservice

import (
	"github.com/Myriad-Dreamin/go-ves/ves/control"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)

func (svc *Service) fillPutFields(
	c controller.MContext, object *model.Object,
	req *control.PutObjectRequest) (fields []string) {
	return nil
}
