package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) RegisterRouter(r gin.IRouter) {
	r.GET("/info", c.getUserInfo)
}

func (c *UserController) getUserInfo(ctx *gin.Context) {
	value := ctx.Query("open_id")
	if value == "" {
		ctx.Error(errors.New("[UserInfo] without open_id"))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	info, err := sand.Application.Services().Users().QueryByOpenID(value)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "info": info})
}
