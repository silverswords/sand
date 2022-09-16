package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type ProductController struct {
}

func NewProductController() *ProductController {
	return &ProductController{}
}

func (c *ProductController) RegisterRouter(r gin.IRouter) {
	r.POST("/create", c.create)
}

func (c *ProductController) create(ctx *gin.Context) {
	var (
		req struct {
			StoreID    uint64  `json:"store_id,omitempty"`
			CategoryID uint64  `json:"category_id,omitempty"`
			Price      float64 `json:"price,omitempty"`
			PhotoUrls  string  `json:"photo_urls,omitempty"`
			MainTitle  string  `json:"main_title,omitempty"`
			Subtitle   string  `json:"subtitle,omitempty"`
			Status     uint8   `json:"status,omitempty"`
			Stock      uint32  `json:"stock,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	var product = &model.Product{
		StoreID:    req.StoreID,
		CategoryID: req.CategoryID,
		Price:      req.Price,
		PhotoUrls:  req.PhotoUrls,
		MainTitle:  req.MainTitle,
		Subtitle:   req.Subtitle,
		Status:     req.Status,
		Stock:      req.Stock,
	}

	if err := sand.Application.Services().Products().Create(product); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
