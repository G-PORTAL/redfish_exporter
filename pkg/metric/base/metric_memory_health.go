package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishMemoryHealth ID = "redfish_memory_health"
)

var RedfishMemoryHealthLabels = []string{"system_id", "memory_id"}

func (b *Metrics) WithRedfishMemoryHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) {
	b.SetGauge(RedfishMemoryHealth, labels, float64(health))
}
