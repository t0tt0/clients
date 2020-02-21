package sessionservice

import (
	"github.com/HyperService-Consortium/go-ves/lib/backend/wrapper"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/HyperService-Consortium/go-ves/ves/model"
)

func (svc *Service) getSession(sessionID []byte) (*model.Session, error) {
	ses, err := svc.db.QueryGUIDByBytes(sessionID)
	if err != nil {
		return nil, wrapper.Wrap(types.CodeSessionFindError, err)
	} else if ses == nil {
		return nil, wrapper.WrapCode(types.CodeSessionNotFind)
	}
	return ses, nil
}
