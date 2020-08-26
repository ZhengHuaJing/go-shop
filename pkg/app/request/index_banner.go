package request

type IndexBannerForm struct {
	ID        int    `json:"-" form:"-"`
	ProductID int    `json:"product_id" form:"product_id" valid:"Min(1)"`
	Index     int    `json:"index" form:"index" valid:"Min(1)"`
	ImageUrl  string `json:"image_url" form:"image_url" valid:"MinSize(1)"`
}
