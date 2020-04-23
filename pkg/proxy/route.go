package proxy

import (
	"github.com/creasty/defaults"
)

// Route -
type Route struct {
	Name   string `yaml:"name,omitempty"`
	Host   string `yaml:"host,omitempty"`
	Scheme string `yaml:"scheme" default:"http"`
}

// UnmarshalYAML - yaml interface
func (r *Route) UnmarshalYAML(unmarshal func(interface{}) error) error {
	if err := defaults.Set(r); err != nil {
		return err
	}

	type plain Route
	if err := unmarshal((*plain)(r)); err != nil {
		return err
	}

	return nil
}
