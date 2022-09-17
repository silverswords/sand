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
	r.POST("/detial", c.detial)
	r.POST("/list/default", c.listDefault)
	r.POST("/list/by-category-id", c.listByCategoryID)
	r.POST("/list/by-store-id", c.listByStoreID)
	r.POST("/modify/property", c.modifyProperty)
	r.POST("/modify/category-id", c.modifyCategoryID)
	r.POST("/modify/store-id", c.modifyStoreID)
	r.POST("/modify/status", c.modifyStatus)
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

func (c *ProductController) detial(ctx *gin.Context) {
	var (
		req struct {
			ProductID uint64 `json:"product_id,omitempty"`
		}
	)
	ctx.ShouldBind(&req)

	product, err := sand.Application.Services().Products().QueryDetialByProductID(req.ProductID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": product})
}

func (c *ProductController) listDefault(ctx *gin.Context) {
	products, err := sand.Application.Services().Products().ListAllProducts()
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "retult": products})
}

func (c *ProductController) listByCategoryID(ctx *gin.Context) {
	var (
		req struct {
			CategoryID uint64 `json:"category_id,omitempty"`
		}
	)

	products, err := sand.Application.Services().Products().ListByCategoryID(req.CategoryID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "retult": products})
}

func (c *ProductController) listByStoreID(ctx *gin.Context) {
	var (
		req struct {
			StoreID uint64 `json:"store_id,omitempty"`
		}
	)

	products, err := sand.Application.Services().Products().ListByStoreId(req.StoreID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "retult": products})
}

func (c *ProductController) modifyProperty(ctx *gin.Context) {

}

func (c *ProductController) modifyCategoryID(ctx *gin.Context) {

}

func (c *ProductController) modifyStoreID(ctx *gin.Context) {

}

func (c *ProductController) modifyStatus(ctx *gin.Context) {

}
