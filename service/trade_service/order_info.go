package trade_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistOrderInfoByID(oi model.OrderInfo) (bool, error) {
	if err := global.DB.Where("id = ?", oi.ID).First(&oi).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddOrderInfo(oi model.OrderInfo) (*model.OrderInfo, error) {
	if err := global.DB.Create(&oi).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Preload("User").Where("order_no = ?", oi.OrderNo).First(&oi).Error; err != nil {
		return nil, err
	}
	// 获取订单中的商品
	var orderProducts []model.OrderProduct
	if err := global.DB.Preload("Product").Where("order_info_id = ?", oi.ID).Find(&orderProducts).
		Error; err != nil {
		return nil, err
	}
	oi.OrderProducts = orderProducts

	return &oi, nil
}

func DeleteOrderInfo(oi model.OrderInfo) error {
	if err := global.DB.Delete(model.OrderInfo{}, "id = ?", oi.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteOrderInfo() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.OrderInfo{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateOrderInfo(oi model.OrderInfo) (*model.OrderInfo, error) {
	if err := global.DB.Model(model.OrderInfo{}).Where("id = ?", oi.ID).Updates(oi).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", oi.ID).First(&oi).Error; err != nil {
		return nil, err
	}

	return &oi, nil
}

func GetOrderInfo(oi model.OrderInfo) (*model.OrderInfo, error) {
	if err := global.DB.Preload("User").Where("id = ?", oi.ID).First(&oi).Error; err != nil {
		return nil, err
	}

	// 获取订单中的商品
	var orderProducts []model.OrderProduct
	if err := global.DB.Preload("Product").Where("order_info_id = ?", oi.ID).Find(&orderProducts).
		Error; err != nil {
		return nil, err
	}
	oi.OrderProducts = orderProducts

	return &oi, nil
}

func GetOrderInfoByOrderNo(oi model.OrderInfo) (*model.OrderInfo, error) {
	if err := global.DB.Preload("User").Where("order_no = ?", oi.OrderNo).First(&oi).Error; err != nil {
		return nil, err
	}

	// 获取订单中的商品
	var orderProducts []model.OrderProduct
	if err := global.DB.Preload("Product").Where("order_info_id = ?", oi.ID).Find(&orderProducts).
		Error; err != nil {
		return nil, err
	}
	oi.OrderProducts = orderProducts

	return &oi, nil
}

func GetOrderInfos(oi model.OrderInfo, pageNum, pageSize int) ([]model.OrderInfo, error) {
	var ois []model.OrderInfo

	err := global.DB.Preload("User").Where(oi).Offset(pageNum).Limit(pageSize).Find(&ois).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	// 获取订单中的商品
	for i := 0; i < len(ois); i++ {
		var orderProducts []model.OrderProduct
		if err := global.DB.Preload("Product").Where("order_info_id = ?", ois[i].ID).Find(&orderProducts).
			Error; err != nil {
			return nil, err
		}

		ois[i].OrderProducts = orderProducts
	}

	return ois, nil
}

func GetOrderInfoTotal(oi model.OrderInfo) (int, error) {
	var count int

	if err := global.DB.Model(&model.OrderInfo{}).Where(oi).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
