package request

type ShoppingCartForm struct {
	ID        int `json:"-" form:"-"`
	UserID    int `json:"user_id" form:"user_id" valid:"Min(1)"`
	ProductID int `json:"product_id" form:"product_id" valid:"Min(1)"`
	Num       int `json:"num" form:"num" valid:"Min(1)"`
}
