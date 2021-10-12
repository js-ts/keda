// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/scaling/scale_handler.go

// Package mock_scaling is a generated GoMock package.
package mock_scaling

import (
	"github.com/kedacore/keda/v2/pkg/scaling/cache"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	scalers "github.com/kedacore/keda/v2/pkg/scalers"
)

// MockScaleHandler is a mock of ScaleHandler interface.
type MockScaleHandler struct {
	ctrl     *gomock.Controller
	recorder *MockScaleHandlerMockRecorder
}

// MockScaleHandlerMockRecorder is the mock recorder for MockScaleHandler.
type MockScaleHandlerMockRecorder struct {
	mock *MockScaleHandler
}

// NewMockScaleHandler creates a new mock instance.
func NewMockScaleHandler(ctrl *gomock.Controller) *MockScaleHandler {
	mock := &MockScaleHandler{ctrl: ctrl}
	mock.recorder = &MockScaleHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScaleHandler) EXPECT() *MockScaleHandlerMockRecorder {
	return m.recorder
}

// DeleteScalableObject mocks base method.
func (m *MockScaleHandler) DeleteScalableObject(scalableObject interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteScalableObject", scalableObject)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteScalableObject indicates an expected call of DeleteScalableObject.
func (mr *MockScaleHandlerMockRecorder) DeleteScalableObject(scalableObject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteScalableObject", reflect.TypeOf((*MockScaleHandler)(nil).DeleteScalableObject), scalableObject)
}

// GetScalers mocks base method.
func (m *MockScaleHandler) GetScalers(scalableObject interface{}) ([]scalers.Scaler, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScalers", scalableObject)
	ret0, _ := ret[0].([]scalers.Scaler)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockScaleHandler) GetScalersCache(scalableObject interface{}) (*cache.ScalersCache, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetScalersCache", scalableObject)
	ret0, _ := ret[0].(*cache.ScalersCache)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (m *MockScaleHandler) ClearScalersCache(name, namespace, kind string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ClearScalersCache", name, namespace, kind)
}

// GetScalers indicates an expected call of GetScalers.
func (mr *MockScaleHandlerMockRecorder) GetScalers(scalableObject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetScalers", reflect.TypeOf((*MockScaleHandler)(nil).GetScalers), scalableObject)
}

// HandleScalableObject mocks base method.
func (m *MockScaleHandler) HandleScalableObject(scalableObject interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandleScalableObject", scalableObject)
	ret0, _ := ret[0].(error)
	return ret0
}

// HandleScalableObject indicates an expected call of HandleScalableObject.
func (mr *MockScaleHandlerMockRecorder) HandleScalableObject(scalableObject interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandleScalableObject", reflect.TypeOf((*MockScaleHandler)(nil).HandleScalableObject), scalableObject)
}
