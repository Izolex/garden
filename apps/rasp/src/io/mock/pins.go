// Code generated by MockGen. DO NOT EDIT.
// Source: pins.go

// Package mock is a generated GoMock package.
package mock

import (
	io "main/io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPins is a mock of Pins interface.
type MockPins struct {
	ctrl     *gomock.Controller
	recorder *MockPinsMockRecorder
}

// MockPinsMockRecorder is the mock recorder for MockPins.
type MockPinsMockRecorder struct {
	mock *MockPins
}

// NewMockPins creates a new mock instance.
func NewMockPins(ctrl *gomock.Controller) *MockPins {
	mock := &MockPins{ctrl: ctrl}
	mock.recorder = &MockPinsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPins) EXPECT() *MockPinsMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockPins) Get(name string) io.Pin {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", name)
	ret0, _ := ret[0].(io.Pin)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockPinsMockRecorder) Get(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPins)(nil).Get), name)
}
