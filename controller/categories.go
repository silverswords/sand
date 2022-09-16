package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type CategoryController struct {
}

func NewcategoryController() *CategoryController {
	return &CategoryController{}
}

func (c *CategoryController) RegisterRouter(r gin.IRouter) {
	r.POST("/create", c.create)
}

func (c *CategoryController) create(ctx *gin.Context) {
	var (
		req struct {
			ParentID uint64 `json:"parent_id,omitempty"`
			Name     string `json:"name,omitempty"`
			Status   uint8  `json:"status,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	var category = &model.Category{
		ParentID: req.ParentID,
		Name:     req.Name,
		Status:   req.Status,
	}

	if err := sand.Application.Services().Category().Create(category); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
