package mock

import (
	"github.com/golang/mock/gomock"

	"github.com/rfyiamcool/go-metrics"
)

// NewMinimalMetrics returns a mock metrics implement as a stub for test
func NewMinimalMetrics(ctrl *gomock.Controller) metrics.Metrics {
	m := NewMockMetrics(ctrl)
	m.EXPECT().EmitGauge(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	m.EXPECT().EmitCounter(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	m.EXPECT().EmitHistogram(gomock.Any(), gomock.Any(), gomock.Any()).AnyTimes()
	m.EXPECT().GetHttpHandlers().AnyTimes()
	m.EXPECT().GetGrpcServerOption().AnyTimes()
	return m
}
