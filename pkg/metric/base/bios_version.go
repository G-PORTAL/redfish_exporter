package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	BiosVersion ID = "redfish_bios_version"
)

var RedfishBiosVersionLabels = []string{"system_id", "version"}

func WithRedfishBiosVersionMetric(labels prometheus.Labels) Metric {
	metric := &GaugeMetric{
		ID: BiosVersion,
		Gauge: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: string(BiosVersion),
			Help: "Indicates the BIOS version of the system.",
		}, RedfishBiosVersionLabels).With(labels),
	}

	metric.Set(1)

	return metric
}
