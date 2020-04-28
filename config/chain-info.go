package config

import (
	"encoding/hex"
	"errors"
	ChainType "github.com/HyperService-Consortium/go-uip/const/chain_type"

	merkleprooftype "github.com/HyperService-Consortium/go-uip/const/merkle-proof-type"
	"github.com/HyperService-Consortium/go-uip/uip"
)

type ChainInfo struct {
	Host      string
	ChainType ChainType.Type
}

func (c *ChainInfo) GetHost() string {
	return c.Host
}

func (c *ChainInfo) GetChainType() ChainType.Type {
	return c.ChainType
}

func getRelay(domain uint64) (uip.Account, error) {
	switch domain {
	case 0:
		return nil, errors.New("nil domain is not allowed")
	case 1: // ethereum chain 1
		b, err := hex.DecodeString("0ac45f1e6b8d47ac4c73aee62c52794b5898da9f")
		return &uip.AccountImpl{
			ChainId: domain,
			Address: b,
		}, err
	case 2: // ethereum chain 2
		b, err := hex.DecodeString("d051a43d3ea62afff3632bca3d5abf68bc6fd737")
		return &uip.AccountImpl{
			ChainId: domain,
			Address: b,
		}, err
	case 3: // tendermint chain 1
		//2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff
		b, err := hex.DecodeString("2333eeffffffffffffff2333eeffffffffffffff2333eeffffffffffffffffff")
		return &uip.AccountImpl{
			ChainId: domain,
			Address: b,
		}, err
	case 4: // tendermint chain 2
		b, err := hex.DecodeString("2333bbffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff")
		return &uip.AccountImpl{
			ChainId: domain,
			Address: b,
		}, err
	case 6:
		b, err := hex.DecodeString("5539f2a34aed86f9032c908bf404eee3b0bdbddc")
		return &uip.AccountImpl{
			ChainId: domain,
			Address: b,
		}, err
	case 7:
		b, err := hex.DecodeString("0201a1327c8e7ee9a8f1611913f7478f368b8b14")
		return &uip.AccountImpl{
			ChainId: domain,
			Address: b,
		}, err
	case 8:
		b, err := hex.DecodeString("0201a1327c8e7ee9a8f1611913f7478f368b8b14")
		return &uip.AccountImpl{
			ChainId: domain,
			Address: b,
		}, err
	default:
		return nil, errors.New("not found")
	}
}

func searchAccount(name string, chainId uint64) (uip.Account, error) {
	switch name {
	case "":
		return nil, errors.New("nil name is not allowed")
	case "a1": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 1: // ethereum chain 1
			b, err := hex.DecodeString("d051a43d3ea62afff3632bca3d5abf68bc6fd737")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 2: // ethereum chain 2
			b, err := hex.DecodeString("93334ae4b2d42ebba8cc7c797bfeb02bfb3349d6")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("604bdd2dd4b7e1b761e2ac96db99bb2bda386bb0d075b51a8f49c5103ebaa985")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 4: // ethereum chain 1
			b, err := hex.DecodeString("604bdd2dd4b7e1b761e2ac96db99bb2bda386bb0d075b51a8f49c5103ebaa985")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 6:
			b, err := hex.DecodeString("5539f2a34aed86f9032c908bf404eee3b0bdbddc")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 7:
			b, err := hex.DecodeString("4b3a59cd1183ab81b3c31b5a22bce26adf928ac2")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 8:
			b, err := hex.DecodeString("4b3a59cd1183ab81b3c31b5a22bce26adf928ac2")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	case "a2": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 1: // ethereum chain 1
			b, err := hex.DecodeString("47a1cdb6594d6efed3a6b917f2fbaa2bbcf61a2e")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 2: // ethereum chain 2
			b, err := hex.DecodeString("981739a13593980763de3353340617ef16da6354")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("2333eeffffffffffffff2333eeffffffffffffff2333eeffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
			//b, err := hex.DecodeString("cfe900c7a56f87882f0e18e26851bce7b7e61ebeca6c4b235fa360d627dfac63")
			//return &uip.AccountImpl{
			//	ChainId: chainId,
			//	Address: b,
			//}, err
		case 4: // tendermint chain 1
			b, err := hex.DecodeString("2333eeffffffffffffff2333eeffffffffffffff2333eeffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
			//b, err := hex.DecodeString("4f7a1b3d9f2f8f3e2c7e7729bc873fc55e607e47309941391a7a82673e563887")
			//return &uip.AccountImpl{
			//	ChainId: chainId,
			//	Address: b,
			//}, err
		case 6: // ethereum chain 1
			b, err := hex.DecodeString("2b5680581553c2312dba96cb8d7639cc049cece7")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 7: // ethereum chain 3
			b, err := hex.DecodeString("6bce60cc3c882ccc7da13876583a4064eb6c04c9")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 8: // ethereum chain 4
			b, err := hex.DecodeString("6bce60cc3c882ccc7da13876583a4064eb6c04c9")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	case "a3": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("2333eeffffffffffffff2333eeffffffffffffff2333eeffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 4: // tendermint chain 1
			b, err := hex.DecodeString("2333eeffffffffffffff2333eeffffffffffffff2333eeffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	case "a4": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("2333ffffffffffffffff2333ffffffffffffffff2333ffffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 4: // tendermint chain 1
			b, err := hex.DecodeString("2333ffffffffffffffff2333ffffffffffffffff2333ffffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	case "a5": // ethereum chain 1
		switch chainId {
		case 0:
			return nil, errors.New("nil domain is not allowed")
		case 3: // tendermint chain 1
			b, err := hex.DecodeString("2333bfffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		case 4: // tendermint chain 1
			b, err := hex.DecodeString("2333bfffffffffffffff2333bbffffffffffffff2333bbffffffffffffffffff")
			return &uip.AccountImpl{
				ChainId: chainId,
				Address: b,
			}, err
		default:
			return nil, errors.New("not found")
		}
	default:
		return nil, errors.New("not registered uip")
	}
}

func getTransactionProofType(chainId uint64) (uip.MerkleProofType, error) {
	// ethereum chain 1
	switch chainId {
	case 0:
		return merkleprooftype.Invalid, errors.New("nil domain is not allowed")
	case 1: // ethereum chain 1
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	case 2: // ethereum chain 2
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	case 3: // tendermint chain 1
		return merkleprooftype.MerklePatriciaTrieUsingKeccak256, nil
	case 4: // tendermint chain 2
		return merkleprooftype.MerklePatriciaTrieUsingKeccak256, nil
	case 5: // ethereum chain 3
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	case 6: // ethereum chain 4
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	case 7: // ethereum chain 5
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	case 8: // ethereum chain 6
		return merkleprooftype.SecureMerklePatriciaTrieUsingKeccak256, nil
	default:
		return merkleprooftype.Invalid, errors.New("not found")
	}
}
