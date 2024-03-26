package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

type ID string

type Metrics struct {
	metrics MetricCollection
}

func (b *Metrics) With(metrics ...Metric) {
	if b.metrics == nil {
		b.metrics = make(MetricCollection)
	}

	for _, metric := range metrics {
		b.metrics[metric.GetID()] = metric
	}
}

func (b *Metrics) Registry() *prometheus.Registry {
	registry := prometheus.NewRegistry()
	for _, collector := range b.metrics {
		registry.MustRegister(collector)
	}

	return registry
}
