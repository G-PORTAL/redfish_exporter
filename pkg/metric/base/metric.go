package base

import (
	"github.com/prometheus/client_golang/prometheus"
)

type MetricCollection []Metric

type Metric interface {
	GetID() ID
	prometheus.Collector
	prometheus.Metric
}
