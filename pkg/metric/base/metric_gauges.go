package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

type GaugeMetric struct {
	ID ID

	prometheus.Gauge
}

func (g *GaugeMetric) GetID() ID {
	return g.ID
}
