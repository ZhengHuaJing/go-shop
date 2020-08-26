package hot_search_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistHotSearchByKeyword(hs model.HotSearch) (bool, error) {
	if err := global.DB.Where("keyword = ?", hs.Keyword).First(&model.HotSearch{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistHotSearchByID(hs model.HotSearch) (bool, error) {
	if err := global.DB.Where("id = ?", hs.ID).First(&hs).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddHotSearch(hs model.HotSearch) (*model.HotSearch, error) {
	if err := global.DB.Create(&hs).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("keyword = ?", hs.Keyword).First(&hs).Error; err != nil {
		return nil, err
	}

	return &hs, nil
}

func DeleteHotSearch(hs model.HotSearch) error {
	if err := global.DB.Delete(model.HotSearch{}, "id = ?", hs.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteHotSearch() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.HotSearch{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateHotSearch(hs model.HotSearch) (*model.HotSearch, error) {
	if err := global.DB.Model(model.HotSearch{}).Where("id = ?", hs.ID).Updates(hs).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", hs.ID).First(&hs).Error; err != nil {
		return nil, err
	}

	return &hs, nil
}

func GetHotSearch(hs model.HotSearch) (*model.HotSearch, error) {
	if err := global.DB.Where("id = ?", hs.ID).First(&hs).Error; err != nil {
		return nil, err
	}

	return &hs, nil
}

func GetHotSearchs(hs model.HotSearch, pageNum, pageSize int) ([]model.HotSearch, error) {
	var hss []model.HotSearch

	err := global.DB.Where(hs).Offset(pageNum).Limit(pageSize).Find(&hss).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return hss, nil
}

func GetHotSearchTotal(hs model.HotSearch) (int, error) {
	var count int

	if err := global.DB.Model(&model.HotSearch{}).Where(hs).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
