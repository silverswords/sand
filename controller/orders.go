package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type OrderController struct {
}

func NewOrderController() *OrderController {
	return &OrderController{}
}

func (c *OrderController) RegisterRouter(r gin.IRouter) {
	r.POST("/create", c.create)
}

func (c *OrderController) create(ctx *gin.Context) {
	var (
		req struct {
			UserID        uint64  `json:"user_id,omitempty"`
			ProductID     uint64  `json:"product_id,omitempty"`
			UserAddressID uint64  `json:"user_address_id,omitempty"`
			TotalPrice    float64 `json:"total_price,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	var order = &model.Order{
		UserID:        req.UserID,
		UserAddressID: req.UserAddressID,
		TotalPrice:    req.TotalPrice,
	}

	if err := sand.Application.Services().Orders().Create(order); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
