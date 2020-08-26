package model

import "github.com/jinzhu/gorm"

// 商品
type Product struct {
	gorm.Model

	CategoryID int      `json:"category_id" gorm:"comment:'分类ID'"`
	Category   Category `json:"category"`

	BannerUrls    []string `json:"banner_urls" gorm:"-"`
	BannerUrlStrs string   `json:"-" gorm:"comment:'商品轮播图，每个url以分号分割'"`
	FrontImageUrl string   `json:"front_image_url" gorm:"comment:'商品封面图片'"`
	ProductSN     string   `json:"product_sn" gorm:"comment:'商品唯一货号'"`
	Name          string   `json:"name" gorm:"comment:'商品名'"`
	ClickNum      int      `json:"click_num" gorm:"comment:'点击数'"`
	SoldNum       int      `json:"sold_num" gorm:"comment:'商品销量'"`
	CollectNum    int      `json:"collect_num" gorm:"comment:'收藏数'"`
	LeftNum       int      `json:"left_num" gorm:"comment:'商品库存'"`
	MarketPrice   float64  `json:"market_price" gorm:"comment:'市场价格'"`
	ShopPrice     float64  `json:"shop_price" gorm:"comment:'本店价格'"`
	Brief         string   `json:"brief" gorm:"comment:'商品简短描述'"`
	Desc          string   `json:"desc" gorm:"comment:'商品详细描述'"`
	IsFreight     int      `json:"is_freight" gorm:"default:0;comment:'是否承担运费'"`
	IsNew         int      `json:"is_new" gorm:"default:0;comment:'是否新品'"`
	IsHot         int      `json:"is_hot" gorm:"default:0;comment:'是否热销'"`
}
