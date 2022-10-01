package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/silverswords/sand"
	"github.com/silverswords/sand/model"
)

type orderDetail struct {
	ProductID uint64  `json:"product_id"`
	Quantity  uint32  `json:"quantity"`
	Price     float64 `json:"price"`
}

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
			FromCart     bool          `json:"from_cart,omitempty"`
			UserID       uint64        `json:"user_id,omitempty"`
			TotalPrice   uint32        `json:"total_price,omitempty"`
			Name         string        `json:"name,omitempty"`
			Phone        string        `json:"phone,omitempty"`
			ProvinceName string        `json:"province_name,omitempty"`
			CityName     string        `json:"city_name,omitempty"`
			CountyName   string        `json:"county_name,omitempty"`
			DetailInfo   string        `json:"detail_info,omitempty"`
			Description  string        `json:"description,omitempty"`
			Details      []orderDetail `json:"details,omitempty"`
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

	if err = sand.GetApplication().Services().OrdersCreate(order, orderDetails, address, req.FromCart); err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	user, err := sand.GetApplication().Services().UsersQueryByID(req.UserID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	prepayID, appID, err := sand.GetApplication().Services().GetPrepayInfo(req.Description,
		string(order.ID), int(order.TotalPrice), user.OpenID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	prepayInfo, err := sand.GetApplication().Services().GetSignedInfo(prepayID, appID)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "prepay_info": prepayInfo})
}

func (c *OrderController) modifyStatus(ctx *gin.Context) {
	var (
		req struct {
			UserID  uint64 `json:"user_id"`
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

	if err = sand.GetApplication().Services().OrdersModifyStatus(req.UserID, req.OrderID, req.Status); err != nil {
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

	orderInfos, err := sand.GetApplication().Services().OrdersQueryByUserIDAndStatus(req.UserID, req.Status)
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
			UserID  uint64 `json:"user_id"`
			OrderID uint64 `json:"order_id"`
		}
	)

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	orderDetail, err := sand.GetApplication().Services().OrdersQueryDetailsByOrderID(req.UserID, req.OrderID)
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
			UserID  uint64 `json:"user_id"`
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

	err = sand.GetApplication().Services().OrdersDelete(req.UserID, req.OrderID)
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
			UserID       uint64 `json:"user_id"`
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

	err = sand.GetApplication().Services().OrdersModifyAddress(req.UserID, req.OrderID, address)
	if err != nil {
		ctx.Error(err)
		ctx.JSON(http.StatusBadGateway, gin.H{"status": http.StatusBadGateway})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": http.StatusOK})
}
