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
	r.POST("/info", c.info)
	r.POST("/delete", c.delete)
	r.POST("/modify/quantity", c.modifyQuantity)
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

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	item := &model.CartItem{
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

func (c *CartController) info(ctx *gin.Context) {
	var (
		req struct {
			UserID uint64 `json:"user_id"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	infos, err := sand.GetApplication().Services().ShoppingCarts().Query(req.UserID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "infos": infos})
}

func (s *CartController) delete(ctx *gin.Context) {
	var (
		req struct {
			UserID uint64   `json:"user_id"`
			ItemID []uint64 `json:"item_ids"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = sand.GetApplication().Services().ShoppingCarts().Delete(req.UserID, req.ItemID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (s *CartController) modifyQuantity(ctx *gin.Context) {
	var (
		req struct {
			UserID   uint64 `json:"user_id"`
			ItemID   uint64 `json:"item_id"`
			Quantity uint32 `json:"quantity"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = sand.GetApplication().Services().ShoppingCarts().ModifyQuantity(req.UserID, req.ItemID, req.Quantity)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
