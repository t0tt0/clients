package nsbi

import (
	uiptypes "github.com/HyperService-Consortium/go-uip/uiptypes"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
	"github.com/Myriad-Dreamin/go-ves/ves/control"
)

type NSBClientImpl struct {
	control.NSBClient
	signer uiptypes.Signer
}

func NSBInterfaceImpl(host string, signer uiptypes.Signer) *NSBClientImpl {
	return &NSBClientImpl{nsbcli.NewNSBClient(host), signer}
}

func NSBInterfaceFromClient(nsbClient control.NSBClient, signer uiptypes.Signer) *NSBClientImpl {
	return &NSBClientImpl{nsbClient, signer}
}

func (nsb *NSBClientImpl) SaveAttestation(iscAddress []byte, attestation uiptypes.Attestation) error {
	// todo
	return nil
}

func (nsb *NSBClientImpl) InsuranceClaim(iscAddress []byte, attestation uiptypes.Attestation) error {
	_, err := nsb.NSBClient.InsuranceClaim(nsb.signer, iscAddress, attestation.GetTid(), attestation.GetAid())
	return err
}

func (nsb *NSBClientImpl) SettleContract(iscAddress []byte) error {
	_, err := nsb.NSBClient.SettleContract(nsb.signer, iscAddress)
	return err
}
