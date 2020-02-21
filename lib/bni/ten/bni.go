package bni

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/types"
)

type BN struct {
	dns    types.ChainDNSInterface
	signer uip.Signer
}

func NewBN(dns types.ChainDNSInterface) *BN {
	return &BN{dns: dns}
}
