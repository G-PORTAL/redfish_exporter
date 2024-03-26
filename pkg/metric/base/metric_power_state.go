package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishPowerState ID = "redfish_power_state"
)

var RedfishPowerStateLabels = []string{"system_id"}

type RedfishPowerStateType float64

const (
	RedfishPowerStateON  RedfishPowerStateType = 1
	RedfishPowerStateOFF RedfishPowerStateType = 2
)

func WithRedfishPowerStateMetric(state RedfishPowerStateType, labels prometheus.Labels) Metric {
	metric := &GaugeMetric{
		ID: RedfishPowerState,
		Gauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: string(RedfishPowerState),
			Help: "Indicates the power state of the system. 1 = ON, 2 = OFF.",
		}, RedfishPowerStateLabels).With(labels),
	}

	metric.Set(float64(state))

	return metric
}
