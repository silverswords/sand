package logger

import (
	"encoding/json"
	"io/ioutil"
)

type Builder struct {
	*Config
}

func init() {

}

func CreateBuilder(path string) *Builder {
	c := &Config{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, c); err != nil {
		panic(err)
	}

	return &Builder{
		Config: c,
	}
}

func (b *Builder) Build(opts ...Option) *Builder {
	for _, opt := range opts {
		opt(b)
	}

	return b
}
