package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishDriveHealth ID = "redfish_drive_health"
)

var RedfishDriveHealthLabels = []string{"system_id", "drive_id", "storage_id"}

func (b *Metrics) WithRedfishDriveHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) {
	b.SetGauge(RedfishDriveHealth, labels, float64(health))
}
