package prometheus

import (
	"runtime"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"

	"github.com/rfyiamcool/go-metrics"
	"github.com/rfyiamcool/go-metrics/mock"
)

var (
	testGlobalTag            = metrics.Tag("test_global_tag", "test")
	testCounterTag           = metrics.Tag("test_counter_tag", "test")
	testGaugeTag             = metrics.Tag("test_gauge_tag", "test")
	testHistogramTag         = metrics.Tag("test_histogram_tag", "test")
	testCounterMetricsName   = "test.counter"
	testGaugeMetricsName     = "test.gauge"
	testHistogramMetricsName = "test.histogram"
)

var expectedMetrics = []string{
	"go_gc_duration_seconds_summary",
	"go_gc_duration_seconds_sum",
	"go_gc_duration_seconds_count",
	"go_goroutines",
	"go_info",
	"go_memstats_alloc_bytes",
	"go_memstats_alloc_bytes_total",
	"go_memstats_buck_hash_sys_bytes",
	"go_memstats_frees_total",
	"go_memstats_gc_cpu_fraction",
	"go_memstats_gc_sys_bytes",
	"go_memstats_heap_alloc_bytes",
	"go_memstats_heap_idle_bytes",
	"go_memstats_heap_inuse_bytes",
	"go_memstats_heap_objects",
	"go_memstats_heap_released_bytes",
	"go_memstats_heap_sys_bytes",
	"go_memstats_last_gc_time_seconds",
	"go_memstats_lookups_total",
	"go_memstats_mallocs_total",
	"go_memstats_mcache_inuse_bytes",
	"go_memstats_mcache_sys_bytes",
	"go_memstats_mspan_inuse_bytes",
	"go_memstats_mspan_sys_bytes",
	"go_memstats_next_gc_bytes",
	"go_memstats_other_sys_bytes",
	"go_memstats_stack_inuse_bytes",
	"go_memstats_stack_sys_bytes",
	"go_memstats_sys_bytes",
	"go_threads",
	"test_counter",
	"test_gauge",
	"test_histogram_bucket",
}

func TestEmitMetrics(t *testing.T) {

	// registry may be polluted by other module init before this one
	// so create new instances to make them correct
	newRegistry := prometheus.NewRegistry()
	registerer, gather = newRegistry, newRegistry
	defer func() {
		// reset after test
		registerer, gather = prometheus.DefaultRegisterer, prometheus.DefaultGatherer
	}()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ast := assert.New(t)

	t.Log(runtime.GOOS)

	p := NewMetrics(testGlobalTag)

	err := p.EmitCounter(testCounterMetricsName, 1, testCounterTag)
	ast.NoError(err)
	err = p.EmitCounter(testCounterMetricsName, 1, testCounterTag)
	ast.NoError(err)

	err = p.EmitGauge(testGaugeMetricsName, 1, testGaugeTag)
	ast.NoError(err)
	err = p.EmitGauge(testGaugeMetricsName, 1, testGaugeTag)
	ast.NoError(err)

	err = p.EmitHistogram(testHistogramMetricsName, 1, testHistogramTag)
	ast.NoError(err)
	err = p.EmitHistogram(testHistogramMetricsName, 1, testHistogramTag)
	ast.NoError(err)

	mockMetrics := mock.NewMockMetrics(ctrl)

	for _, m := range expectedMetrics {
		mockMetrics.EXPECT().EmitGauge(m, gomock.Any(), gomock.Any()).AnyTimes()
	}

	err = emitMetrics(mockMetrics)
	ast.NoError(err)
}
