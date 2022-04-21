package api

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Metrics is a wrapper interface for prometheus
type Metrics interface {
	// Counter creates or returns a prometheus counter by key
	// if the key is registered by other interface, it will be panic
	Counter(key string) prometheus.Counter

	// Gauge creates or returns a prometheus gauge by key
	// if the key is registered by other interface, it will be panic
	Gauge(key string) prometheus.Gauge

	// Histogram creates or returns a prometheus histogram by key
	// if the key is registered by other interface, it will be panic
	Histogram(key string, buckets []float64) prometheus.Histogram

	// Summary creates or returns a summary histogram by key, objectives default 50%, 90% 95% and 99%
	// if the key is registered by other interface, it will be panic
	Summary(key string, objectives map[float64]float64) prometheus.Summary

	// UnregisterAll unregister all metrics.  (Mostly for testing.)
	UnregisterAll()
}

// MetricsWithLabels is a wrapper interface for prometheus.Metrics.
// Usually create only one metcis object, because the name will registry once.
type MetricsWithLabels interface {
	// Counter creates or returns a prometheus counter by labels
	Counter(labels map[string]string) (prometheus.Counter, error)

	// Gauge creates or returns a prometheus gauge by labels
	Gauge(labels map[string]string) (prometheus.Gauge, error)

	// Histogram creates or returns a prometheus histogram by labels
	Histogram(labels map[string]string, buckets []float64) (prometheus.Histogram, error)

	// Summary creates or returns a summary histogram by labels, objectives default 50%, 90% 95% and 99%
	Summary(labels map[string]string, objectives map[float64]float64) (prometheus.Summary, error)

	// Delete prometheus.Metrics With Labels
	Delete(labels map[string]string) bool

	// Free nnregister from Prometheus
	Free()
}
