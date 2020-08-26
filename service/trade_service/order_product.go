package trade_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistOrderProductByID(op model.OrderProduct) (bool, error) {
	if err := global.DB.Where("id = ?", op.ID).First(&op).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddOrderProduct(op model.OrderProduct) error {
	if err := global.DB.Create(&op).Error; err != nil {
		return err
	}

	return nil
}

func DeleteOrderProduct(op model.OrderProduct) error {
	if err := global.DB.Delete(model.OrderProduct{}, "id = ?", op.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteOrderProduct() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.OrderProduct{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateOrderProduct(op model.OrderProduct) (*model.OrderProduct, error) {
	if err := global.DB.Model(model.OrderProduct{}).Where("id = ?", op.ID).Updates(op).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", op.ID).First(&op).Error; err != nil {
		return nil, err
	}

	return &op, nil
}

func GetOrderProduct(op model.OrderProduct) (*model.OrderProduct, error) {
	if err := global.DB.Where("id = ?", op.ID).First(&op).Error; err != nil {
		return nil, err
	}

	return &op, nil
}

func GetOrderProducts(op model.OrderProduct, pageNum, pageSize int) ([]model.OrderProduct, error) {
	var ops []model.OrderProduct

	err := global.DB.Where(op).Offset(pageNum).Limit(pageSize).Find(&ops).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ops, nil
}

func GetOrderProductTotal(op model.OrderProduct) (int, error) {
	var count int

	if err := global.DB.Model(&model.OrderProduct{}).Where(op).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
