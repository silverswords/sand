package controller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type UserController struct {
}

func (c *UserController) RegisterRouter(r gin.IRouter) {
	r.POST("/login", c.login)
	r.GET("/info", c.getUserInfo)
}

func (c *UserController) login(ctx *gin.Context) {
	var (
		req struct {
			Code string `json:"code"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	loginResp, err := sand.GetApplication().Services().WeChat().Login(req.Code)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	_, err = sand.GetApplication().Services().Users().QueryByOpenID(loginResp.OpenID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
	}

	user := &model.User{UnionID: loginResp.UnionID, OpenID: loginResp.OpenID}
	err = sand.GetApplication().Services().Users().Create(user)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (c *UserController) getUserInfo(ctx *gin.Context) {
	value := ctx.Query("open_id")
	if value == "" {
		ctx.Error(errors.New("[UserInfo] without open_id"))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	info, err := sand.GetApplication().Services().Users().QueryByOpenID(value)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "info": info})
}
