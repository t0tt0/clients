package abstraction

import (
	database2 "github.com/Myriad-Dreamin/go-ves/ves/model/internal/database"
)

type SessionDB interface {
	Create(s *database2.Session) (int64, error)
	Update(s *database2.Session) (int64, error)
	UpdateFields(s *database2.Session, fields []string) (int64, error)
	Delete(s *database2.Session) (int64, error)
	Filter(f *database2.SessionFilter) (user []database2.Session, err error)
	FilterI(f interface{}) (interface{}, error)
	ID(id uint) (session *database2.Session, err error)
	QueryGUID(iscAddress string) (session *database2.Session, err error)
	QueryGUIDByBytes(iscAddress []byte) (session *database2.Session, err error)

	Query(opts ...SessionQueryOption) (objs []database2.Session, err error)
	Scan(desc interface{}, opts ...SessionQueryOption) (err error)

}


type SessionQueryOption interface {
	implementsSessionQuery() SessionQueryOption
}

