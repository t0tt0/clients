package dep_uip

import (
	"bytes"
	"github.com/HyperService-Consortium/go-uip/lib/serial"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type ContractMetaEncoder struct{}

func (*ContractMetaEncoder) Marshal(meta *opintent.ContractInvokeMeta) (_ []byte, err error) {
	var w = bytes.NewBuffer(nil)
	serial.Write(w, meta.Code, &err)
	serial.Write(w, meta.Meta, &err)
	serial.Write(w, meta.FuncName, &err)
	serial.Write(w, uint64(len(meta.Params)), &err)
	for i := range meta.Params {
		opintent.EncodeVTok(w, meta.Params[i], &err)
	}
	return w.Bytes(), err
}
func (*ContractMetaEncoder) Unmarshal(b []byte, meta *opintent.ContractInvokeMeta) (err error) {
	var r = bytes.NewReader(b)
	serial.Read(r, &meta.Code, &err)
	serial.Read(r, &meta.Meta, &err)
	serial.Read(r, &meta.FuncName, &err)
	var paramsLength uint64
	serial.Read(r, &paramsLength, &err)
	if err != nil {
		return
	}
	meta.Params = make([]uip.VTok, paramsLength)
	for i := range meta.Params {
		opintent.DecodeVTok(r, &meta.Params[i], &err)
	}
	return
}
