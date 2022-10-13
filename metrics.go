package metrics

import (
	"net/http"

	"google.golang.org/grpc"
)

func Tag(name, value string) T {
	return T{
		Name:  name,
		Value: value,
	}
}

type T struct {
	Name  string
	Value string
}

// Metrics abstracts metrics common define
type Metrics interface {

	// GetGrpcServerOption returns the grpc.ServerOptions for metrics interceptors
	GetGrpcServerOption() []grpc.ServerOption

	// GetHttpHandlers returns the map from path to handler register some handle for metrics
	GetHttpHandlers() map[string]http.Handler

	// EmitCounter emits the count
	EmitCounter(name string, value interface{}, tags ...T) error

	// EmitGauge emits the state
	EmitGauge(name string, value interface{}, tags ...T) error

	// EmitHistogram emits the histogram
	EmitHistogram(name string, value interface{}, tags ...T) error // histogram
}
