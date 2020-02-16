package mock

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/golang/mock/gomock"
)

type matchSignature struct {
	s uiptypes.Signature
}

func MatchSignature(s uiptypes.Signature) gomock.Matcher {
	return matchSignature{s: s}
}

func (m matchSignature) Matches(x interface{}) bool {
	if s, ok := x.(uiptypes.Signature); ok {
		return bytes.Equal(s.GetContent(), m.s.GetContent()) && s.GetSignatureType() == m.s.GetSignatureType()
	}
	return false
}

func (m matchSignature) String() string {
	return "match the properties on uip-types.Signature"
}
