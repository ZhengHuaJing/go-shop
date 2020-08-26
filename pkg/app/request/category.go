package request

type CategoryForm struct {
	ID               int    `json:"-" form:"-"`
	ParentCategoryID int    `json:"parent_category_id" form:"parent_category_id"`
	Name             string `json:"name" form:"name" valid:"MinSize(1)"`
	Code             string `json:"code" form:"code" valid:"MinSize(1)"`
	Desc             string `json:"desc" form:"desc"`
	CategoryType     int    `json:"category_type" form:"category_type"  valid:"Min(1)"`
	IsTab            int    `json:"is_tab" form:"is_tab" valid:"Match(/^[0|1]{1}$/)"`
}
