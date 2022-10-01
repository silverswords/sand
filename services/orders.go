package services

import (
	"errors"
	"fmt"

	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
	"gorm.io/gorm"
)

type orders struct {
	interfaces.DatabaseAccessor
}

type address struct {
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	ProvinceName string `json:"province_name"`
	CityName     string `json:"city_name"`
	CountyName   string `json:"county_name"`
	DetailInfo   string `json:"detail_info"`
}

type productItem struct {
	ID        uint64 `json:"id"`
	MainTitle string `json:"main_title"`
	Spec      string `json:"spec"`
	Price     uint32 `json:"price"`
	PhotoUrls string `json:"photo_urls"`
	Quantity  uint32 `json:"quantity"`
}

type orderInfo struct {
	ID         uint64         `json:"id"`
	Status     uint8          `json:"status"`
	TotalPrice uint32         `json:"total_price"`
	Products   []*productItem `json:"products"`
}

type orderDetail struct {
	*orderInfo
	Address *address `json:"address"`
}

func CreateOrdersService(accessor interfaces.DatabaseAccessor) Orders {
	return &orders{
		DatabaseAccessor: accessor,
	}
}

func (s *orders) OrdersCreate(o *model.Order, d []*model.OrderDetail, a *model.UserAddress, fromCart bool) error {
	return s.GetDefaultGormDB().Transaction(func(tx *gorm.DB) error {
		var (
			addresses []*model.UserAddress
			isExist   bool
			err       error
		)

		err = tx.Model(model.UserAddress{}).Where("user_id = ?", a.UserID).Find(&addresses).Error
		if err != nil {
			return err
		}

		for _, address := range addresses {
			if address.UserName == a.UserName && address.UserPhone == a.UserPhone &&
				address.ProvinceName == a.ProvinceName && address.CityName == a.CityName &&
				address.CountyName == a.CountyName && address.DetailInfo == a.DetailInfo {
				isExist = true
			}
		}

		if !isExist {
			err = tx.Model(model.UserAddress{}).Create(a).Error
			if err != nil {
				return err
			}
		}

		o.UserAddressID = a.ID
		if err := tx.Model(model.Order{}).Create(o).Error; err != nil {
			return err
		}

		for _, d := range d {
			d.OrderID = o.ID
		}

		if err := tx.Model(model.OrderDetail{}).Create(d).Error; err != nil {
			return err
		}

		for _, detail := range d {
			var (
				product *model.Product
			)

			err := tx.Model(model.Product{}).Where("id = ?", detail.ProductID).Take(&product).Error
			if err != nil {
				return err
			}

			if product.Stock < detail.Quantity {
				return fmt.Errorf("product[%d] stock is not enough", detail.ProductID)
			}

			err = tx.Model(model.Product{}).Where("id = ?", detail.ProductID).
				Update("stock", (product.Stock - detail.Quantity)).Error
			if err != nil {
				return err
			}

			if fromCart {
				err = tx.Model(model.CartItem{}).Where("product_id = ? AND quantity = ?", detail.ProductID, detail.Quantity).
					Delete(&model.CartItem{}).Error
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
}

func (s *orders) OrdersQueryByUserIDAndStatus(userID uint64, status uint8) ([]*orderInfo, error) {
	var (
		orders     []*model.Order
		orderInfos []*orderInfo
	)

	err := s.GetDefaultGormDB().Model(model.Order{}).Where("user_id = ? And status = ?", userID, status).
		Order("created_at desc").Find(&orders).Error
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		var (
			details  []*model.OrderDetail
			products []*productItem
		)

		err := s.GetDefaultGormDB().Model(model.OrderDetail{}).Where("order_id = ?", order.ID).
			Find(&details).Error

		if err != nil {
			return nil, err
		}

		for _, detail := range details {
			var (
				product *model.Product
				item    *productItem
			)

			err := s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", detail.ProductID).
				Find(&product).Error

			if err != nil {
				return nil, err
			}

			item = &productItem{
				ID:        product.ID,
				MainTitle: product.MainTitle,
				Spec:      product.Spec,
				Price:     product.Price,
				PhotoUrls: product.PhotoUrls,
				Quantity:  detail.Quantity,
			}

			products = append(products, item)
		}

		info := &orderInfo{
			ID:         order.ID,
			TotalPrice: order.TotalPrice,
			Status:     order.Status,
			Products:   products,
		}
		orderInfos = append(orderInfos, info)
	}

	return orderInfos, nil
}

func (s *orders) OrdersQueryDetailsByOrderID(userID uint64, orderID uint64) (*orderDetail, error) {
	var (
		order        *model.Order
		orders       []*model.Order
		orderDetails []*model.OrderDetail
		products     []*productItem
		addr         *model.UserAddress
		orderIsExist bool
		err          error
	)

	err = s.GetDefaultGormDB().Model(model.Order{}).Where("user_id = ?", userID).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		if order.ID == orderID {
			orderIsExist = true
		}
	}

	if !orderIsExist {
		return nil, errors.New("order does not exist for this user")
	}

	err = s.GetDefaultGormDB().Model(model.Order{}).Where("id = ?", orderID).Take(&order).Error
	if err != nil {
		return nil, err
	}

	err = s.GetDefaultGormDB().Model(model.OrderDetail{}).Where("order_id = ?", orderID).
		Find(&orderDetails).Error
	if err != nil {
		return nil, err
	}

	err = s.GetDefaultGormDB().Model(model.UserAddress{}).Where("id = ?", order.UserAddressID).
		Take(&addr).Error
	if err != nil {
		return nil, err
	}

	for _, detail := range orderDetails {
		var (
			product *model.Product
			item    *productItem
		)

		err = s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", detail.ProductID).
			Take(&product).Error
		if err != nil {
			return nil, err
		}

		item = &productItem{
			ID:        product.ID,
			MainTitle: product.MainTitle,
			Spec:      product.Spec,
			Price:     product.Price,
			PhotoUrls: product.PhotoUrls,
			Quantity:  detail.Quantity,
		}

		products = append(products, item)
	}

	info := &orderInfo{
		ID:         order.ID,
		TotalPrice: order.TotalPrice,
		Status:     order.Status,
		Products:   products,
	}

	address := &address{
		Name:         addr.UserName,
		Phone:        addr.UserPhone,
		ProvinceName: addr.ProvinceName,
		CityName:     addr.CityName,
		CountyName:   addr.CountyName,
		DetailInfo:   addr.DetailInfo,
	}

	detail := &orderDetail{
		orderInfo: info,
		Address:   address,
	}

	return detail, err
}

func (s *orders) OrdersModifyStatus(userID uint64, orderID uint64, status uint8) error {
	result := s.GetDefaultGormDB().Model(model.Order{}).Where("user_id = ? AND id = ?", userID, orderID).
		Update("status", status)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *orders) OrdersModifyAddress(userID uint64, orderID uint64, a *model.UserAddress) error {
	var order *model.Order

	err := s.GetDefaultGormDB().Model(model.Order{}).Where("id = ?", orderID).Take(&order).Error
	if err != nil {
		return err
	}

	return s.GetDefaultGormDB().Model(model.UserAddress{}).Where("id = ?", order.UserAddressID).Updates(a).Error
}

func (s *orders) OrdersDelete(userID uint64, orderID uint64) error {
	return s.GetDefaultGormDB().Transaction(func(tx *gorm.DB) error {
		var (
			orderDetails []*model.OrderDetail
			err          error
		)

		result := tx.Model(model.Order{}).Where("user_id = ?", userID).Delete(model.Order{}, orderID)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return errInvalidNoRowsAffected
		}

		err = tx.Model(model.OrderDetail{}).Where("order_id = ?", orderID).Find(&orderDetails).Error
		if err != nil {
			return err
		}

		err = tx.Model(model.OrderDetail{}).Where("order_id = ?", orderID).Delete(&model.OrderDetail{}).Error
		if err != nil {
			return err
		}

		for _, orderDetail := range orderDetails {
			var (
				product *model.Product
			)

			err = tx.Model(model.Product{}).Where("id", orderDetail.ProductID).Take(&product).Error
			if err != nil {
				return err
			}

			err = tx.Model(model.Product{}).Where("id = ?", orderDetail.ProductID).
				Update("stock", (product.Stock + orderDetail.Quantity)).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}
