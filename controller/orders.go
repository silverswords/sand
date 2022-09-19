package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type OrderController struct {
}

type OrderInfo struct {
	*model.Order
	Details []*model.OrderDetail
}

func (c *OrderController) RegisterRouter(r gin.IRouter) {
	r.POST("/create", c.create)
	r.POST("/modify/status", c.modifyStatus)
	r.POST("/info", c.listOrdersByUserIDAndStatus)
}

func (c *OrderController) create(ctx *gin.Context) {
	var (
		req struct {
			UserID        uint64              `json:"user_id,omitempty"`
			ProductID     uint64              `json:"product_id,omitempty"`
			UserAddressID uint64              `json:"user_address_id,omitempty"`
			TotalPrice    float64             `json:"total_price,omitempty"`
			Details       []model.OrderDetail `json:"details,omitempty"`
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
		UserID:        req.UserID,
		UserAddressID: req.UserAddressID,
		TotalPrice:    req.TotalPrice,
	}

	if err = sand.GetApplication().Services().Orders().Create(order); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	for _, detail := range req.Details {
		orderDetail := &model.OrderDetail{
			OrderID:   order.ID,
			ProductID: detail.ProductID,
			Quantity:  detail.Quantity,
			Price:     detail.Price,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	if err = sand.GetApplication().Services().OrderDetails().Create(orderDetails); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "order_id": order.ID})
}

func (c *OrderController) modifyStatus(ctx *gin.Context) {
	var (
		req struct {
			OrderID uint64 `json:"order_id"`
			Status  uint8  `json:"status"`
		}
		err error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	order := &model.Order{
		Status: req.Status,
	}
	order.ID = req.OrderID

	if err = sand.GetApplication().Services().Orders().Modify(order); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"statue": http.StatusOK})
}

func (c *OrderController) listOrdersByUserIDAndStatus(ctx *gin.Context) {
	var (
		req struct {
			UserID uint64 `json:"user_id"`
			Status uint8  `json:"status"`
		}
		orderInfos []OrderInfo
		err        error
	)

	err = ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	orders, err := sand.GetApplication().Services().Orders().QueryByUserIDAndStatus(req.UserID, req.Status)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	for _, order := range orders {
		var orderInfo OrderInfo
		orderInfo.Order = order
		details, err := sand.GetApplication().Services().OrderDetails().QueryByOrderID(order.ID)
		if err != nil {
			ctx.Error(err)
			ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
			return
		}

		orderInfo.Details = append(orderInfo.Details, details...)

		orderInfos = append(orderInfos, orderInfo)
	}

	ctx.JSON(http.StatusOK, gin.H{"statue": http.StatusOK, "order_infos": orderInfos})
}
