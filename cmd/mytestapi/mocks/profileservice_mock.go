// Code generated by MockGen. DO NOT EDIT.
// Source: mytestapi/cmd/mytestapi (interfaces: IProfileService)

// Package mocks is a generated GoMock package.
package mocks

import (
	mytestapi "mytestapi/cmd/mytestapi/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIProfileService is a mock of IProfileService interface.
type MockIProfileService struct {
	ctrl     *gomock.Controller
	recorder *MockIProfileServiceMockRecorder
}

// MockIProfileServiceMockRecorder is the mock recorder for MockIProfileService.
type MockIProfileServiceMockRecorder struct {
	mock *MockIProfileService
}

// NewMockIProfileService creates a new mock instance.
func NewMockIProfileService(ctrl *gomock.Controller) *MockIProfileService {
	mock := &MockIProfileService{ctrl: ctrl}
	mock.recorder = &MockIProfileServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIProfileService) EXPECT() *MockIProfileServiceMockRecorder {
	return m.recorder
}

// GetProfileByUsername mocks base method.
func (m *MockIProfileService) GetProfileByUsername(arg0 string) (*mytestapi.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfileByUsername", arg0)
	ret0, _ := ret[0].(*mytestapi.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfileByUsername indicates an expected call of GetProfileByUsername.
func (mr *MockIProfileServiceMockRecorder) GetProfileByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfileByUsername", reflect.TypeOf((*MockIProfileService)(nil).GetProfileByUsername), arg0)
}

// GetProfiles mocks base method.
func (m *MockIProfileService) GetProfiles() ([]mytestapi.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfiles")
	ret0, _ := ret[0].([]mytestapi.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfiles indicates an expected call of GetProfiles.
func (mr *MockIProfileServiceMockRecorder) GetProfiles() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfiles", reflect.TypeOf((*MockIProfileService)(nil).GetProfiles))
}
