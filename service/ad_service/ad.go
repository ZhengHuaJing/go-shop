package ad_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistAdByID(ad model.Ad) (bool, error) {
	if err := global.DB.Where("id = ?", ad.ID).First(&ad).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddAd(ad model.Ad) (*model.Ad, error) {
	if err := global.DB.Create(&ad).Error; err != nil {
		return nil, err
	}

	err := global.DB.Where("category_id = ? AND product_id = ?", ad.CategoryID, ad.ProductID).First(&ad).Error
	if err != nil {
		return nil, err
	}

	return &ad, nil
}

func DeleteAd(ad model.Ad) error {
	if err := global.DB.Delete(model.Ad{}, "id = ?", ad.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteAd() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.Ad{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateAd(ad model.Ad) (*model.Ad, error) {
	if err := global.DB.Model(model.Ad{}).Where("id = ?", ad.ID).Updates(ad).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", ad.ID).First(&ad).Error; err != nil {
		return nil, err
	}

	return &ad, nil
}

func GetAd(ad model.Ad) (*model.Ad, error) {
	if err := global.DB.Where("id = ?", ad.ID).First(&ad).Error; err != nil {
		return nil, err
	}

	return &ad, nil
}

func GetAds(ad model.Ad, pageNum, pageSize int) ([]model.Ad, error) {
	var ads []model.Ad

	err := global.DB.Where(ad).Offset(pageNum).Limit(pageSize).Find(&ads).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return ads, nil
}

func GetAdTotal(ad model.Ad) (int, error) {
	var count int

	if err := global.DB.Model(&model.Ad{}).Where(ad).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
