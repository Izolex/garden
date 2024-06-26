// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	measurement "shared/model/entity/measurement"

	gomock "github.com/golang/mock/gomock"
)

// MockInserter is a mock of Inserter interface.
type MockInserter struct {
	ctrl     *gomock.Controller
	recorder *MockInserterMockRecorder
}

// MockInserterMockRecorder is the mock recorder for MockInserter.
type MockInserterMockRecorder struct {
	mock *MockInserter
}

// NewMockInserter creates a new mock instance.
func NewMockInserter(ctrl *gomock.Controller) *MockInserter {
	mock := &MockInserter{ctrl: ctrl}
	mock.recorder = &MockInserterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInserter) EXPECT() *MockInserterMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockInserter) Insert(measurement *measurement.Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", measurement)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockInserterMockRecorder) Insert(measurement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockInserter)(nil).Insert), measurement)
}

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// Insert mocks base method.
func (m *MockStorage) Insert(measurement *measurement.Entity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", measurement)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockStorageMockRecorder) Insert(measurement interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockStorage)(nil).Insert), measurement)
}
