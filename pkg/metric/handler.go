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

	client, err := api.NewClient(params.Host, params.Username, params.Password, false)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

		return
	}

	log.Println("finished with handling")
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
	Username string
	Password string
	Host     string
}

func extractCollectorParams(req *http.Request) collectorParams {
	cfg := config.GetConfig()
	params := collectorParams{
		Host:     "",
		Username: cfg.Redfish.Username,
		Password: cfg.Redfish.Password,
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

	return params
}
