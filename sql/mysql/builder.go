package mysql

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	json "github.com/json-iterator/go"
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

func (b *Builder) Run() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		b.UserName, b.Password, b.Host, b.Port, b.Database, b.Charset, true, "Local")

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	return db
}
