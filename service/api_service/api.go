package api_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistApiByID(a model.Api) (bool, error) {
	if err := global.DB.Where("id = ?", a.ID).First(&a).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddApi(a model.Api) (*model.Api, error) {
	if err := global.DB.Create(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func GetApi(a model.Api) (*model.Api, error) {
	if err := global.DB.Where("id = ?", a.ID).First(&a).Error; err != nil {
		return nil, err
	}

	return &a, nil
}

func GetApis(a model.Api, pageNum, pageSize int) ([]model.Api, error) {
	var apis []model.Api

	err := global.DB.Where(a).Offset(pageNum).Limit(pageSize).Find(&apis).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return apis, nil
}

func GetApiTotal(a model.Api) (int, error) {
	var count int

	if err := global.DB.Model(&model.Api{}).Where(a).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
