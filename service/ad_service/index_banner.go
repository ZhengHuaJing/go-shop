package ad_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistIndexBannerByProductID(ib model.IndexBanner) (bool, error) {
	if err := global.DB.Where("product_id = ?", ib.ProductID).First(&model.IndexBanner{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistIndexBannerByID(ib model.IndexBanner) (bool, error) {
	if err := global.DB.Where("id = ?", ib.ID).First(&ib).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddIndexBanner(ib model.IndexBanner) (*model.IndexBanner, error) {
	if err := global.DB.Create(&ib).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("product_id = ?", ib.ProductID).First(&ib).Error; err != nil {
		return nil, err
	}

	return &ib, nil
}

func DeleteIndexBanner(ib model.IndexBanner) error {
	if err := global.DB.Delete(model.IndexBanner{}, "id = ?", ib.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteIndexBanner() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.IndexBanner{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateIndexBanner(ib model.IndexBanner) (*model.IndexBanner, error) {
	if err := global.DB.Model(model.IndexBanner{}).Where("id = ?", ib.ID).Updates(ib).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", ib.ID).First(&ib).Error; err != nil {
		return nil, err
	}

	return &ib, nil
}

func GetIndexBanner(ib model.IndexBanner) (*model.IndexBanner, error) {
	if err := global.DB.Where("id = ?", ib.ID).First(&ib).Error; err != nil {
		return nil, err
	}

	return &ib, nil
}

func GetIndexBanners(ib model.IndexBanner, pageNum, pageSize int) ([]model.IndexBanner, error) {
	var ibs []model.IndexBanner

	err := global.DB.Where(ib).Offset(pageNum).Limit(pageSize).Find(&ibs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ibs, nil
}

func GetIndexBannerTotal(ib model.IndexBanner) (int, error) {
	var count int

	if err := global.DB.Model(&model.IndexBanner{}).Where(ib).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
