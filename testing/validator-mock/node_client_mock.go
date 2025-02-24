// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/prysmaticlabs/prysm/v5/validator/client/iface (interfaces: NodeClient)
//
// Generated by this command:
//
//	mockgen -package=validator_mock -destination=testing/validator-mock/node_client_mock.go github.com/prysmaticlabs/prysm/v5/validator/client/iface NodeClient
//

// Package validator_mock is a generated GoMock package.
package validator_mock

import (
	context "context"
	reflect "reflect"

	"github.com/prysmaticlabs/prysm/v5/api/client/beacon"
	eth "github.com/prysmaticlabs/prysm/v5/proto/prysm/v1alpha1"
	gomock "go.uber.org/mock/gomock"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// MockNodeClient is a mock of NodeClient interface.
type MockNodeClient struct {
	ctrl          *gomock.Controller
	recorder      *MockNodeClientMockRecorder
	healthTracker *beacon.NodeHealthTracker
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient.
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance.
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	mock.healthTracker = beacon.NewNodeHealthTracker(mock)
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// GetGenesis mocks base method.
func (m *MockNodeClient) GetGenesis(arg0 context.Context, arg1 *emptypb.Empty) (*eth.Genesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesis", arg0, arg1)
	ret0, _ := ret[0].(*eth.Genesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesis indicates an expected call of GetGenesis.
func (mr *MockNodeClientMockRecorder) GetGenesis(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesis", reflect.TypeOf((*MockNodeClient)(nil).GetGenesis), arg0, arg1)
}

// GetSyncStatus mocks base method.
func (m *MockNodeClient) GetSyncStatus(arg0 context.Context, arg1 *emptypb.Empty) (*eth.SyncStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncStatus", arg0, arg1)
	ret0, _ := ret[0].(*eth.SyncStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncStatus indicates an expected call of GetSyncStatus.
func (mr *MockNodeClientMockRecorder) GetSyncStatus(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncStatus", reflect.TypeOf((*MockNodeClient)(nil).GetSyncStatus), arg0, arg1)
}

// GetVersion mocks base method.
func (m *MockNodeClient) GetVersion(arg0 context.Context, arg1 *emptypb.Empty) (*eth.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0, arg1)
	ret0, _ := ret[0].(*eth.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion.
func (mr *MockNodeClientMockRecorder) GetVersion(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockNodeClient)(nil).GetVersion), arg0, arg1)
}

// IsHealthy mocks base method.
func (m *MockNodeClient) IsHealthy(arg0 context.Context) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsHealthy", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsHealthy indicates an expected call of IsHealthy.
func (mr *MockNodeClientMockRecorder) IsHealthy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsHealthy", reflect.TypeOf((*MockNodeClient)(nil).IsHealthy), arg0)
}

// ListPeers mocks base method.
func (m *MockNodeClient) ListPeers(arg0 context.Context, arg1 *emptypb.Empty) (*eth.Peers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPeers", arg0, arg1)
	ret0, _ := ret[0].(*eth.Peers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPeers indicates an expected call of ListPeers.
func (mr *MockNodeClientMockRecorder) ListPeers(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPeers", reflect.TypeOf((*MockNodeClient)(nil).ListPeers), arg0, arg1)
}

func (m *MockNodeClient) HealthTracker() *beacon.NodeHealthTracker {
	return m.healthTracker
}
