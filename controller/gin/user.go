package gin

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/services"
)

func RegisterUser(r gin.IRouter) {
	if err := services.CreateUserTable(); err != nil {
		log.Fatal(err)
	}
}
