package mock

import (
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/rfyiamcool/go-metrics"
)

func TestNewMinimalMetrics(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMinimalMetrics(ctrl)

	t.Run("no tag", func(t *testing.T) {
		var tags []metrics.T
		m.EmitGauge("a", "b", tags...)
		m.EmitHistogram("a", "b", tags...)
		m.EmitCounter("a", "b", tags...)
	})

	t.Run("1 tag", func(t *testing.T) {
		tags := []metrics.T{metrics.Tag("a", "b")}
		m.EmitGauge("a", "b", tags...)
		m.EmitHistogram("a", "b", tags...)
		m.EmitCounter("a", "b", tags...)
	})

	t.Run("1+ tag", func(t *testing.T) {
		tags := []metrics.T{metrics.Tag("a", "b"), metrics.Tag("a", "b")}
		m.EmitGauge("a", "b", tags...)
		m.EmitHistogram("a", "b", tags...)
		m.EmitCounter("a", "b", tags...)
	})

}
