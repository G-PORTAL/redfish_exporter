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

func (b *Metrics) WithRedfishPowerStateMetric(state RedfishPowerStateType, labels prometheus.Labels) {
	b.SetGauge(RedfishPowerState, labels, float64(state))
}
