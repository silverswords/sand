package services

import (
	"github.com/silverswords/sand/core/interfaces"
	"github.com/silverswords/sand/model"
)

type carts struct {
	interfaces.DatabaseAccessor
}

type itemInfo struct {
	MainTitle  string `json:"main_title"`
	Price      uint32 `json:"price"`
	Quantity   uint32 `json:"quantity"`
	PhotoUrls  string `json:"photo_urls"`
	TotalPrice uint32 `json:"total_price"`
}

func CreateCartsService(accessor interfaces.DatabaseAccessor) Carts {
	return &carts{
		DatabaseAccessor: accessor,
	}
}

func (s *carts) CartsCreate(sc *model.CartItem) error {
	return s.GetDefaultGormDB().Model(model.CartItem{}).Create(sc).Error
}

func (s *carts) CartsQuery(user_id uint64) ([]*itemInfo, error) {
	var (
		cartItems []*model.CartItem
		product   *model.Product
		itemInfos []*itemInfo
	)

	err := s.GetDefaultGormDB().Model(model.CartItem{}).Where("user_id = ?", user_id).
		Order("updated_at desc").Find(&cartItems).Error
	if err != nil {
		return nil, err
	}

	for _, cartItem := range cartItems {
		err = s.GetDefaultGormDB().Model(model.Product{}).Where("id = ?", cartItem.ProductID).
			Take(&product).Error
		if err != nil {
			return nil, err
		}

		info := &itemInfo{
			MainTitle:  product.MainTitle,
			Price:      product.Price,
			Quantity:   cartItem.Quantity,
			PhotoUrls:  product.PhotoUrls,
			TotalPrice: product.Price * cartItem.Quantity,
		}

		itemInfos = append(itemInfos, info)
	}

	return itemInfos, nil
}

func (s *carts) CartsDelete(userID uint64, itemIDs []uint64) error {
	result := s.GetDefaultGormDB().Model(model.CartItem{}).Where("user_id = ?", userID).
		Delete(&model.CartItem{}, &itemIDs)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}

func (s *carts) CartsModifyQuantity(userID uint64, itemID uint64, quantity uint32) error {
	result := s.GetDefaultGormDB().Model(model.CartItem{}).Where("user_id = ? AND id = ?", userID, itemID).
		Update("quantity", quantity)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errInvalidNoRowsAffected
	}

	return nil
}
