package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	RedfishBiosVersion ID = "redfish_bios_version"
)

var RedfishBiosVersionLabels = []string{"system_id", "version"}

func (b *Metrics) WithRedfishBiosVersionMetric(labels prometheus.Labels) {
	b.SetGauge(RedfishBiosVersion, labels, 1.0)
}
