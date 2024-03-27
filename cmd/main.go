package main

import (
	"flag"
	"log"

	"github.com/g-portal/redfish_exporter/pkg/config"
	"github.com/g-portal/redfish_exporter/pkg/metric"
	"github.com/g-portal/redfish_exporter/pkg/middleware"
	"github.com/gin-gonic/gin"
)

var (
	listenAddr string
	configPath string
)

func init() {
	flag.StringVar(&configPath,
		"config.file",
		"/etc/redfish_exporter/config.yml",
		"Defines the path to the platform management config",
	)
	flag.StringVar(&listenAddr,
		"web.listen-address",
		"0.0.0.0:9096",
		"Address the exporter listens on",
	)
	flag.Parse()

	config.SetPath(configPath)
}

func main() {
	if conf := config.GetConfig(); !conf.Verbose {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()
	server.Use(gin.Recovery())
	server.Use(middleware.ErrorLogger)
	server.Use(gin.Logger())

	server.GET("/metrics", metric.Handle)
	log.Println("Starting listening on: " + listenAddr)
	if err := server.Run(listenAddr); err != nil {
		log.Printf("Error starting http server: %v", err)
	}
}
