package request

type ID struct {
	ID int `json:"id" form:"id" valid:"Min(1)"` // ID
}
