package base_service

import (
	"github.com/Myriad-Dreamin/minimum-lib/controller"
)


type dHookObject interface {
	DeleteHook(c controller.MContext, obj interface{}) bool
}

type FObject interface {
	CreateEntity(id uint) interface{}
	GetEntity(id uint) (interface{}, error)
}

type DObject interface {
	dHookObject
}

type RObject interface {
	ResponseGet(c controller.MContext, obj interface{}) interface{}
	ResponseInspect(c controller.MContext, obj interface{}) interface{}
}

type UObject interface {
	GetPutRequest() interface{}
	FillPutFields(c controller.MContext, object interface{}, req interface{}) []string
}
type CObject interface {
	SerializePost(c controller.MContext) interface{}
	ResponsePost(obj interface{}) interface{}
}
