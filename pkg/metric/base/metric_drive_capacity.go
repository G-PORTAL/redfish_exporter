package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishDriveCapacity ID = "redfish_drive_capacity"
)

var RedfishDriveCapacityLabels = []string{"system_id", "drive_id", "storage_id"}

func (b *Metrics) WithRedfishDriveCapacityMetric(size float64, labels prometheus.Labels) {
	b.SetGauge(RedfishDriveCapacity, labels, size)
}
