package api

import (
	"fmt"

	"github.com/g-portal/redfish_exporter/pkg/drivers/redfish"
	"github.com/prometheus/client_golang/prometheus"
)

type Client interface {
	Connect(host, username, password string, tlsVerify bool) error
	GetMetrics() (*prometheus.Registry, error)
	Disconnect() error
}

func NewClient(host, username, password string, tlsVerify bool) (Client, error) {
	client := &redfish.Redfish{}
	err := client.Connect(host, username, password, tlsVerify)
	if err != nil {
		return nil, fmt.Errorf("error connecting to client: %w", err)
	}

	return client, nil
}
