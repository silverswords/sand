package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/gorilla/websocket"
	"github.com/silverswords/sand/pkg/proxy"
)

func main() {
	flag.Parse()

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		r := gin.Default()
		queryG := r.Group("/query")

		queryG.GET("/ping", func(c *gin.Context) {
			a := c.Query("a")

			c.JSON(200, gin.H{
				"message": a,
			})
		})

		var upgrader = websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		queryG.GET("/ws", func(c *gin.Context) {
			conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

			if err != nil {
				log.Printf("[WebSocket]Failed to upgrade ws: %+v", err)
				return
			}

			for {
				t, msg, err := conn.ReadMessage()
				if err != nil {
					log.Printf("[WebSocket] %+v", err)
					break
				}
				conn.WriteMessage(t, msg)
			}
		})

		r.Run(":10000")

		wg.Done()
	}()

	go func() {
		r := gin.Default()
		queryG := r.Group("/mail")

		queryG.GET("/ping", func(c *gin.Context) {
			a := c.Query("a")

			c.JSON(200, gin.H{
				"message": "mail:" + a,
			})
		})

		r.Run(":10001")

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

		handlerMap := map[string]http.Handler{}

		for i := 0; i < len(c.Routes); i++ {
			route := &c.Routes[i]
			glog.V(2).Info("[Loading] route - ", route.Name)

			handlerMap[route.Name] = proxy.BuildProxy(route)
		}

		r := gin.Default()

		r.NoRoute(func(c *gin.Context) {
			path := c.Request.URL.Path
			glog.V(2).Info("[Proxy Entry] Path ->", path)

			strParts := strings.Split(path, "/")
			strParts = strParts[0 : len(strParts)-1]

			group := strings.Join(strParts, "/")

			if handler, ok := handlerMap[group]; ok {
				handler.ServeHTTP(c.Writer, c.Request)
			}
		})

		r.Run(":10010")

		wg.Done()
	}()

	wg.Wait()
	glog.Flush()
}
