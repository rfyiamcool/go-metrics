package mock

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	"github.com/rfyiamcool/go-metrics"
	grpc "google.golang.org/grpc"
)

// MockMetrics is a mock of Metrics interface
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// GetGrpcServerOption mocks base method
func (m *MockMetrics) GetGrpcServerOption() []grpc.ServerOption {
	ret := m.ctrl.Call(m, "GetGrpcServerOption")
	ret0, _ := ret[0].([]grpc.ServerOption)
	return ret0
}

// GetGrpcServerOption indicates an expected call of GetGrpcServerOption
func (mr *MockMetricsMockRecorder) GetGrpcServerOption() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGrpcServerOption", reflect.TypeOf((*MockMetrics)(nil).GetGrpcServerOption))
}

// GetHttpHandlers mocks base method
func (m *MockMetrics) GetHttpHandlers() map[string]http.Handler {
	ret := m.ctrl.Call(m, "GetHttpHandlers")
	ret0, _ := ret[0].(map[string]http.Handler)
	return ret0
}

// GetHttpHandlers indicates an expected call of GetHttpHandlers
func (mr *MockMetricsMockRecorder) GetHttpHandlers() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHttpHandlers", reflect.TypeOf((*MockMetrics)(nil).GetHttpHandlers))
}

// EmitCounter mocks base method
func (m *MockMetrics) EmitCounter(name string, value interface{}, tags ...metrics.T) error {
	varargs := []interface{}{name, value}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EmitCounter", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitCounter indicates an expected call of EmitCounter
func (mr *MockMetricsMockRecorder) EmitCounter(name, value interface{}, tags ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, value}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitCounter", reflect.TypeOf((*MockMetrics)(nil).EmitCounter), varargs...)
}

// EmitGauge mocks base method
func (m *MockMetrics) EmitGauge(name string, value interface{}, tags ...metrics.T) error {
	varargs := []interface{}{name, value}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EmitGauge", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitGauge indicates an expected call of EmitGauge
func (mr *MockMetricsMockRecorder) EmitGauge(name, value interface{}, tags ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, value}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitGauge", reflect.TypeOf((*MockMetrics)(nil).EmitGauge), varargs...)
}

// EmitHistogram mocks base method
func (m *MockMetrics) EmitHistogram(name string, value interface{}, tags ...metrics.T) error {
	varargs := []interface{}{name, value}
	for _, a := range tags {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EmitHistogram", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// EmitHistogram indicates an expected call of EmitHistogram
func (mr *MockMetricsMockRecorder) EmitHistogram(name, value interface{}, tags ...interface{}) *gomock.Call {
	varargs := append([]interface{}{name, value}, tags...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EmitHistogram", reflect.TypeOf((*MockMetrics)(nil).EmitHistogram), varargs...)
}
