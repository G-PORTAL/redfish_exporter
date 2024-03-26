package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishStorageHealth ID = "redfish_storage_health"
)

var RedfishStorageHealthLabels = []string{"system_id", "storage_id"}

func (b *Metrics) WithRedfishStorageHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) {
	b.SetGauge(RedfishStorageHealth, labels, float64(health))
}
