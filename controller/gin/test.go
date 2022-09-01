package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/services"
)

func RegisterTest(r gin.IRouter) {
	r.GET("", helloworld)
}

func helloworld(ctx *gin.Context) {
	result := services.Helloworld()

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": result})
}
