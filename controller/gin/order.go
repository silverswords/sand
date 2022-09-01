package gin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/services"
)

func RegisterOrder(r gin.IRouter) {
	if err := services.CreateOrderTable(); err != nil {
		log.Fatal(err)
	}

	r.GET("")
}
