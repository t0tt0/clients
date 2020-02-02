package bni

import (
	"github.com/HyperService-Consortium/go-uip/uiptypes"
	"github.com/Myriad-Dreamin/go-ves/types"
)

type BN struct {
	dns    types.ChainDNSInterface
	signer uiptypes.Signer
}

func NewBN(dns types.ChainDNSInterface) *BN {
	return &BN{dns: dns}
}
