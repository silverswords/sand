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
	r.POST("/list/parent-directory", c.listParentDirectories)
	r.POST("/list/sub-categories", c.listSubCategories)
	r.POST("/modify/status", c.modifyStatus)
	r.POST("/modify/name", c.modifyName)
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

func (c *CategoryController) listParentDirectories(ctx *gin.Context) {
	var (
		req struct {
			ID       uint64 `json:"id,omitempty"`
			Name     string `json:"name,omitempty"`
			ParentID uint64 `json:"parent_id,omitempty"`
			Status   uint8  `json:"status,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	categories, err := sand.Application.Services().Category().ListAllParentDirectory()
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": categories})
}

func (c *CategoryController) listSubCategories(ctx *gin.Context) {
	var (
		req struct {
			ParentID uint64 `json:"parent_id,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	categories, err := sand.Application.Services().Category().ListChildrenByParentID(req.ParentID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "result": categories})
}

func (c *CategoryController) modifyStatus(ctx *gin.Context) {
	var (
		req struct {
			ID     uint64 `json:"id,omitempty"`
			Status uint8  `json:"status,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := sand.Application.Services().Category().ModifyCategoryStatus(req.ID, req.Status); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (c *CategoryController) modifyName(ctx *gin.Context) {
	var (
		req struct {
			ID   uint64 `json:"id,omitempty"`
			Name string `json:"name,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err := sand.Application.Services().Category().ModifyCategoryName(req.ID, req.Name); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
