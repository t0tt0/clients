package bni

type MerkleProofInfo struct {
	Proof [][]byte `json:"proof"`
	Key   []byte   `json:"key"`
	Value []byte   `json:"value"`
}
