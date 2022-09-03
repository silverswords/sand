package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand/models"
)

type ProductController struct {
	db *sql.DB
}

func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{
		db: db,
	}
}

func (c *ProductController) RegisterProduct(r gin.IRouter) {
	if err := models.CreateProductTable(c.db); err != nil {
		log.Fatal(err)
	}

	r.GET("/insert", c.insertProduct)
	r.GET("/list", c.getAllProduct)
	r.GET("/detail", c.getProductInfoByID)
	r.GET("/virtualStorePros", c.getVirtualStorePros)
}

func (c *ProductController) insertProduct(ctx *gin.Context) {
	var (
		req struct {
			ProductID string      `json:"pro_id,omitempty"`
			StoreID   string      `json:"store_id,omitempty"`
			Price     float64     `json:"price,omitempty"`
			MainTitle string      `json:"main_title,omitempty"`
			Subtitle  string      `json:"subtitle,omitempty"`
			Images    interface{} `json:"images,omitempty"`
			Stock     uint32      `json:"stock,omitempty"`
			Status    uint8       `json:"status,omitempty"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := models.InsertProduct(c.db, req.ProductID, req.StoreID, req.Price, req.MainTitle, req.Subtitle, req.Images, req.Stock, req.Status, time.Now()); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (c *ProductController) getAllProduct(ctx *gin.Context) {
	brifeInfo, err := models.ListAllProduce(c.db)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}

func (c *ProductController) getProductInfoByID(ctx *gin.Context) {
	var (
		req struct {
			ProductID string `json:"product_id"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	detial, err := models.GetProductInfoByID(c.db, req.ProductID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "detial": detial})
}

func (c *ProductController) getVirtualStorePros(ctx *gin.Context) {
	var (
		req struct {
			VirtualStoreID string `json:"virtual_store_id"`
		}
	)

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	brifeInfo, err := models.GetVirtualStoreProsByID(c.db, req.VirtualStoreID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "brifeInfo": brifeInfo})
}
