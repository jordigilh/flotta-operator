// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/project-flotta/flotta-operator/internal/edgeapi/backend/k8s (interfaces: RepositoryFacade)

// Package k8s is a generated GoMock package.
package k8s

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
	client "sigs.k8s.io/controller-runtime/pkg/client"

	v1alpha1 "github.com/project-flotta/flotta-operator/api/v1alpha1"
)

// MockRepositoryFacade is a mock of RepositoryFacade interface.
type MockRepositoryFacade struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryFacadeMockRecorder
}

// MockRepositoryFacadeMockRecorder is the mock recorder for MockRepositoryFacade.
type MockRepositoryFacadeMockRecorder struct {
	mock *MockRepositoryFacade
}

// NewMockRepositoryFacade creates a new mock instance.
func NewMockRepositoryFacade(ctrl *gomock.Controller) *MockRepositoryFacade {
	mock := &MockRepositoryFacade{ctrl: ctrl}
	mock.recorder = &MockRepositoryFacadeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryFacade) EXPECT() *MockRepositoryFacadeMockRecorder {
	return m.recorder
}

// CreateEdgeDeviceSignedRequest mocks base method.
func (m *MockRepositoryFacade) CreateEdgeDeviceSignedRequest(arg0 context.Context, arg1 *v1alpha1.EdgeDeviceSignedRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateEdgeDeviceSignedRequest", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateEdgeDeviceSignedRequest indicates an expected call of CreateEdgeDeviceSignedRequest.
func (mr *MockRepositoryFacadeMockRecorder) CreateEdgeDeviceSignedRequest(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEdgeDeviceSignedRequest", reflect.TypeOf((*MockRepositoryFacade)(nil).CreateEdgeDeviceSignedRequest), arg0, arg1)
}

// GetConfigMap mocks base method.
func (m *MockRepositoryFacade) GetConfigMap(arg0 context.Context, arg1, arg2 string) (*v1.ConfigMap, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfigMap", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.ConfigMap)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConfigMap indicates an expected call of GetConfigMap.
func (mr *MockRepositoryFacadeMockRecorder) GetConfigMap(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfigMap", reflect.TypeOf((*MockRepositoryFacade)(nil).GetConfigMap), arg0, arg1, arg2)
}

// GetEdgeDevice mocks base method.
func (m *MockRepositoryFacade) GetEdgeDevice(arg0 context.Context, arg1, arg2 string) (*v1alpha1.EdgeDevice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEdgeDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.EdgeDevice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEdgeDevice indicates an expected call of GetEdgeDevice.
func (mr *MockRepositoryFacadeMockRecorder) GetEdgeDevice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEdgeDevice", reflect.TypeOf((*MockRepositoryFacade)(nil).GetEdgeDevice), arg0, arg1, arg2)
}

// GetEdgeDeviceSet mocks base method.
func (m *MockRepositoryFacade) GetEdgeDeviceSet(arg0 context.Context, arg1, arg2 string) (*v1alpha1.EdgeDeviceSet, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEdgeDeviceSet", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.EdgeDeviceSet)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEdgeDeviceSet indicates an expected call of GetEdgeDeviceSet.
func (mr *MockRepositoryFacadeMockRecorder) GetEdgeDeviceSet(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEdgeDeviceSet", reflect.TypeOf((*MockRepositoryFacade)(nil).GetEdgeDeviceSet), arg0, arg1, arg2)
}

// GetEdgeDeviceSignedRequest mocks base method.
func (m *MockRepositoryFacade) GetEdgeDeviceSignedRequest(arg0 context.Context, arg1, arg2 string) (*v1alpha1.EdgeDeviceSignedRequest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEdgeDeviceSignedRequest", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.EdgeDeviceSignedRequest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEdgeDeviceSignedRequest indicates an expected call of GetEdgeDeviceSignedRequest.
func (mr *MockRepositoryFacadeMockRecorder) GetEdgeDeviceSignedRequest(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEdgeDeviceSignedRequest", reflect.TypeOf((*MockRepositoryFacade)(nil).GetEdgeDeviceSignedRequest), arg0, arg1, arg2)
}

// GetEdgeWorkload mocks base method.
func (m *MockRepositoryFacade) GetEdgeWorkload(arg0 context.Context, arg1, arg2 string) (*v1alpha1.EdgeWorkload, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEdgeWorkload", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.EdgeWorkload)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEdgeWorkload indicates an expected call of GetEdgeWorkload.
func (mr *MockRepositoryFacadeMockRecorder) GetEdgeWorkload(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEdgeWorkload", reflect.TypeOf((*MockRepositoryFacade)(nil).GetEdgeWorkload), arg0, arg1, arg2)
}

// GetPlaybookExecution mocks base method.
func (m *MockRepositoryFacade) GetPlaybookExecution(arg0 context.Context, arg1, arg2 string) (*v1alpha1.PlaybookExecution, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPlaybookExecution", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1alpha1.PlaybookExecution)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPlaybookExecution indicates an expected call of GetPlaybookExecution.
func (mr *MockRepositoryFacadeMockRecorder) GetPlaybookExecution(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPlaybookExecution", reflect.TypeOf((*MockRepositoryFacade)(nil).GetPlaybookExecution), arg0, arg1, arg2)
}

// GetSecret mocks base method.
func (m *MockRepositoryFacade) GetSecret(arg0 context.Context, arg1, arg2 string) (*v1.Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecret", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecret indicates an expected call of GetSecret.
func (mr *MockRepositoryFacadeMockRecorder) GetSecret(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecret", reflect.TypeOf((*MockRepositoryFacade)(nil).GetSecret), arg0, arg1, arg2)
}

// PatchEdgeDevice mocks base method.
func (m *MockRepositoryFacade) PatchEdgeDevice(arg0 context.Context, arg1, arg2 *v1alpha1.EdgeDevice) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchEdgeDevice", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEdgeDevice indicates an expected call of PatchEdgeDevice.
func (mr *MockRepositoryFacadeMockRecorder) PatchEdgeDevice(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEdgeDevice", reflect.TypeOf((*MockRepositoryFacade)(nil).PatchEdgeDevice), arg0, arg1, arg2)
}

// PatchEdgeDeviceStatus mocks base method.
func (m *MockRepositoryFacade) PatchEdgeDeviceStatus(arg0 context.Context, arg1 *v1alpha1.EdgeDevice, arg2 *client.Patch) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PatchEdgeDeviceStatus", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// PatchEdgeDeviceStatus indicates an expected call of PatchEdgeDeviceStatus.
func (mr *MockRepositoryFacadeMockRecorder) PatchEdgeDeviceStatus(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PatchEdgeDeviceStatus", reflect.TypeOf((*MockRepositoryFacade)(nil).PatchEdgeDeviceStatus), arg0, arg1, arg2)
}

// UpdateEdgeDeviceLabels mocks base method.
func (m *MockRepositoryFacade) UpdateEdgeDeviceLabels(arg0 context.Context, arg1 *v1alpha1.EdgeDevice, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateEdgeDeviceLabels", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateEdgeDeviceLabels indicates an expected call of UpdateEdgeDeviceLabels.
func (mr *MockRepositoryFacadeMockRecorder) UpdateEdgeDeviceLabels(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEdgeDeviceLabels", reflect.TypeOf((*MockRepositoryFacade)(nil).UpdateEdgeDeviceLabels), arg0, arg1, arg2)
}
