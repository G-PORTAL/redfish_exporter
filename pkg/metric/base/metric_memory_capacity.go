package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishMemoryCapacity ID = "redfish_memory_capacity"
)

var RedfishMemoryCapacityLabels = []string{"system_id", "memory_id"}

func (b *Metrics) WithRedfishMemoryCapacityMetric(size int, labels prometheus.Labels) {
	b.SetGauge(RedfishMemoryCapacity, labels, float64(size))
}
