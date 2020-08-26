package request

type HotSearchForm struct {
	ID      int    `json:"-" form:"-"`
	Keyword string `json:"keyword" form:"keyword" valid:"Required"`
	Index   int    `json:"index" form:"index" valid:"Min(1)"`
}
