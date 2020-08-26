package user_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistAddressByID(a model.Address) (bool, error) {
	if err := global.DB.Where("id = ?", a.ID).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddAddress(a model.Address) (*model.Address, error) {
	if err := global.DB.Create(&a).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("detail_address = ?", a.DetailAddress).First(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func DeleteAddress(a model.Address) error {
	if err := global.DB.Delete(model.Address{}, "id = ?", a.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteAddress() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.Address{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateAddress(a model.Address) (*model.Address, error) {
	if err := global.DB.Model(model.Address{}).Where("id = ?", a.ID).Updates(a).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", a.ID).First(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func GetAddress(a model.Address) (*model.Address, error) {
	if err := global.DB.Where("id = ?", a.ID).First(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func GetAddresses(a model.Address, pageNum, pageSize int) ([]model.Address, error) {
	var as []model.Address

	err := global.DB.Where(a).Offset(pageNum).Limit(pageSize).Find(&as).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return as, nil
}

func GetAddressTotal(a model.Address) (int, error) {
	var count int

	if err := global.DB.Model(&model.Address{}).Where(a).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
