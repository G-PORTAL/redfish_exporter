package base

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type ID string

type Metrics struct {
	collectors map[ID]prometheus.Collector
}

func NewMetrics() Metrics {
	return Metrics{
		collectors: map[ID]prometheus.Collector{
			RedfishBiosVersion: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishBiosVersion),
				Help: "Indicates the BIOS version of the system.",
			}, RedfishBiosVersionLabels),
			RedfishDriveCapacity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishDriveCapacity),
				Help: "Indicates the capacity of the drive in MiB.",
			}, RedfishDriveCapacityLabels),
			RedfishDriveHealth: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishDriveHealth),
				Help: "Indicates the health status of the drive. 1 = OK, 2 = Warning, 3 = Critical.",
			}, RedfishDriveHealthLabels),
			RedfishHealth: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishHealth),
				Help: "Indicates the health status of the system. 1 = OK, 2 = Warning, 3 = Critical.",
			}, RedfishHealthLabels),
			RedfishMemoryCapacity: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishMemoryCapacity),
				Help: "Indicates the capacity of the memory in MiB.",
			}, RedfishMemoryCapacityLabels),
			RedfishMemoryHealth: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishMemoryHealth),
				Help: "Indicates the health status of the memory. 1 = OK, 2 = Warning, 3 = Critical.",
			}, RedfishMemoryHealthLabels),
			RedfishPowerState: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishPowerState),
				Help: "Indicates the power state of the system. 1 = ON, 2 = OFF.",
			}, RedfishPowerStateLabels),
			RedfishProcessorHealth: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishProcessorHealth),
				Help: "Indicates the health status of the processor. 1 = OK, 2 = Warning, 3 = Critical.",
			}, RedfishProcessorStateLabels),
			RedfishStorageHealth: prometheus.NewGaugeVec(prometheus.GaugeOpts{
				Name: string(RedfishStorageHealth),
				Help: "Indicates the health status of the storage. 1 = OK, 2 = Warning, 3 = Critical.",
			}, RedfishStorageHealthLabels),
		},
	}
}

func (b *Metrics) SetGauge(id ID, labels prometheus.Labels, value float64) {
	if collector, ok := b.collectors[id]; ok {
		if gauge, ok := collector.(*prometheus.GaugeVec); ok {
			metric := gauge.With(labels)
			metric.Set(value)
		}
	}
}

func (b *Metrics) Registry() *prometheus.Registry {
	registry := prometheus.NewRegistry()
	for id, collector := range b.collectors {
		log.Printf("Registering collector %q.", id)
		registry.MustRegister(collector)
	}

	return registry
}
