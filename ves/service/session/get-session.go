package sessionservice

import (
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/model"
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
