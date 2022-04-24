package api

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics is a wrapper interface for prometheus
type Metrics interface {
	// Counter creates or returns a prometheus counter by name
	// if the key is registered by other interface, it will return nil
	// Counter: invoke CounterWithLabelsWithError(name, nil) and ignore error
	// CounterWithLabels: invoke CounterWithLabelsWithError(name, dynamicLables) and ignore error
	Counter(name string) prometheus.Counter
	CounterWithLabelsWithError(name string, dynamicLabels map[string]string) (prometheus.Counter, error)
	CounterWithLabels(name string, dynamicLabels map[string]string) prometheus.Counter

	// Gauge creates or returns a prometheus gauge by key
	// if the key is registered by other interface, it will return nil
	// Gauge: invoke GaugeWithLabelsWithError(name, nil) and ignore error
	// GaugeWithLabels: invoke GaugeWithLabelsWithError(name, dynamicLables) and ignore error
	Gauge(key string) prometheus.Gauge
	GaugeWithLabels(name string, dynamicLabels map[string]string) prometheus.Gauge
	GaugeWithLabelsWithError(name string, dynamicLabels map[string]string) (prometheus.Gauge, error)

	// Histogram creates or returns a prometheus histogram by key
	// if the key is registered by other interface, it will be return nil
	// Histogram: invoke HistogramWithLabelsWithError(name, buckets, nil) and ignore error
	// HistogramWithLabels: invoke HistogramWithLabelsWithError(name, buckets, dynamicLables) and ignore error
	Histogram(key string, buckets []float64) prometheus.Histogram
	HistogramWithLabels(name string, buckets []float64, dynamicLabels map[string]string) prometheus.Histogram
	HistogramWithLabelsWithError(name string, buckets []float64, dynamicLabels map[string]string) (prometheus.Histogram, error)

	// Summary creates or returns a summary histogram by key, objectives default 50%, 90% 95% and 99%
	// if the key is registered by other interface, it will be return nil
	// Summary: invoke SummaryWithLabelsWithError(name, objectives, nil) and ignore error
	// SummaryWithLables: invoke SummaryWithLabelsWithError(name, objectives, dynamicLables) and ignore error
	Summary(key string, objectives map[float64]float64) prometheus.Summary
	SummaryWithLables(name string, objectives map[float64]float64, dynamicLabels map[string]string) prometheus.Summary
	SummaryWithLabelsWithError(name string, objectives map[float64]float64, dynamicLabels map[string]string) (prometheus.Summary, error)

	// UnregisterAll unregister all metrics.  (Mostly for testing.)
	UnregisterAll()

	// Delete With metricsVec, not unregister metrics
	DeleteWithLabels(name string, labels map[string]string) bool
}
