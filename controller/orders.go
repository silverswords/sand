package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type OrderController struct {
}

func (c *OrderController) RegisterRouter(r gin.IRouter) {
	r.POST("/create", c.create)
	r.POST("/delete", c.delete)
	r.POST("/infos", c.listOrdersByUserIDAndStatus)
	r.POST("/detail", c.detailByOrderID)
	r.POST("/modify/status", c.modifyStatus)
	r.POST("/modify/address", c.modifyAddress)
}

func (c *OrderController) create(ctx *gin.Context) {
	var (
		req struct {
			FromCart     bool                `json:"from_cart,omitempty"`
			UserID       uint64              `json:"user_id,omitempty"`
			TotalPrice   float64             `json:"total_price,omitempty"`
			Details      []model.OrderDetail `json:"details,omitempty"`
			Name         string              `json:"name,omitempty"`
			Phone        string              `json:"phone,omitempty"`
			ProvinceName string              `json:"province_name,omitempty"`
			CityName     string              `json:"city_name,omitempty"`
			CountyName   string              `json:"county_name,omitempty"`
			DetailInfo   string              `json:"detail_info,omitempty"`
		}
		orderDetails []*model.OrderDetail
		err          error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	order := &model.Order{
		UserID:     req.UserID,
		TotalPrice: req.TotalPrice,
	}

	address := &model.UserAddress{
		UserID:       req.UserID,
		UserName:     req.Name,
		UserPhone:    req.Phone,
		ProvinceName: req.ProvinceName,
		CityName:     req.CityName,
		CountyName:   req.CountyName,
		DetailInfo:   req.DetailInfo,
	}

	for _, detail := range req.Details {
		orderDetail := &model.OrderDetail{
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
			Price:     detail.Price,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	if err = sand.GetApplication().Services().Orders().Create(order, orderDetails, address, req.FromCart); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "order_id": order.ID})
}

func (c *OrderController) modifyStatus(ctx *gin.Context) {
	var (
		req struct {
			OrderID uint64 `json:"order_id,omitempty"`
			Status  uint8  `json:"status,omitempty"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	if err = sand.GetApplication().Services().Orders().ModifyStatus(req.OrderID, req.Status); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (c *OrderController) listOrdersByUserIDAndStatus(ctx *gin.Context) {
	var (
		req struct {
			UserID uint64 `json:"user_id"`
			Status uint8  `json:"status"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	orderInfos, err := sand.GetApplication().Services().Orders().QueryByUserIDAndStatus(req.UserID, req.Status)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "order_infos": orderInfos})
}

func (c *OrderController) detailByOrderID(ctx *gin.Context) {
	var (
		req struct {
			OrderID uint64 `json:"order_id"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	orderDetail, err := sand.GetApplication().Services().Orders().QueryDetailsByOrderID(req.OrderID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "order_detail": orderDetail})
}

func (c *OrderController) delete(ctx *gin.Context) {
	var (
		req struct {
			OrderID uint64 `json:"order_id"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	err = sand.GetApplication().Services().Orders().Delete(req.OrderID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}

func (c *OrderController) modifyAddress(ctx *gin.Context) {
	var (
		req struct {
			OrderID      uint64 `json:"order_id,omitempty"`
			Name         string `json:"name,omitempty"`
			Phone        string `json:"phone,omitempty"`
			ProvinceName string `json:"province_name,omitempty"`
			CityName     string `json:"city_name,omitempty"`
			CountyName   string `json:"county_name,omitempty"`
			DetailInfo   string `json:"detail_info,omitempty"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	var address = &model.UserAddress{
		UserName:     req.Name,
		UserPhone:    req.Phone,
		ProvinceName: req.ProvinceName,
		CityName:     req.CityName,
		CountyName:   req.CountyName,
		DetailInfo:   req.DetailInfo,
	}

	err = sand.GetApplication().Services().Orders().ModifyAddress(req.OrderID, address)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
