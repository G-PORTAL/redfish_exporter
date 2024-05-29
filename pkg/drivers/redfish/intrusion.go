package redfish

import (
	"fmt"

	"github.com/stmcginnis/gofish/redfish"
)

func (rf *Redfish) ResetChassisIntrusion() error {
	chassisCollection, err := rf.client.Service.Chassis()
	if err != nil {
		return fmt.Errorf("error getting chassis: %w", err)
	}

	for _, chassis := range chassisCollection {
		if chassis.PhysicalSecurity.IntrusionSensor == redfish.HardwareIntrusionIntrusionSensor {
			if err = chassis.Patch(chassis.ODataID, map[string]interface{}{
				"PhysicalSecurity": map[string]interface{}{
					"IntrusionSensor": redfish.NormalIntrusionSensor,
				},
			}); err != nil {
				return fmt.Errorf("error updating chassis: %w", err)
			}
		}
	}

	return nil
}
