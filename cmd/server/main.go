package main

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/pkg/proxy"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(2)
	go func() {
		r := gin.Default()
		queryG := r.Group("/query")

		queryG.GET("/ping", func(c *gin.Context) {
			a := c.Query("a")

			c.JSON(200, gin.H{
				"message": a,
			})
		})

		r.Run(":10000")

		wg.Done()
	}()

	go func() {
		content, err := ioutil.ReadFile("./cmd/config/config.yaml")
		if err != nil {
			log.Fatal(err)
		}
		c, err := proxy.NewConfig(content)
		if err != nil {
			log.Fatal(err)
		}

		r := gin.Default()

		for i := 0; i < len(c.Routes); i++ {
			route := &c.Routes[i]
			log.Printf("route - [%s]\n", route.Name)

			p := proxy.BuildProxy(route)

			gp := r.Group(route.Name)
			gp.Use(func(c *gin.Context) {
				log.Printf("[%s]", c.Request.URL.Path)
				p.ServeHTTP(c.Writer, c.Request)
			})

			gp.GET("/blackhole", func(c *gin.Context) {
			})
		}

		r.Run(":10010")

		wg.Done()
	}()

	wg.Wait()
}
