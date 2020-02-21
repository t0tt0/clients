package types

import uip "github.com/HyperService-Consortium/go-uip/uip"

type isc_address = []byte

type NSBInterface interface {
	SaveAttestation(isc_address, uip.Attestation) error
	InsuranceClaim(isc_address, uip.Attestation) error
	SettleContract(isc_address) error
}

type success_or_not = bool
type help_info = string
type Session interface {
	// session must has isc_address or other guid

	// session is a kv-object
	KVObject

	SetSigner(uip.Signer)

	GetGUID() isc_address
	GetAccounts() []uip.Account
	GetTransaction(transaction_local_id) transaction
	GetTransactions() []transaction

	GetAckCount() uint32
	GetTransactingTransaction() (transaction_local_id, error)

	// error reports Internal errors, help_info reports Logic errors
	InitFromOpIntents(opIntents uip.OpIntents) error
	AckForInit(uip.Account, uip.Signature) (success_or_not, help_info, error)
	NotifyAttestation(NSBInterface, uip.BlockChainInterface, uip.Attestation) (success_or_not, help_info, error)
	ProcessAttestation(NSBInterface, uip.BlockChainInterface, uip.Attestation) (success_or_not, help_info, error)

	SyncFromISC() error
}

// the database which used by others

type transaction_id = uint64
type getter = func([]byte) error
type SessionBase interface {
	// insert accounts maps from guid to account
	InsertSessionInfo(MultiIndex, Index, Session) error

	// find accounts which guid is corresponding to user
	FindSessionInfo(MultiIndex, Index, isc_address) (Session, error)

	UpdateSessionInfo(MultiIndex, Index, Session) error

	DeleteSessionInfo(MultiIndex, Index, isc_address) error

	FindTransaction(Index, isc_address, transaction_id, getter) error

	ActivateSession(isc_address)
	InactivateSession(isc_address)
}
