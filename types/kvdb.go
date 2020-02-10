package types

type SessionKV interface {
	SetKV(iscAddress []byte, provedKey []byte, provedValue []byte) error
	GetKV(iscAddress []byte, provedKey []byte) (provedValue []byte, err error)
}

type KVSetter interface {
	Set(provedKey []byte, provedValue []byte) error
}

type KVGetter interface {
	Get([]byte) ([]byte, error)
}
