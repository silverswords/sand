package web

import (
	"io/ioutil"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"github.com/silverswords/sand/controller"
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

func (b *Builder) Run() *Server {
	listener, err := net.Listen("tcp", b.Host+":"+b.Config.Addr)
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	registerRouter(engine)

	return &Server{
		Engine: engine,

		server: http.Server{
			Handler: engine,
		},

		listener: listener,
	}
}

func registerRouter(engine *gin.Engine) {
	routerBasicGroup := engine.Group("/api/v1")

	order := &controller.OrderController{}
	order.RegisterRouter(routerBasicGroup.Group("order"))
	user := &controller.UserController{}
	user.RegisterRouter(routerBasicGroup.Group("user"))
	product := &controller.ProductController{}
	product.RegisterRouter(routerBasicGroup.Group("product"))
	cart := &controller.CartController{}
	cart.RegisterRouter(routerBasicGroup.Group("cart"))
}
