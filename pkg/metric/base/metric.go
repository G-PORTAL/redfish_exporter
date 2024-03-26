package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricCollection map[ID]Metric

type Metric interface {
	GetID() ID
	prometheus.Collector
	prometheus.Metric
}
