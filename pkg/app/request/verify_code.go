package request

type VerifyCodeForm struct {
	Email string `json:"email" form:"email" valid:"Email"`
}
