package metrics

import (
	"log"

	"github.com/g-portal/redfish_exporter/pkg/oem/hpe"
	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// vendorSpecifics is a helper function to collect vendor specific metrics.
func (m *Metrics) vendorSpecifics(system *redfish.ComputerSystem) {
	switch system.Manufacturer {
	case hpe.Manufacturer:
		//nolint: nestif
		if hpeStorage, err := hpe.GetSmartStorage(system); err == nil {
			if arrayControllers, err := hpeStorage.ArrayControllers(); err == nil {
				for _, controller := range arrayControllers {
					m.WithRedfishStorageHealthMetric(convertHealthStatus(hpeStorage.Status.Health,
						hpeStorage.Status.State == common.EnabledState), map[string]string{
						"system_id":  system.ID,
						"storage_id": controller.ID,
					})

					if drives, err := controller.PhysicalDrives(); err == nil {
						for _, drive := range drives {
							m.WithRedfishDriveHealthMetric(convertHealthStatus(drive.Status.Health,
								drive.Status.State == common.EnabledState), map[string]string{
								"system_id":  system.ID,
								"storage_id": controller.ID,
								"drive_id":   drive.ID,
							})
							m.WithRedfishDriveCapacityMetric(float64(drive.CapacityMiB), map[string]string{
								"system_id":  system.ID,
								"storage_id": controller.ID,
								"drive_id":   drive.ID,
							})
						}
					} else {
						log.Printf("error getting drives: %s", err)
					}
				}
			} else {
				log.Printf("error getting array controllers: %s", err)
			}
		} else {
			log.Printf("error getting SmartStorage: %s", err)
		}
	default:
		// no vendor specific metrics to collect
	}
}
