package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/silverswords/sand/pkg/proxy"
)

func main() {
	content, err := ioutil.ReadFile("./cmd/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	c, err := proxy.NewConfig(content)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%#v", c)
}
