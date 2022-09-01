package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/models/structs"
	"github.com/silverswords/sand/services"
)

func RegisterOrder(r gin.IRouter) {
	if err := services.CreateOrderTable(); err != nil {
		log.Fatal(err)
	}

	r.POST("/insert", insertOrder)
	r.POST("/modify/status", modifyOrderStatus)
	r.GET("/brifeInfoByOpenID", getOrderBrifeInfoByOpenID)
	r.GET("/brifeInfoByStoreID", getOrderBrifeInfoByStoreID)
	r.GET("/detail", getOrderDetialByOrderID)
}

func insertOrder(ctx *gin.Context) {
	var order structs.Order

	if err := ctx.ShouldBind(&order); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := services.InsertOrder(order); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func getOrderBrifeInfoByOpenID(ctx *gin.Context) {
	var open_id string

	if err := ctx.ShouldBind(&open_id); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	brifeInfo, err := services.GetOrderBrifeInfoByOpenID(open_id)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}

func getOrderBrifeInfoByStoreID(ctx *gin.Context) {
	var store_id string

	if err := ctx.ShouldBind(&store_id); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	brifeInfo, err := services.GetOrderBrifeInfoByStoreID(store_id)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}

func getOrderDetialByOrderID(ctx *gin.Context) {
	var order_id string

	if err := ctx.ShouldBind(&order_id); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	detial, err := services.GetOrderDetialByOrderID(order_id)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "detial": detial})
}

func modifyOrderStatus(ctx *gin.Context) {
	var order_status *structs.OrderStatus

	if err := ctx.ShouldBind(&order_status); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := services.ModifyOrderStatus(order_status.OrderID, order_status.Status); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
