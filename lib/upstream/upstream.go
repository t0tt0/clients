package upstream

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	opintent "github.com/HyperService-Consortium/go-uip/op-intent"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/tidwall/gjson"
	"math/big"
	"reflect"
	"testing"
)

type mcs struct{}
type ts struct {}
type _serializer struct {
	TransactionIntent ts
	Meta              struct {
		Contract mcs
	}
}

var Serializer = _serializer{}

func (mcs) Unmarshal(b []byte, meta *uip.ContractInvokeMeta) error {
	return json.Unmarshal(b, meta)
}

func (mcs) Marshal(meta *uip.ContractInvokeMeta) ([]byte, error) {
	return json.Marshal(meta)
}

func (ts) Unmarshal(b []byte, meta *opintent.TransactionIntent) error {
	return json.Unmarshal(b, meta)
}

func (ts) Marshal(meta *opintent.TransactionIntent) ([]byte, error) {
	return json.Marshal(meta)
}

type Kv struct {
	K string
	V interface{}
}

type GJSONAssertion struct {
	kvs []Kv
}

func (g GJSONAssertion) AssertBytes(object []byte) (err error) {
	for _, assertKeyValue := range g.kvs {
		k, v := assertKeyValue.K, assertKeyValue.V
		if err = g.compare(gjson.GetBytes(object, k), v); err != nil {
			return fmt.Errorf("compared failed on %v, assertion error %v", k, err)
		}
	}
	return
}

var int64T = reflect.TypeOf(int64(1))

func (g GJSONAssertion) compare(bytes gjson.Result, v interface{}) error {
	t := reflect.TypeOf(v)
	switch bytes.Type {
	case gjson.Null:
		if v != nil {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.False:
		if t.Kind() != reflect.Bool || v != false {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.True:
		if t.Kind() != reflect.Bool || v != true {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.Number:
		if !t.ConvertibleTo(int64T) ||
			reflect.ValueOf(v).Convert(int64T).Int() != bytes.Int() {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.String:
		if t.Kind() != reflect.String || v != bytes.String() {
			return fmt.Errorf("compare failed: %v %v", bytes, v)
		}
	case gjson.JSON:
		return fmt.Errorf("not basic comparable data")
	default:
		panic("unknown g-json type")
	}
	return nil
}

func GJSONWant(kvs ...Kv) GJSONAssertion {
	return GJSONAssertion{kvs: kvs}
}

type MockBNIStorage struct {
	Data []MockData
}

func (m MockBNIStorage) GetTransactionProof(chainID uip.ChainID, blockID uip.BlockID, color []byte) (uip.MerkleProof, error) {
	panic("implement me")
}

func (m MockBNIStorage) GetStorageAt(chainID uip.ChainID, typeID uip.TypeID, contractAddress uip.ContractAddress, pos []byte, description []byte) (uip.Variable, error) {
	for _, d := range m.Data {
		if d.ChainID == chainID && d.TypeID == typeID &&
			bytes.Equal(d.ContractAddress, contractAddress) &&
			bytes.Equal(d.Pos, pos) &&
			bytes.Equal(d.Description, description) {
			return d.V, nil
		}
	}
	return nil, errors.New("not found")
}

func (m *MockBNIStorage) insertMockData(data []MockData) {
	m.Data = append(m.Data, data...)
}

type testFunc = func(t *testing.T)

type bNIStorageTestSet struct {
	s uip.Storage
}

type MockData struct {
	ChainID         uip.ChainID
	TypeID          uip.TypeID
	ContractAddress uip.ContractAddress
	Pos             []byte
	Description     []byte
	V               uip.Variable
}

type MockValue struct {
	T value_type.Type
	V interface{}
}

func (m MockValue) GetType() uip.TypeID {
	return m.T
}

func (m MockValue) GetValue() interface{} {
	return m.V
}

var bytesType = reflect.TypeOf(new([]byte)).Elem()
var bigIntType = reflect.TypeOf(new(big.Int))

func StorageValueIsValid(v interface{}, t value_type.Type) bool {
	runtimeT := reflect.TypeOf(v)
	switch t {
	case value_type.Bytes,
		value_type.SliceInt8, value_type.SliceInt16, value_type.SliceInt32,
		value_type.SliceInt64, value_type.SliceInt128, value_type.SliceInt256,
		value_type.SliceUint8, value_type.SliceUint16, value_type.SliceUint32,
		value_type.SliceUint64, value_type.SliceUint128, value_type.SliceUint256:
		if v == nil {
			return true
		}
		switch t {
		case value_type.Bytes:
			if runtimeT != bytesType {
				return false
			}
		case value_type.SliceUint8:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint8 {
				return false
			}
		case value_type.SliceUint16:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint16 {
				return false
			}
		case value_type.SliceUint32:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint32 {
				return false
			}
		case value_type.SliceUint64:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Uint64 {
				return false
			}
		case value_type.SliceUint128:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		case value_type.SliceUint256:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		case value_type.SliceInt8:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int8 {
				return false
			}
		case value_type.SliceInt16:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int16 {
				return false
			}
		case value_type.SliceInt32:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int32 {
				return false
			}
		case value_type.SliceInt64:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem().Kind() != reflect.Int64 {
				return false
			}
		case value_type.SliceInt128:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		case value_type.SliceInt256:
			if runtimeT.Kind() != reflect.Slice || runtimeT.Elem() != bigIntType {
				return false
			}
		default:
			return false
		}
	case value_type.String, value_type.Bool, value_type.Uint128, value_type.Uint256, value_type.Int128, value_type.Int256,
		value_type.Int8, value_type.Int16, value_type.Int32, value_type.Int64,
		value_type.Uint8, value_type.Uint16, value_type.Uint32, value_type.Uint64:
		if v == nil {
			return false
		}
		switch t {
		case value_type.String:
			if runtimeT.Kind() != reflect.String {
				return false
			}
		case value_type.Uint128, value_type.Uint256, value_type.Int128, value_type.Int256:
			if runtimeT != bigIntType {
				return false
			}
		case value_type.Uint8:
			if runtimeT.Kind() != reflect.Uint8 {
				return false
			}
		case value_type.Uint16:
			if runtimeT.Kind() != reflect.Uint16 {
				return false
			}
		case value_type.Uint32:
			if runtimeT.Kind() != reflect.Uint32 {
				return false
			}
		case value_type.Uint64:
			if runtimeT.Kind() != reflect.Uint64 {
				return false
			}
		case value_type.Int8:
			if runtimeT.Kind() != reflect.Int8 {
				return false
			}
		case value_type.Int16:
			if runtimeT.Kind() != reflect.Int16 {
				return false
			}
		case value_type.Int32:
			if runtimeT.Kind() != reflect.Int32 {
				return false
			}
		case value_type.Int64:
			if runtimeT.Kind() != reflect.Int64 {
				return false
			}
		case value_type.Bool:
			if runtimeT.Kind() != reflect.Bool {
				return false
			}
		default:
			return false
		}
	default:
		return false
	}
	return true
}

func (m MockValue) IsValid() bool {
	return StorageValueIsValid(m.V, m.T)
}

func (b bNIStorageTestSet) MockingData() []MockData {
	return []MockData{
		{
			ChainID:         2,
			TypeID:          value_type.Int64,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int64, int64(10)},
		},
		{
			ChainID:         2,
			TypeID:          value_type.Int32,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int32, int32(11)},
		},
		{
			ChainID:         2,
			TypeID:          value_type.Int128,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int128, bigInt3},
		},
		{
			ChainID:         2,
			TypeID:          value_type.Int256,
			ContractAddress: make([]byte, 32),
			Pos:             make([]byte, 2),
			Description:     make([]byte, 2),
			V:               MockValue{value_type.Int256, bigInt3},
		},
	}
}

func (b bNIStorageTestSet) RunTests(t *testing.T) {
	t.Run("testGetInt32", b.testGetInt32)
	t.Run("testGetInt64", b.testGetInt64)
	t.Run("testGetInt128", b.testGetInt128)
	t.Run("testGetInt256", b.testGetInt256)
}

func assertType(l *testing.T, x uip.Variable, t value_type.Type, k reflect.Kind) bool {
	l.Helper()
	if x.GetType() != t {
		l.Fatal("bad type")
		return false
	}
	v0 := x.GetValue()
	v := reflect.ValueOf(v0)
	if v.Type().Kind() != k {
		l.Fatal("bad value type")
		return false
	}
	return true
}

func assertTypeOf(l *testing.T, x uip.Variable, t value_type.Type, r reflect.Type) bool {
	l.Helper()
	if x.GetType() != t {
		l.Fatal("bad type")
		return false
	}
	v0 := x.GetValue()
	v := reflect.ValueOf(v0)
	if v.Type() != r {
		l.Fatal("bad value type")
		return false
	}
	return true
}

func (b bNIStorageTestSet) testGetInt32(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int32, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertType(t, x, value_type.Int32, reflect.Int32) {
		return
	}
	if x.GetValue().(int32) != 11 {
		t.Fatal("bad value")
	}
}

func (b bNIStorageTestSet) testGetInt64(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int64, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertType(t, x, value_type.Int64, reflect.Int64) {
		return
	}
	if x.GetValue().(int64) != 10 {
		t.Fatal("bad value")
	}
}

var bigInt3 = big.NewInt(3)

func (b bNIStorageTestSet) testGetInt128(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int128, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertTypeOf(t, x, value_type.Int128, reflect.TypeOf(bigInt3)) {
		return
	}
	if x.GetValue().(*big.Int).Cmp(bigInt3) != 0 {
		t.Fatal("bad value")
	}
}

func (b bNIStorageTestSet) testGetInt256(t *testing.T) {
	x := sugar.HandlerError(b.s.GetStorageAt(2, value_type.Int256, make([]byte, 32), make([]byte, 2), make([]byte, 2))).(uip.Variable)
	if !assertTypeOf(t, x, value_type.Int256, reflect.TypeOf(bigInt3)) {
		return
	}
	if x.GetValue().(*big.Int).Cmp(bigInt3) != 0 {
		t.Fatal("bad value")
	}
}
