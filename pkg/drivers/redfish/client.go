package redfish

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/g-portal/redfish_exporter/pkg/config"
	"github.com/g-portal/redfish_exporter/pkg/drivers/redfish/metrics"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stmcginnis/gofish"
)

type Redfish struct {
	client *gofish.APIClient
}

const Timeout = 30 * time.Second

var ErrDefaultTransport = errors.New("could not convert http.DefaultTransport to *http.Transport")

func (rf *Redfish) Connect(host, username, password string, verifyTLS bool) error {
	var err error

	transport, ok := http.DefaultTransport.(*http.Transport)
	if !ok {
		return ErrDefaultTransport
	}

	transport.TLSHandshakeTimeout = Timeout
	transport.TLSClientConfig = &tls.Config{
		MinVersion:         tls.VersionTLS11,
		InsecureSkipVerify: !verifyTLS,
		CipherSuites: []uint16{
			// See https://go.dev/src/crypto/tls/cipher_suites.go for future updates
			// TLS 1.0 - 1.2 cipher suites
			tls.TLS_RSA_WITH_RC4_128_SHA,
			tls.TLS_RSA_WITH_3DES_EDE_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_RC4_128_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_RC4_128_SHA,
			tls.TLS_ECDHE_RSA_WITH_3DES_EDE_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256,
			// TLS 1.3 cipher suites.
			tls.TLS_AES_128_GCM_SHA256,
			tls.TLS_AES_256_GCM_SHA384,
			tls.TLS_CHACHA20_POLY1305_SHA256,
		},
	}

	httpClient := http.DefaultClient
	httpClient.Transport = transport
	httpClient.Timeout = Timeout

	cfg := gofish.ClientConfig{
		Endpoint:            host,
		Username:            username,
		Password:            password,
		Insecure:            !verifyTLS,
		TLSHandshakeTimeout: int(Timeout.Seconds()),
		HTTPClient:          httpClient,
	}

	if verbose := config.GetConfig().Verbose; verbose {
		cfg.DumpWriter = os.Stdout
	}

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
