package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var (
	configPath     string
	exporterConfig *Config
)

func SetPath(opt string) {
	configPath = opt
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Error loading config file: %v", err)
	}
	exporterConfig = &Config{}
	exporterConfig.reload()
}

// reload Utility function for reloading the config file.
func (c *Config) reload() {
	configContent, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	err = yaml.Unmarshal(configContent, c)
	if err != nil {
		log.Fatalf("Unable to parse config file: %v", err)
	}
	log.Println("Config file reloaded")
}

// GetConfig Returns the current config instance.
func GetConfig() *Config {
	return exporterConfig
}
