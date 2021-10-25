package nsbi

import (
	nsbcli "github.com/HyperService-Consortium/NSB/lib/nsb-client"
	uip "github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/ves/control"
)

type NSBClientImpl struct {
	control.NSBClient
	signer uip.Signer
}

func NSBInterfaceImpl(host string, signer uip.Signer) *NSBClientImpl {
	return &NSBClientImpl{nsbcli.NewNSBClient(host), signer}
}

func NSBInterfaceFromClient(nsbClient control.NSBClient, signer uip.Signer) *NSBClientImpl {
	return &NSBClientImpl{nsbClient, signer}
}

func (nsb *NSBClientImpl) SaveAttestation(iscAddress []byte, attestation uip.Attestation) error {
	// todo
	return nil
}

func (nsb *NSBClientImpl) InsuranceClaim(iscAddress []byte, attestation uip.Attestation) error {
	_, err := nsb.NSBClient.InsuranceClaim(nsb.signer, iscAddress, attestation.GetTid(), attestation.GetAid())
	return err
}

func (nsb *NSBClientImpl) SettleContract(iscAddress []byte) error {
	_, err := nsb.NSBClient.SettleContract(nsb.signer, iscAddress)
	return err
}

func (nsb *NSBClientImpl) ISCGetPC(iscAddress []byte) (uint64, error) {
	return nsb.NSBClient.ISCGetPC(nsb.signer, iscAddress)
}
