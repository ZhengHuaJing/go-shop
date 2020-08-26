package product_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

func ExistCategoryByName(c model.Category) (bool, error) {
	if err := global.DB.Where("name = ?", c.Name).First(&model.Category{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistCategoryByID(c model.Category) (bool, error) {
	if err := global.DB.Where("id = ?", c.ID).First(&c).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddCategory(c model.Category) (*model.Category, error) {
	if err := global.DB.Create(&c).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("name = ?", c.Name).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func DeleteCategory(c model.Category) error {
	if err := global.DB.Delete(model.Category{}, "id = ?", c.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteCategory() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.Category{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateCategory(c model.Category) (*model.Category, error) {
	if err := global.DB.Model(model.Category{}).Where("id = ?", c.ID).Updates(c).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", c.ID).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func GetCategory(c model.Category) (*model.Category, error) {
	if err := global.DB.Where("id = ?", c.ID).First(&c).Error; err != nil {
		return nil, err
	}

	return &c, nil
}

func GetCategorys(c model.Category, pageNum, pageSize int) ([]model.Category, error) {
	var cs []model.Category

	err := global.DB.Where(c).Offset(pageNum).Limit(pageSize).Find(&cs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return cs, nil
}

func GetCategoryTotal(c model.Category) (int, error) {
	var count int

	if err := global.DB.Model(&model.Category{}).Where(c).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
