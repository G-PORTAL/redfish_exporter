package config

type Config struct {
	Redfish struct {
		// Username Default username for collecting metrics
		Username string `yaml:"username"`

		// Password Default password for collecting metrics
		Password string `yaml:"password"`
	} `yaml:"redfish"`
}
