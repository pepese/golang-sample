// Code generated by MockGen. DO NOT EDIT.
// Source: app/sample.go

// Package mock_app is a generated GoMock package.
package mock_app

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSampleInterfacer is a mock of SampleInterfacer interface.
type MockSampleInterfacer struct {
	ctrl     *gomock.Controller
	recorder *MockSampleInterfacerMockRecorder
}

// MockSampleInterfacerMockRecorder is the mock recorder for MockSampleInterfacer.
type MockSampleInterfacerMockRecorder struct {
	mock *MockSampleInterfacer
}

// NewMockSampleInterfacer creates a new mock instance.
func NewMockSampleInterfacer(ctrl *gomock.Controller) *MockSampleInterfacer {
	mock := &MockSampleInterfacer{ctrl: ctrl}
	mock.recorder = &MockSampleInterfacerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSampleInterfacer) EXPECT() *MockSampleInterfacerMockRecorder {
	return m.recorder
}

// SampleFunc mocks base method.
func (m *MockSampleInterfacer) SampleFunc(in int) int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SampleFunc", in)
	ret0, _ := ret[0].(int)
	return ret0
}

// SampleFunc indicates an expected call of SampleFunc.
func (mr *MockSampleInterfacerMockRecorder) SampleFunc(in interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SampleFunc", reflect.TypeOf((*MockSampleInterfacer)(nil).SampleFunc), in)
}
