package metrics

import (
	"fmt"

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
			"system_id": system.ODataID,
		}), base.WithRedfishPowerStateMetric(convertPowerState(system.PowerState), map[string]string{
			"system_id": system.ODataID,
		}))
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
