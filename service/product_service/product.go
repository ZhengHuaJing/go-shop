package product_service

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
	"strings"
)

func ExistProductByName(p model.Product) (bool, error) {
	if err := global.DB.Where("name = ?", p.Name).First(&model.Product{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func ExistProductByID(p model.Product) (bool, error) {
	if err := global.DB.Where("id = ?", p.ID).First(&p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func AddProduct(p model.Product) (*model.Product, error) {
	if err := global.DB.Create(&p).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("product_sn = ?", p.ProductSN).First(&p).Error; err != nil {
		return nil, err
	}
	// 处理轮播图
	p.BannerUrls = strings.Split(p.BannerUrlStrs, ";")

	return &p, nil
}

func DeleteProduct(p model.Product) error {
	if err := global.DB.Delete(model.Product{}, "id = ?", p.ID).Error; err != nil {
		return err
	}

	return nil
}

func CleanSoftDeleteProduct() error {
	if err := global.DB.Unscoped().Where("deleted_at is not null").Delete(model.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func UpdateProduct(p model.Product) (*model.Product, error) {
	if err := global.DB.Model(model.Product{}).Where("id = ?", p.ID).Updates(p).Error; err != nil {
		return nil, err
	}

	if err := global.DB.Where("id = ?", p.ID).First(&p).Error; err != nil {
		return nil, err
	}
	// 处理轮播图
	p.BannerUrls = strings.Split(p.BannerUrlStrs, ";")

	return &p, nil
}

func GetProduct(p model.Product) (*model.Product, error) {
	if err := global.DB.Preload("Category").Where("id = ?", p.ID).First(&p).Error; err != nil {
		return nil, err
	}
	// 处理轮播图
	p.BannerUrls = strings.Split(p.BannerUrlStrs, ";")

	return &p, nil
}

func GetProducts(p model.Product, pageNum, pageSize int) ([]model.Product, error) {
	var ps []model.Product

	err := global.DB.Preload("Category").Where(p).Offset(pageNum).Limit(pageSize).Find(&ps).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	// 处理轮播图
	for _, p := range ps {
		p.BannerUrls = strings.Split(p.BannerUrlStrs, ";")
	}

	return ps, nil
}

func GetProductTotal(p model.Product) (int, error) {
	var count int

	if err := global.DB.Model(&model.Product{}).Where(p).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}
