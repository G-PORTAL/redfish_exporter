package actions

import (
	"fmt"

	"github.com/g-portal/redfish_exporter/pkg/api"
	"github.com/g-portal/redfish_exporter/pkg/config"
)

const (
	ResetIntrusionAlert config.Action = "reset_intrusion_alert"
)

func Execute(action config.Action, client api.Client) error {
	switch action {
	case ResetIntrusionAlert:
		err := client.ResetChassisIntrusion()
		if err != nil {
			return fmt.Errorf("error resetting chassis intrusion: %w", err)
		}

		return nil
	default:
		//nolint: err113
		return fmt.Errorf("unknown action %s", action)
	}
}
