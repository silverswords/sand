package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type CartController struct {
}

func (c *CartController) RegisterRouter(r gin.IRouter) {
	r.POST("/create", c.create)
}

func (c *CartController) create(ctx *gin.Context) {
	var (
		req struct {
			UserID    uint64 `json:"user_id"`
			ProductID uint64 `json:"product_id"`
			Quantity  uint32 `json:"quantity"`
		}
		err error
	)

	err = ctx.ShouldBind(req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	item := &model.ShoppingCart{
		UserID:    req.UserID,
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	err = sand.GetApplication().Services().ShoppingCarts().Create(item)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
