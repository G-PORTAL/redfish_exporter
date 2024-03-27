package metric

import (
	"log"
	"net/http"

	"github.com/g-portal/redfish_exporter/pkg/api"
	"github.com/g-portal/redfish_exporter/pkg/config"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handle(ctx *gin.Context) {
	params := extractCollectorParams(ctx.Request)

	client, err := api.NewClient(params.Host, params.Username, params.Password, params.VerifyTLS)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	defer func() {
		err := client.Disconnect()
		if err != nil {
			log.Printf("error disconnecting: %v", err)
		}
	}()

	// Get metrics from the client.
	registry, err := client.GetMetrics()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	// Delegate http serving to Prometheus client library, which will call collector.Collect.
	promhttp.HandlerFor(registry, promhttp.HandlerOpts{}).ServeHTTP(ctx.Writer, ctx.Request)
}

type collectorParams struct {
	Username  string
	Password  string
	Host      string
	VerifyTLS bool
}

func extractCollectorParams(req *http.Request) collectorParams {
	cfg := config.GetConfig()
	params := collectorParams{
		Host:      "",
		Username:  cfg.Redfish.Username,
		Password:  cfg.Redfish.Password,
		VerifyTLS: cfg.Redfish.VerifyTLS,
	}
	if host := req.URL.Query().Get("host"); host != "" {
		params.Host = host
	}
	if username := req.URL.Query().Get("username"); username != "" {
		params.Username = username
	}
	if password := req.URL.Query().Get("password"); password != "" {
		params.Password = password
	}

	if verifyTLS := req.URL.Query().Get("verify_tls"); verifyTLS == "true" {
		params.VerifyTLS = true
	} else if verifyTLS == "false" {
		params.VerifyTLS = false
	}

	return params
}
