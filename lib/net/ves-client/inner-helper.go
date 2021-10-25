package vesclient

import (
	"github.com/HyperService-Consortium/go-uip/uip"
	uiprpc_base "github.com/HyperService-Consortium/go-ves/grpc/uiprpc-base"
	"github.com/HyperService-Consortium/go-ves/lib/basic/encoding"
	"github.com/gogo/protobuf/proto"
)

func encodeAddress(src []byte) string {
	return encoding.EncodeHex(src)
}

func decodeAddress(src string) ([]byte, error) {
	return encoding.DecodeHex(src)
}

func encodeAddition(src []byte) string {
	return encoding.EncodeHex(src)
}

func decodeAddition(src string) ([]byte, error) {
	return encoding.DecodeHex(src)
}

func stringSliceToBytesSlice(ss []string) (bs [][]byte) {
	bs = make([][]byte, len(ss))
	for i := range ss {
		bs[i] = []byte(ss[i])
	}
	return
}

func (vc *VesClient) unmarshalProto(message []byte, target proto.Message) bool {
	err := proto.Unmarshal(message, target)
	if err != nil {
		vc.logger.Error("unmarshal protobuf error", "error", err)
	}
	return err == nil
}

func SignatureFromSTDToRPC(signature uip.Signature) *uiprpc_base.Signature {
	return &uiprpc_base.Signature{
		SignatureType: uip.SignatureTypeUnderlyingType(signature.GetSignatureType()),
		Content:       signature.GetContent(),
	}
}

// nextAttestation require oldAttestation not be nil
func nextAttestation(
	oldAttestation *uiprpc_base.Attestation, newSignature *uiprpc_base.Signature,
) *uiprpc_base.Attestation {
	return &uiprpc_base.Attestation{
		Tid:        oldAttestation.Tid,
		Aid:        oldAttestation.Aid + 1,
		Content:    oldAttestation.Content,
		Signatures: append(oldAttestation.Signatures, newSignature),
	}
}
