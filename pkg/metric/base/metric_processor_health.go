package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishProcessorHealth ID = "redfish_processor_health"
)

var RedfishProcessorStateLabels = []string{"system_id", "processor_id"}

func WithRedfishProcessorHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) Metric {
	metric := &GaugeMetric{
		ID: RedfishProcessorHealth,
		Gauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: string(RedfishProcessorHealth),
			Help: "Indicates the health status of the processor. 1 = OK, 2 = Warning, 3 = Critical.",
		}, RedfishProcessorStateLabels).With(labels),
	}

	metric.Set(float64(health))

	return metric
}
