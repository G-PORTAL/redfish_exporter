package base

import (
	"log"

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
	for id, collector := range b.metrics {
		log.Printf("Registering metric %s\n", id)

		// debug log?
		registry.MustRegister(collector)
	}

	return registry
}
