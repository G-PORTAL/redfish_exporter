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

func (b *Metrics) WithRedfishHealthMetric(health RedfishHealthStatus, labels prometheus.Labels) {
	b.SetGauge(RedfishHealth, labels, float64(health))
}
