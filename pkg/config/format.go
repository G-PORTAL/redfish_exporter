package config

type Action string

type Config struct {
	Verbose bool `yaml:"verbose"`

	Redfish struct {
		// Username Default username for collecting metrics
		Username string `yaml:"username"`

		// Password Default password for collecting metrics
		Password string `yaml:"password"`

		// Verify Whether to verify the TLS certificate
		VerifyTLS bool `yaml:"verifyTls"`
	} `yaml:"redfish"`

	PreActions []Action `yaml:"preActions"`
}
