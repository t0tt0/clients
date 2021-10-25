package base_service

import dorm_crud_dao "github.com/Myriad-Dreamin/go-model-traits/dorm-crud-dao"

type CRUDObjectToolLite interface {
	FObject
	DObject
	RObject
	UObject
	CObject
}

type CRUDService struct {
	Tool CRUDObjectToolLite
	k    string
	CServiceInterface
	RServiceInterface
	UServiceInterface
	DServiceInterface
}

type CRUDModel interface {
	Create(object interface{}) (aff int64, err error)
	Delete(object interface{}) (aff int64, err error)
	UpdateFields(object dorm_crud_dao.ORMObject, fields []string) (aff int64, err error)
}

func NewCRUDService(tool CRUDObjectToolLite, m CRUDModel, k string) CRUDService {
	return CRUDService{
		Tool:              tool,
		k:                 k,
		CServiceInterface: NewCService(tool, m.Create, k),
		RServiceInterface: NewRService(tool, k),
		UServiceInterface: NewUService(tool, m.UpdateFields, k),
		DServiceInterface: NewDService(tool, m.Delete, k),
	}
}
