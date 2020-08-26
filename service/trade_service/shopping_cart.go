package trade_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistShoppingCartByID(sc model.ShoppingCart) (bool, error) {
	if err := global.DB.Where("id = ?", sc.ID).First(&sc).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddShoppingCart(sc model.ShoppingCart) (*model.ShoppingCart, error) {
	if err := global.DB.Create(&sc).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("user_id = ? AND product_id = ?", sc.UserID, sc.ProductID).First(&sc).Error; err != nil {
		return nil, err
	}

	return &sc, nil
}

func DeleteShoppingCart(sc model.ShoppingCart) error {
	if err := global.DB.Delete(model.ShoppingCart{}, "id = ?", sc.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteShoppingCart() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.ShoppingCart{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateShoppingCart(sc model.ShoppingCart) (*model.ShoppingCart, error) {
	if err := global.DB.Model(model.ShoppingCart{}).Where("id = ?", sc.ID).Updates(sc).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", sc.ID).First(&sc).Error; err != nil {
		return nil, err
	}

	return &sc, nil
}

func GetShoppingCart(sc model.ShoppingCart) (*model.ShoppingCart, error) {
	if err := global.DB.Where("id = ?", sc.ID).First(&sc).Error; err != nil {
		return nil, err
	}

	return &sc, nil
}

func GetShoppingCarts(sc model.ShoppingCart, pageNum, pageSize int) ([]model.ShoppingCart, error) {
	var scs []model.ShoppingCart

	err := global.DB.Where(sc).Offset(pageNum).Limit(pageSize).Find(&scs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return scs, nil
}

func GetShoppingCartTotal(sc model.ShoppingCart) (int, error) {
	var count int

	if err := global.DB.Model(&model.ShoppingCart{}).Where(sc).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
