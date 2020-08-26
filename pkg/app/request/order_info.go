package request

type OrderProduct struct {
	ProductID int `json:"product_id" form:"product_id" valid:"Min(1)"`
	Num       int `json:"num" form:"num" valid:"Min(1)"`
}

type OrderInfoForm struct {
	ID            int            `json:"-" form:"-"`
	UserID        int            `json:"user_id" form:"user_id" valid:"Min(1)"`
	PostScript    string         `json:"post_script" form:"post_script"`                        // 订单留言
	OrderProducts []OrderProduct `json:"order_products" form:"order_products" valid:"Required"` // 订单中的商品
	ClientType    int            `json:"client_type" form:"client_type" valid:"Min(1)"`         // 1: Web, 2: Web App
	ReturnUrl     string         `json:"return_url" form:"return_url" valid:"Required"`         // 支付成功后跳转页面
}
