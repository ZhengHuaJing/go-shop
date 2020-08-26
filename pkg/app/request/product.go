package request

type ProductForm struct {
	ID            int      `json:"-" form:"-"`
	BannerUrls    []string `json:"banner_urls" form:"banner_urls" valid:"Required"`
	CategoryID    int      `json:"category_id" form:"category_id" valid:"Min(1)"`
	Name          string   `json:"name" form:"name" valid:"MinSize(1)"`
	LeftNum       int      `json:"left_num" form:"left_num" valid:"Min(1)"`
	MarketPrice   float64  `json:"market_price" form:"market_price" valid:"Required"`
	ShopPrice     float64  `json:"shop_price" form:"shop_price" valid:"Required"`
	Brief         string   `json:"brief" form:"brief" valid:"MinSize(5)"`
	Desc          string   `json:"desc" form:"desc" valid:"MinSize(20)"`
	FrontImageUrl string   `json:"front_image_url" form:"front_image_url" valid:"Required"`
	IsFreight     int      `json:"is_freight" form:"is_freight" valid:"Match(/^[0|1]{1}$/)"`
	IsNew         int      `json:"is_new" form:"is_new" valid:"Match(/^[0|1]{1}$/)"`
	IsHot         int      `json:"is_hot" form:"is_hot" valid:"Match(/^[0|1]{1}$/)"`
}
