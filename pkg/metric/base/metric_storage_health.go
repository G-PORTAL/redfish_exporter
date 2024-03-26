package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishStorageHealth ID = "redfish_storage_health"
)

var RedfishStorageStateLabels = []string{"system_id", "storage_id"}

func WithRedfishStorageHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) Metric {
	metric := &GaugeMetric{
		ID: RedfishStorageHealth,
		Gauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: string(RedfishStorageHealth),
			Help: "Indicates the health status of the storage. 1 = OK, 2 = Warning, 3 = Critical.",
		}, RedfishStorageStateLabels).With(labels),
	}

	metric.Set(float64(health))

	return metric
}
