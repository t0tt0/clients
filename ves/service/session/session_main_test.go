package sessionservice

import (
	"encoding/json"
	"github.com/HyperService-Consortium/go-uip/const/value_type"
	"github.com/HyperService-Consortium/go-uip/uip"
	"github.com/Myriad-Dreamin/go-ves/lib/backend/wrapper"
	logger2 "github.com/Myriad-Dreamin/go-ves/lib/basic/log"
	"github.com/Myriad-Dreamin/go-ves/types"
	"github.com/Myriad-Dreamin/go-ves/ves/config"
	"github.com/Myriad-Dreamin/go-ves/ves/mock"
	"github.com/Myriad-Dreamin/minimum-lib/logger"
	"github.com/Myriad-Dreamin/minimum-lib/sugar"
	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"testing"
)

var describer = wrapper.Describer{
	Pack: "github.com/Myriad-Dreamin/go-ves/ves/service/session",
	Rel:  sugar.HandlerError(os.Getwd()).(string)}

var (
	sessionIDNotFound                   = []byte("xx")
	sessionIDFindTransactionError       = []byte("xy")
	sessionIDPushTransactionNotNil      = []byte("xz")
	sessionIDAttestationSendErrorNotOk  = []byte("yx")
	sessionIDFindError                  = []byte("yy")
	sessionIDAttestationSendErrorNotOk2 = []byte("yz")
	sessionIDOk                         = []byte("zx")
	sessionIDOk2                        = []byte("zy")
	sessionIDAttestationSendError       = []byte("zz")
	sessionIDGetBlockChainError         = []byte("xxx")

	sessionIDDeserializeTransactionError         = []byte("xyz")
	sessionIDFindSessionWithAcknowledgeError     = []byte("x")
	sessionIDFindSessionWithGetAcknowledgedError = []byte("y")
)

type ChainInfo struct {
	ChainType uip.ChainType
	ChainHost string
}

func (c ChainInfo) GetChainType() uip.ChainType {
	return c.ChainType
}

func (c ChainInfo) GetChainHost() string {
	return c.ChainHost
}

func MockSessionDB(ctl *gomock.Controller) *mock.SessionDB {
	return mock.NewSessionDB(ctl)
}

func MockSessionFSet(ctl *gomock.Controller) *mock.SessionFSet {
	return mock.NewSessionFSet(ctl)
}

func MockSessionAccountDB(ctl *gomock.Controller) *mock.SessionAccountDB {
	return mock.NewSessionAccountDB(ctl)
}

func MockCentralVESClient(ctl *gomock.Controller) *mock.CentralVESClient {
	return mock.NewCentralVESClient(ctl)
}

func MockChainDNS(ctl *gomock.Controller) *mock.ChainDNS {
	return mock.NewChainDNS(ctl)
}

func MockBlockChainInterface(ctl *gomock.Controller) *mock.BlockChainInterface {
	return mock.NewBlockChainInterface(ctl)
}

func MockStorageHandler(ctl *gomock.Controller) *mock.StorageHandler {
	return mock.NewStorageHandler(ctl)
}

func createService(options ...interface{}) *Service {
	ensureTestLogger()
	f := &Service{
		logger: testLogger,
		cfg:    config.Default(),
	}
	for i := range options {
		switch o := options[i].(type) {
		case *mock.SessionDB:
			f.db = o
		case *mock.SessionFSet:
			f.sesFSet = o
		case *mock.SessionAccountDB:
			f.accountDB = o
		case *mock.CentralVESClient:
			f.cVes = o
		case *mock.ChainDNS:
			f.dns = o
		case *mock.StorageHandler:
			f.storageHandler = o
		}
	}

	return f
}

func ensureTestLogger() {
	if testLogger == nil {
		if testing.Verbose() {
			var err error
			testLogger, err = logger.NewZapLogger(
				logger.NewZapDevelopmentSugarOption(), zapcore.DebugLevel)
			if err != nil {
				log.Fatal("init vesLogger error", "error", err)
			}
		} else {
			testLogger = logger2.NewNopLogger()
		}
	}
}

var testLogger logger.Logger

func checkErrorCode(t *testing.T, err error, i int) {
	t.Helper()
	if i != types.CodeOK {
		if f, ok := wrapper.FromError(err); ok {
			if f.GetCode() != i {
				t.Errorf("not expected code, error code %v, wantCode %v", f.GetCode(), i)
			} else {
				ensureTestLogger()
				testLogger.Info("expected good error", "error", describer.Describe(err))
			}
		} else {
			t.Error("not frame error wrapped")
		}
	}
}

func marshal(x interface{}) []byte {
	return sugar.HandlerError(json.Marshal(x)).([]byte)
}

func valueTypeToString(t value_type.Type) string {
	switch t {
	case value_type.Bytes:
		return "bytes"
	case value_type.String:
		return "string"
	case value_type.Uint8:
		return "uint8"
	case value_type.Uint16:
		return "uint16"
	case value_type.Uint32:
		return "uint32"
	case value_type.Uint64:
		return "uint64"
	case value_type.Uint128:
		return "uint128"
	case value_type.Uint256:
		return "uint256"
	case value_type.Int8:
		return "int8"
	case value_type.Int16:
		return "int16"
	case value_type.Int32:
		return "int32"
	case value_type.Int64:
		return "int64"
	case value_type.Int128:
		return "int128"
	case value_type.Int256:
		return "int256"
	case value_type.SliceUint8:
		return "[]uint8"
	case value_type.SliceUint16:
		return "[]uint16"
	case value_type.SliceUint32:
		return "[]uint32"
	case value_type.SliceUint64:
		return "[]uint64"
	case value_type.SliceUint128:
		return "[]uint128"
	case value_type.SliceUint256:
		return "[]uint256"
	case value_type.SliceInt8:
		return "[]int8"
	case value_type.SliceInt16:
		return "[]int16"
	case value_type.SliceInt32:
		return "[]int32"
	case value_type.SliceInt64:
		return "[]int64"
	case value_type.SliceInt128:
		return "[]int128"
	case value_type.SliceInt256:
		return "[]int256"
	case value_type.Bool:
		return "bool"
	default:
		return "<error-type>"
	}

}

func newRawMeta(t value_type.Type, value string) uip.RawParam {
	return uip.RawParam{
		Type: valueTypeToString(t),
		Value: marshal(map[string]interface{}{
			"constant": value,
		}),
	}
}

func newVarRawMeta(t value_type.Type, contract, pos, field string) uip.RawParam {
	return uip.RawParam{
		Type: valueTypeToString(t),
		Value: marshal(map[string]interface{}{
			"contract": contract,
			"pos":      pos,
			"field":    field,
		}),
	}
}
