package metrics

import (
	"fmt"
	"log"

	"github.com/g-portal/redfish_exporter/pkg/metric/base"
	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type Metrics struct {
	base.Metrics

	api *gofish.APIClient
}

func NewMetrics(gofish *gofish.APIClient) *Metrics {
	return &Metrics{
		api: gofish,
	}
}

func (m *Metrics) Collect() error {
	service := m.api.GetService()

	systems, err := service.Systems()
	if err != nil {
		return fmt.Errorf("error getting systems: %w", err)
	}

	for _, system := range systems {
		m.With(base.WithRedfishHealthMetric(convertHealthStatus(system.Status.Health), map[string]string{
			"system_id": system.ID,
		}), base.WithRedfishPowerStateMetric(convertPowerState(system.PowerState), map[string]string{
			"system_id": system.ID,
		}))

		if memory, err := system.Memory(); err == nil {
			for _, mem := range memory {
				m.With(base.WithRedfishMemoryHealthMetric(convertHealthStatus(mem.Status.Health), map[string]string{
					"system_id": system.ID,
					"memory_id": mem.ID,
				}))
			}
		} else {
			log.Printf("error getting memory: %s", err)
		}

		if storage, err := system.Storage(); err == nil {
			for _, store := range storage {
				m.With(base.WithRedfishStorageHealthMetric(convertHealthStatus(store.Status.Health), map[string]string{
					"system_id":  system.ID,
					"storage_id": store.ID,
				}))

				if drives, err := store.Drives(); err == nil {
					for _, drive := range drives {
						m.With(base.WithRedfishDriveHealthMetric(convertHealthStatus(drive.Status.Health), map[string]string{
							"system_id":  system.ID,
							"storage_id": store.ID,
							"drive_id":   drive.ID,
						}))
					}
				} else {
					log.Printf("error getting drives: %s", err)
				}
			}
		} else {
			log.Printf("error getting storage: %s", err)
		}

		if cpus, err := system.Processors(); err == nil {
			for _, cpu := range cpus {
				m.With(base.WithRedfishProcessorHealthMetric(convertHealthStatus(cpu.Status.Health), map[string]string{
					"system_id": system.ID,
					"cpu_id":    cpu.ID,
				}))
			}
		} else {
			log.Printf("error getting processors: %s", err)
		}

		if system.BIOSVersion != "" {
			m.With(base.WithRedfishBiosVersionMetric(map[string]string{
				"system_id": system.ID,
				"version":   system.BIOSVersion,
			}))
		}
	}

	return nil
}

func convertHealthStatus(status common.Health) base.RedfishHealthStatus {
	switch status {
	case common.OKHealth:
		return base.RedfishHealthOK
	case common.WarningHealth:
		return base.RedfishHealthWarning
	case common.CriticalHealth:
		return base.RedfishHealthCritical
	default:
		return base.RedfishHealthWarning
	}
}

func convertPowerState(state redfish.PowerState) base.RedfishPowerStateType {
	switch state {
	case redfish.OnPowerState, redfish.PoweringOnPowerState:
		return base.RedfishPowerStateON
	case redfish.OffPowerState, redfish.PoweringOffPowerState:
		return base.RedfishPowerStateOFF
	default:
		return base.RedfishPowerStateON
	}
}
