package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishMemoryHealth ID = "redfish_memory_health"
)

var RedfishMemoryStateLabels = []string{"system_id", "memory_id"}

func WithRedfishMemoryHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) Metric {
	metric := &GaugeMetric{
		ID: RedfishMemoryHealth,
		Gauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: string(RedfishMemoryHealth),
			Help: "Indicates the health status of the memory. 1 = OK, 2 = Warning, 3 = Critical.",
		}, RedfishMemoryStateLabels).With(labels),
	}

	metric.Set(float64(health))

	return metric
}
