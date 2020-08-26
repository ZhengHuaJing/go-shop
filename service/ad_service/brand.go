package ad_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistBrandByName(brand model.Brand) (bool, error) {
	if err := global.DB.Where("name = ?", brand.Name).First(&model.Brand{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistBrandByID(brand model.Brand) (bool, error) {
	if err := global.DB.Where("id = ?", brand.ID).First(&brand).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddBrand(brand model.Brand) (*model.Brand, error) {
	if err := global.DB.Create(&brand).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("name = ?", brand.Name).First(&brand).Error; err != nil {
		return nil, err
	}

	return &brand, nil
}

func DeleteBrand(brand model.Brand) error {
	if err := global.DB.Delete(model.Brand{}, "id = ?", brand.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteBrand() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.Brand{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateBrand(brand model.Brand) (*model.Brand, error) {
	if err := global.DB.Model(model.Brand{}).Where("id = ?", brand.ID).Updates(brand).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", brand.ID).First(&brand).Error; err != nil {
		return nil, err
	}

	return &brand, nil
}

func GetBrand(brand model.Brand) (*model.Brand, error) {
	if err := global.DB.Where("id = ?", brand.ID).First(&brand).Error; err != nil {
		return nil, err
	}

	return &brand, nil
}

func GetBrands(brand model.Brand, pageNum, pageSize int) ([]model.Brand, error) {
	var brands []model.Brand

	err := global.DB.Where(brand).Offset(pageNum).Limit(pageSize).Find(&brands).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return brands, nil
}

func GetBrandTotal(brand model.Brand) (int, error) {
	var count int

	if err := global.DB.Model(&model.Brand{}).Where(brand).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
