package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishDriveHealth ID = "redfish_drive_health"
)

var RedfishDriveStateLabels = []string{"system_id", "drive_id", "storage_id"}

func WithRedfishDriveHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) Metric {
	metric := &GaugeMetric{
		ID: RedfishDriveHealth,
		Gauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: string(RedfishDriveHealth),
			Help: "Indicates the health status of the drive. 1 = OK, 2 = Warning, 3 = Critical.",
		}, RedfishDriveStateLabels).With(labels),
	}

	metric.Set(float64(health))

	return metric
}
