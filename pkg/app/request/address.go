package request

type AddressForm struct {
	ID            int    `json:"-" form:"-"`
	UserID        int    `json:"user_id" form:"user_id" valid:"Min(1)"`
	Province      string `json:"province" form:"province" valid:"Required"`
	City          string `json:"city" form:"city" valid:"Required"`
	District      string `json:"district" form:"district" valid:"Required"`
	DetailAddress string `json:"detail_address" form:"detail_address" valid:"Required"`
	SignerName    string `json:"signer_name" form:"signer_name" valid:"Required"`
	SignerMobile  string `json:"signer_mobile" form:"signer_mobile" valid:"Mobile"`
}
