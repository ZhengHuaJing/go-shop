package request

type AdForm struct {
	ID         int `json:"-" form:"-"`
	CategoryID int `json:"category_id" form:"category_id" valid:"Min(1)"`
	ProductID  int `json:"product_id" form:"product_id" valid:"Min(1)"`
}
