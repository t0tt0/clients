package bni

import (
	"encoding/json"
	"errors"
	merkleproof "github.com/HyperService-Consortium/go-uip/merkle-proof"
	"github.com/HyperService-Consortium/go-uip/uip"
	nsbcli "github.com/Myriad-Dreamin/go-ves/lib/net/nsb-client"
)

func (bn *BN) GetTransactionProof(chainID uint64, blockID []byte, additional []byte) (uip.MerkleProof, error) {
	ci, err := bn.dns.GetChainInfo(chainID)

	if err != nil {
		return nil, err
	}

	resp, err := nsbcli.NewNSBClient(ci.GetChainHost()).GetProof(additional, `"prove_on_tx_trie"`)

	if err != nil {
		return nil, err
	}

	var info MerkleProofInfo
	err = json.Unmarshal([]byte(resp.Info), &info)
	if err != nil {
		return nil, err
	}
	return merkleproof.NewMPTUsingKeccak256(info.Proof, info.Key, info.Value), nil
}

func (bn *BN) GetStorageAt(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	return nil, errors.New("todo")
	//ci, err := bn.dns.GetChainInfo(chainID)
	//if err != nil {
	//	return nil, err
	//}
	//
	//switch typeID {
	//case valuetype.Bool:
	//	s, err := ethclient.NewEthClient((&url.URL{Scheme: "http", Host: ci.GetChainHost(), Path: "/"}).String()).GetStorageAt(contractAddress, pos, "latest")
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	b, err := hex.DecodeString(s[2:])
	//	if err != nil {
	//		return nil, err
	//	}
	//	bs, err := ethabi.NewDecoder().Decodes([]string{"bool"}, b)
	//	return bs[0], nil
	//case valuetype.Uint256:
	//	s, err := ethclient.NewEthClient(ci.GetChainHost()).GetStorageAt(contractAddress, pos, "latest")
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	b, err := hex.DecodeString(s[2:])
	//	if err != nil {
	//		return nil, err
	//	}
	//	bs, err := ethabi.NewDecoder().Decodes([]string{"uint256"}, b)
	//	return bs[0], nil
	//}

	//return nil, nil
}
