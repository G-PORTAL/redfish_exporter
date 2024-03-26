package redfish

import (
	"fmt"

	"github.com/g-portal/redfish_exporter/pkg/drivers/redfish/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish"
)

type Redfish struct {
	client *gofish.APIClient
}

const TLSHandshakeTimeout = 30

func (rf *Redfish) Connect(host, username, password string, verifyTLS bool) error {
	var err error

	cfg := gofish.ClientConfig{
		Endpoint:            host,
		Username:            username,
		Password:            password,
		Insecure:            !verifyTLS,
		TLSHandshakeTimeout: TLSHandshakeTimeout,
	}
	// Debug

	rf.client, err = gofish.Connect(cfg)
	if err != nil {
		return fmt.Errorf("error connecting to redfish: %w", err)
	}

	return nil
}

func (rf *Redfish) GetMetrics() (*prometheus.Registry, error) {
	allMetrics := metrics.NewMetrics(rf.client)
	if err := allMetrics.Collect(); err != nil {
		return nil, fmt.Errorf("error collecting metrics: %w", err)
	}

	return allMetrics.Registry(), nil
}

func (rf *Redfish) Disconnect() error {
	rf.client.Logout()

	return nil
}
