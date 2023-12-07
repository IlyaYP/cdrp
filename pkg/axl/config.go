package axl

import "fmt"

// Config keeps AXL configuration.
type Config struct {
	CUCM   string `mapstructure:"cucm"`
	User   string `mapstructure:"user"`
	Pass   string `mapstructure:"pass"`
	Insec  bool   `mapstructure:"insec"`
	Schema string `mapstructure:"schema"`
	Dump   bool   `mapstructure:"dump"`
}

// Validate performs a basic validation.
func (c Config) Validate() error {
	if c.CUCM == "" {
		return fmt.Errorf("%s field: empty", "CUCM")
	}

	return nil
}

// NewDefaultConfig builds a Config with default values.
func NewDefaultConfig() Config {
	return Config{}
}
