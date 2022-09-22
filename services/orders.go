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
	Details []*model.OrderDetail `json:"details"`
}

func CreateOrdersService(accessor interfaces.DatabaseAccessor) Orders {
	return &orders{
		DatabaseAccessor: accessor,
	}
}

func (s *orders) Create(o *model.Order, d []*model.OrderDetail, fromCart bool) error {
	return s.GetDefaultGormDB().Transaction(func(tx *gorm.DB) error {
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

func (s *orders) Modify(o *model.Order) error {
	return s.GetDefaultGormDB().Model(model.Order{}).Where("id = ?", o.ID).Updates(o).Error
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
			Order:   order,
			Details: details,
		}

		orderInfos = append(orderInfos, orderInfo)
	}

	return orderInfos, nil
}

func (s *orders) QueryDetailsByOrderID(orderID uint64) ([]*model.OrderDetail, error) {
	var details []*model.OrderDetail
	err := s.GetDefaultGormDB().Model(model.OrderDetail{}).Where("order_id = ?", orderID).
		Find(&details).Error

	return details, err
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
