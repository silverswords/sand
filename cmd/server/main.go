package main

import (
	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/pkg/proxy"
)

func main() {
	p := proxy.BuildProxy()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		a := c.Query("a")

		c.JSON(200, gin.H{
			"message": a,
		})
	})

	r.GET("/proxy", func(c *gin.Context) {
		p.ServeHTTP(c.Writer, c.Request)
	})

	r.Run(":10000")
}
