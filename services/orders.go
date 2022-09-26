package services

import (
	"fmt"

	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
	"gorm.io/gorm"
)

type orders struct {
	interfaces.DatabaseAccessor
}

type orderInfo struct {
	*model.Order
	*orderDetail
}

type orderDetail struct {
	Details []*model.OrderDetail `json:"details"`
	Address *model.UserAddress   `json:"address"`
}

func CreateOrdersService(accessor interfaces.DatabaseAccessor) Orders {
	return &orders{
		DatabaseAccessor: accessor,
	}
}

func (s *orders) Create(o *model.Order, d []*model.OrderDetail, a *model.UserAddress, fromCart bool) error {
	return s.GetDefaultGormDB().Transaction(func(tx *gorm.DB) error {
		var (
			addresses []*model.UserAddress
			isExist   bool
			err       error
		)

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

			err = tx.Model(model.UserAddress{}).Where("user_id = ?", a.UserID).Find(&addresses).Error
			if err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *orders) QueryByUserIDAndStatus(userID uint64, status uint8) ([]*orderInfo, error) {
	var (
		orders     []*model.Order
		orderInfos []*orderInfo
	)
	err := s.GetDefaultGormDB().Model(model.Order{}).
		Where("user_id = ? And status = ?", userID, status).Order("created_at").Find(&orders).Error

	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		details, err := s.QueryDetailsByOrderID(order.ID)
		if err != nil {
			return nil, err
		}

		orderInfo := &orderInfo{
			Order:       order,
			orderDetail: details,
		}

		orderInfos = append(orderInfos, orderInfo)
	}

	return orderInfos, nil
}

func (s *orders) QueryDetailsByOrderID(orderID uint64) (*orderDetail, error) {
	var (
		detail = &orderDetail{}
		order  *model.Order
	)

	err := s.GetDefaultGormDB().Model(model.Order{}).Where("id = ?", orderID).Take(&order).Error
	if err != nil {
		return nil, err
	}

	err = s.GetDefaultGormDB().Model(model.OrderDetail{}).Where("order_id = ?", orderID).
		Find(&detail.Details).Error
	if err != nil {
		return nil, err
	}

	err = s.GetDefaultGormDB().Model(model.UserAddress{}).Where("id = ?", order.UserAddressID).
		Take(&detail.Address).Error
	if err != nil {
		return nil, err
	}

	return detail, err
}

func (s *orders) ModifyStatus(id uint64, status uint8) error {
	return s.GetDefaultGormDB().Model(model.Order{}).Where("id = ?", id).Update("status", status).Error
}

func (s *orders) ModifyAddress(orderID uint64, a *model.UserAddress) error {
	var order *model.Order

	err := s.GetDefaultGormDB().Model(model.Order{}).Where("id = ?", orderID).Take(&order).Error
	if err != nil {
		return err
	}

	return s.GetDefaultGormDB().Model(model.UserAddress{}).Where("id = ?", order.UserAddressID).Updates(a).Error
}

func (s *orders) Delete(orderID uint64) error {
	return s.GetDefaultGormDB().Transaction(func(tx *gorm.DB) error {
		var (
			orderDetails []*model.OrderDetail
		)

		err := tx.Model(model.Order{}).Delete(model.Order{}, orderID).Error
		if err != nil {
			return err
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
