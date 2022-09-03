package controller

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/models"
)

type OrderController struct {
	db *sql.DB
}

func NewOrderController(db *sql.DB) *OrderController {
	return &OrderController{
		db: db,
	}
}

func (c *OrderController) Register(r gin.IRouter) {
	if err := models.CreateOrderTable(c.db); err != nil {
		log.Fatal(err)
	}

	r.POST("/insert", c.insert)
	r.POST("/modify/status", c.modifyOrderStatus)
	r.GET("/brifeInfoByOpenID", c.getOrderBrifeInfoByOpenID)
	r.GET("/brifeInfoByStoreID", c.getOrderBrifeInfoByStoreID)
	r.GET("/detail", c.getOrderDetialByOrderID)
}

func (c *OrderController) insert(ctx *gin.Context) {
	var (
		req struct {
			OrderID    uint32  `json:"order_id,omitempty"`
			UserID     uint64  `json:"open_id,omitempty"`
			ProductID  string  `json:"pro_id,omitempty"`
			StoreID    string  `json:"store_id,omitempty"`
			Quantity   uint32  `json:"count,omitempty"`
			TotalPrice float64 `json:"total_price,omitempty"`
			Status     uint8   `json:"status,omitempty"`
			CreateTime string  `json:"create_time"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := models.InsertOrder(c.db, req.OrderID, req.UserID, req.ProductID, req.StoreID, req.Quantity, req.TotalPrice, req.Status); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (c *OrderController) getOrderBrifeInfoByOpenID(ctx *gin.Context) {
	var openId uint64

	if err := ctx.ShouldBind(&openId); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	brifeInfo, err := models.ListOrderByUserID(c.db, openId)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}

func (c *OrderController) getOrderBrifeInfoByStoreID(ctx *gin.Context) {
	var store_id string

	if err := ctx.ShouldBind(&store_id); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	brifeInfo, err := models.ListOrderByStoreID(c.db, store_id)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}

func (c *OrderController) getOrderDetialByOrderID(ctx *gin.Context) {
	var order_id string

	if err := ctx.ShouldBind(&order_id); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	detial, err := models.GetOrderDetialByOrderID(c.db, order_id)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "detial": detial})
}

func (c *OrderController) modifyOrderStatus(ctx *gin.Context) {
	var (
		req struct {
			OrderID string `json:"order_id,omitempty"`
			Status  uint8  `json:"status,omitempty"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := models.ModifyOrderStatus(c.db, req.OrderID, req.Status); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
