package request

type BrandForm struct {
	ID         int    `json:"-" form:"-"`
	CategoryID int    `json:"category_id" form:"category_id" valid:"Min(1)"`
	Name       string `json:"name" form:"name" valid:"MinSize(1)"`
	Desc       string `json:"desc" form:"desc" valid:"MinSize(1)"`
	ImageUrl   string `json:"image_url" form:"image_url" valid:"MinSize(1)"`
}
