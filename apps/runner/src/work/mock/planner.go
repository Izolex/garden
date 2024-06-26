// Code generated by MockGen. DO NOT EDIT.
// Source: planner.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPlanner is a mock of Planner interface.
type MockPlanner struct {
	ctrl     *gomock.Controller
	recorder *MockPlannerMockRecorder
}

// MockPlannerMockRecorder is the mock recorder for MockPlanner.
type MockPlannerMockRecorder struct {
	mock *MockPlanner
}

// NewMockPlanner creates a new mock instance.
func NewMockPlanner(ctrl *gomock.Controller) *MockPlanner {
	mock := &MockPlanner{ctrl: ctrl}
	mock.recorder = &MockPlannerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlanner) EXPECT() *MockPlannerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockPlanner) Run() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Run")
}

// Run indicates an expected call of Run.
func (mr *MockPlannerMockRecorder) Run() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockPlanner)(nil).Run))
}
