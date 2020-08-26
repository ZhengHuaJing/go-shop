package user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistCollectByUserIDAndProductID(c model.Collect) (bool, error) {
	err := global.DB.Where("user_id = ? AND product_id = ?", c.UserID, c.ProductID).First(&model.Collect{}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistCollectByID(c model.Collect) (bool, error) {
	if err := global.DB.Where("id = ?", c.ID).First(&c).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddCollect(c model.Collect) (*model.Collect, error) {
	if err := global.DB.Create(&c).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("user_id = ? AND product_id = ?", c.UserID, c.ProductID).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func DeleteCollect(c model.Collect) error {
	if err := global.DB.Delete(model.Collect{}, "id = ?", c.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteCollect() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.Collect{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateCollect(c model.Collect) (*model.Collect, error) {
	if err := global.DB.Model(model.Collect{}).Where("id = ?", c.ID).Updates(c).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", c.ID).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func GetCollect(c model.Collect) (*model.Collect, error) {
	if err := global.DB.Where("id = ?", c.ID).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func GetCollects(c model.Collect, pageNum, pageSize int) ([]model.Collect, error) {
	var cs []model.Collect

	err := global.DB.Where(c).Offset(pageNum).Limit(pageSize).Find(&cs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return cs, nil
}

func GetCollectTotal(c model.Collect) (int, error) {
	var count int

	if err := global.DB.Model(&model.Collect{}).Where(c).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
