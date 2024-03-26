package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishProcessorHealth ID = "redfish_processor_health"
)

var RedfishProcessorStateLabels = []string{"system_id", "processor_id"}

func (b *Metrics) WithRedfishProcessorHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) {
	b.SetGauge(RedfishProcessorHealth, labels, float64(health))
}
