package index

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/HyperService-Consortium/go-ves/types"
	"github.com/Myriad-Dreamin/gvm"
	"math/big"
)

type StorageHandler struct {
	index types.Index
}

func NewStorageHandler(i types.Index) *StorageHandler {
	return &StorageHandler{index: i}
}

func (g *StorageHandler) GetTransactionProof(chainID uip.ChainID, blockID uip.BlockID, color []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

type variable struct {
	Type  uip.TypeID
	Value interface{}
}

func (v variable) Encode() ([]byte, error) {
	panic("implement me")
}

func (v variable) GetGVMType() gvm.RefType {
	return gvm.RefType(v.Type)
}

func (v variable) Unwrap() interface{} {
	return v.Value
}

func (v variable) GetType() uip.TypeID {
	return v.Type
}

func (v variable) GetValue() interface{} {
	return v.Value
}

type Key struct {
	// todo underlying type
	ChainID         uip.ChainID
	ContractAddress uip.ContractAddress
	Pos             []byte
	Description     []byte
}

type KeyHeader struct {
	// todo underlying type
	ChainID         uip.ChainID
	ContractAddress uint64
	Pos             uint64
	Description     uint64
}

func (g *StorageHandler) GetStorageAt(
	chainID uip.ChainID,
	typeID uip.TypeID,
	contractAddress uip.ContractAddress,
	pos []byte,
	description []byte) (uip.Variable, error) {
	buf, err := toKey(chainID, contractAddress, pos, description)
	if err != nil {
		return nil, err
	}
	b, err := g.index.Get(buf.Bytes())
	if err != nil {
		return nil, err
	}
	var v variable
	v.Type = uip.TypeID(binary.BigEndian.Uint16(b[0:2]))

	// todo, convert?
	if v.Type != typeID {
		return nil, fmt.Errorf("unmatched type: provide %v, but %v", typeID, v.Type)
	}
	// todo: generalize 2
	buf = bytes.NewBuffer(b[2:])
	switch v.Type {
	case value_type.Bool:
		var val bool
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Uint256, value_type.Uint128, value_type.Int128, value_type.Int256:
		var val = big.NewInt(0).SetBytes(buf.Bytes())
		v.Value = val
	case value_type.Uint64:
		var val uint64
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Uint32:
		var val uint32
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Uint16:
		var val uint16
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Uint8:
		var val uint8
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Int64:
		var val int64
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Int32:
		var val int32
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Int16:
		var val int16
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Int8:
		var val int8
		err = binary.Read(buf, binary.BigEndian, &val)
		if err != nil {
			return nil, err
		}
		v.Value = val
	case value_type.Bytes:
		var val = buf.Bytes()
		v.Value = val
	case value_type.String:
		var val = buf.String()
		v.Value = val
	default:
		return nil, fmt.Errorf("not support this type id now: %v", v.Type)
	}
	return v, nil
}

func toKey(chainID uip.ChainID, contractAddress uip.ContractAddress, pos []byte, description []byte) (*bytes.Buffer, error) {
	buf := bytes.NewBufferString("k:")
	//todo normalize key
	err := binary.Write(buf, binary.BigEndian, KeyHeader{
		ChainID:         chainID,
		ContractAddress: uint64(len(contractAddress)),
		Pos:             uint64(len(pos)),
		Description:     uint64(len(description)),
	})
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(contractAddress)
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(pos)
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(description)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

func (g *StorageHandler) SetStorageOf(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte, val uip.Variable) error {

	buf, err := toKey(chainID, contractAddress, pos, description)
	if err != nil {
		return err
	}
	k := buf.Bytes()
	buf = bytes.NewBuffer(nil)
	err = binary.Write(buf, binary.BigEndian, val.GetType())
	if err != nil {
		return err
	}
	switch val.GetType() {
	// todo judge
	case value_type.String:
		_, err = buf.WriteString(val.GetValue().(string))
	case value_type.Bytes:
		_, err = buf.Write(val.GetValue().([]byte))
	case value_type.Uint64, value_type.Uint32, value_type.Uint16, value_type.Uint8,
		value_type.Int64, value_type.Int32, value_type.Int16, value_type.Int8,
		value_type.Bool:
		err = binary.Write(buf, binary.BigEndian, val.GetValue())
	case value_type.Uint128, value_type.Uint256, value_type.Int128, value_type.Int256:
		_, err = buf.Write(val.GetValue().(*big.Int).Bytes())
	default:
		return fmt.Errorf("not support this type id now: %v", val.GetType())
	}
	if err != nil {
		return err
	}

	return g.index.Set(k, buf.Bytes())
}

func clone(b []byte) []byte {
	var c = make([]byte, len(b))
	copy(c, b)
	return c
}

func cloneWithLen(b []byte, l int) []byte {
	var c = make([]byte, l)
	copy(c, b)
	return c
}
