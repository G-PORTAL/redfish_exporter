package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishHealth ID = "redfish_health"
)

var RedfishHealthLabels = []string{"system_id"}

type RedfishHealthStatus float64

const (
	RedfishHealthOK       RedfishHealthStatus = 1
	RedfishHealthWarning  RedfishHealthStatus = 2
	RedfishHealthCritical RedfishHealthStatus = 3
)

func WithRedfishHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) Metric {
	metric := &GaugeMetric{
		ID: RedfishHealth,
		Gauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: string(RedfishHealth),
			Help: "Indicates the health status of the system. 1 = OK, 2 = Warning, 3 = Critical.",
		}, RedfishHealthLabels).With(labels),
	}

	metric.Set(float64(health))

	return metric
}
