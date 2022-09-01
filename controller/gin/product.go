package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/models/structs"
	"github.com/silverswords/sand/services"
)

func RegisterProduct(r gin.IRouter) {
	if err := services.CreateProductTable(); err != nil {
		log.Fatal(err)
	}

	r.GET("/insert", insertProduct)
	r.GET("/getAll", getAllProduct)
	r.GET("/getInfoByID", getProductInfoByID)
	r.GET("/getVirtualStorePros", getVirtualStorePros)
}

func insertProduct(ctx *gin.Context) {
	var product structs.Product
	if err := ctx.ShouldBind(&product); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := services.InsertProduct(product); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func getAllProduct(ctx *gin.Context) {
	brifeInfo, err := services.GetAllProduce()
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}

func getProductInfoByID(ctx *gin.Context) {
	var productID string

	if err := ctx.ShouldBind(&productID); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	detial, err := services.GetProductInfoByID(productID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "detial": detial})
}

func getVirtualStorePros(ctx *gin.Context) {
	var virtualStoreID string

	if err := ctx.ShouldBind(&virtualStoreID); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	brifeInfo, err := services.GetVirtualStoreProsByID(virtualStoreID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}
