package mock

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/golang/mock/gomock"
)

type matchSignature struct {
	s uip.Signature
}

func MatchSignature(s uip.Signature) gomock.Matcher {
	return matchSignature{s: s}
}

func (m matchSignature) Matches(x interface{}) bool {
	if s, ok := x.(uip.Signature); ok {
		return bytes.Equal(s.GetContent(), m.s.GetContent()) && s.GetSignatureType() == m.s.GetSignatureType()
	}
	return false
}

func (m matchSignature) String() string {
	return "match the properties on uip-types.Signature"
}
