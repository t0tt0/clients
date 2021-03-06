// Code generated by MockGen. DO NOT EDIT.
// Source: ../control/external-interface.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	nsb_message "github.com/HyperService-Consortium/NSB/lib/nsb-client/nsb-message"
	uip "github.com/HyperService-Consortium/go-uip/uip"
	uiprpc "github.com/HyperService-Consortium/go-ves/grpc/uiprpc"
	types "github.com/HyperService-Consortium/go-ves/types"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// ChainDNS is a mock of ChainDNS interface
type ChainDNS struct {
	ctrl     *gomock.Controller
	recorder *ChainDNSMockRecorder
}

// ChainDNSMockRecorder is the mock recorder for ChainDNS
type ChainDNSMockRecorder struct {
	mock *ChainDNS
}

// NewChainDNS creates a new mock instance
func NewChainDNS(ctrl *gomock.Controller) *ChainDNS {
	mock := &ChainDNS{ctrl: ctrl}
	mock.recorder = &ChainDNSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *ChainDNS) EXPECT() *ChainDNSMockRecorder {
	return m.recorder
}

// GetChainInfo mocks base method
func (m *ChainDNS) GetChainInfo(chainID uint64) (types.ChainInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainInfo", chainID)
	ret0, _ := ret[0].(types.ChainInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainInfo indicates an expected call of GetChainInfo
func (mr *ChainDNSMockRecorder) GetChainInfo(chainID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainInfo", reflect.TypeOf((*ChainDNS)(nil).GetChainInfo), chainID)
}

// CentralVESClient is a mock of CentralVESClient interface
type CentralVESClient struct {
	ctrl     *gomock.Controller
	recorder *CentralVESClientMockRecorder
}

// CentralVESClientMockRecorder is the mock recorder for CentralVESClient
type CentralVESClientMockRecorder struct {
	mock *CentralVESClient
}

// NewCentralVESClient creates a new mock instance
func NewCentralVESClient(ctrl *gomock.Controller) *CentralVESClient {
	mock := &CentralVESClient{ctrl: ctrl}
	mock.recorder = &CentralVESClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *CentralVESClient) EXPECT() *CentralVESClientMockRecorder {
	return m.recorder
}

// InternalRequestComing mocks base method
func (m *CentralVESClient) InternalRequestComing(ctx context.Context, in *uiprpc.InternalRequestComingRequest, opts ...grpc.CallOption) (*uiprpc.InternalRequestComingReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InternalRequestComing", varargs...)
	ret0, _ := ret[0].(*uiprpc.InternalRequestComingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InternalRequestComing indicates an expected call of InternalRequestComing
func (mr *CentralVESClientMockRecorder) InternalRequestComing(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalRequestComing", reflect.TypeOf((*CentralVESClient)(nil).InternalRequestComing), varargs...)
}

// InternalAttestationSending mocks base method
func (m *CentralVESClient) InternalAttestationSending(ctx context.Context, in *uiprpc.InternalRequestComingRequest, opts ...grpc.CallOption) (*uiprpc.InternalRequestComingReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InternalAttestationSending", varargs...)
	ret0, _ := ret[0].(*uiprpc.InternalRequestComingReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InternalAttestationSending indicates an expected call of InternalAttestationSending
func (mr *CentralVESClientMockRecorder) InternalAttestationSending(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalAttestationSending", reflect.TypeOf((*CentralVESClient)(nil).InternalAttestationSending), varargs...)
}

// InternalCloseSession mocks base method
func (m *CentralVESClient) InternalCloseSession(ctx context.Context, in *uiprpc.InternalCloseSessionRequest, opts ...grpc.CallOption) (*uiprpc.InternalCloseSessionReply, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "InternalCloseSession", varargs...)
	ret0, _ := ret[0].(*uiprpc.InternalCloseSessionReply)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InternalCloseSession indicates an expected call of InternalCloseSession
func (mr *CentralVESClientMockRecorder) InternalCloseSession(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InternalCloseSession", reflect.TypeOf((*CentralVESClient)(nil).InternalCloseSession), varargs...)
}

// NSBClient is a mock of NSBClient interface
type NSBClient struct {
	ctrl     *gomock.Controller
	recorder *NSBClientMockRecorder
}

// NSBClientMockRecorder is the mock recorder for NSBClient
type NSBClientMockRecorder struct {
	mock *NSBClient
}

// NewNSBClient creates a new mock instance
func NewNSBClient(ctrl *gomock.Controller) *NSBClient {
	mock := &NSBClient{ctrl: ctrl}
	mock.recorder = &NSBClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *NSBClient) EXPECT() *NSBClientMockRecorder {
	return m.recorder
}

// FreezeInfo mocks base method
func (m *NSBClient) FreezeInfo(signer uip.Signer, guid []byte, u uint64) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FreezeInfo", signer, guid, u)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FreezeInfo indicates an expected call of FreezeInfo
func (mr *NSBClientMockRecorder) FreezeInfo(signer, guid, u interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FreezeInfo", reflect.TypeOf((*NSBClient)(nil).FreezeInfo), signer, guid, u)
}



// InsuranceClaim mocks base method
func (m *NSBClient) InsuranceClaim(user uip.Signer, contractAddress []byte, tid, aid uint64) (*nsb_message.DeliverTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsuranceClaim", user, contractAddress, tid, aid)
	ret0, _ := ret[0].(*nsb_message.DeliverTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsuranceClaim indicates an expected call of InsuranceClaim
func (mr *NSBClientMockRecorder) InsuranceClaim(user, contractAddress, tid, aid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsuranceClaim", reflect.TypeOf((*NSBClient)(nil).InsuranceClaim), user, contractAddress, tid, aid)
}

// CreateISC mocks base method
func (m *NSBClient) CreateISC(signer uip.Signer, uint32s []uint32, bytes, bytes2 [][]byte, bytes3 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateISC", signer, uint32s, bytes, bytes2, bytes3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateISC indicates an expected call of CreateISC
func (mr *NSBClientMockRecorder) CreateISC(signer, uint32s, bytes, bytes2, bytes3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateISC", reflect.TypeOf((*NSBClient)(nil).CreateISC), signer, uint32s, bytes, bytes2, bytes3)
}

// SettleContract mocks base method
func (m *NSBClient) SettleContract(signer uip.Signer, bytes []byte) (*nsb_message.DeliverTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SettleContract", signer, bytes)
	ret0, _ := ret[0].(*nsb_message.DeliverTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SettleContract indicates an expected call of SettleContract
func (mr *NSBClientMockRecorder) SettleContract(signer, bytes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SettleContract", reflect.TypeOf((*NSBClient)(nil).SettleContract), signer, bytes)
}
