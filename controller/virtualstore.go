package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type VirtualStoreController struct {
}

func NewVirtualStoreController() *VirtualStoreController {
	return &VirtualStoreController{}
}

func (c *VirtualStoreController) RegisterRouter(r gin.IRouter) {
	r.POST("/create", c.create)
}

func (c *VirtualStoreController) create(ctx *gin.Context) {
	var (
		req struct {
			Name   string `json:"name,omitempty"`
			Status uint8  `json:"status,omitempty"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	var virtualStore = &model.VirtualStore{
		Name:   req.Name,
		Status: req.Status,
	}

	if err := sand.Application.Services().VirtualStore().Create(virtualStore); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
