package proxy

import (
	"gopkg.in/yaml.v3"
)

// Config -
type Config struct {
	Version string  `yaml:"version,omitempty"`
	Routes  []Route `yaml:"routes"`
}

// NewConfig -
func NewConfig(content []byte) (*Config, error) {
	c := &Config{}

	if err := yaml.Unmarshal(content, c); err != nil {
		return nil, err
	}

	return c, nil
}
