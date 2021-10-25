package raw_transaction

import (
	"github.com/HyperService-Consortium/go-uip/uip"
)

func FromRaw(b []byte) uip.RawTransaction {
	return uip.RawTransactionImpl(b)
}
