package request

type LoginUserForm struct {
	UserName string `json:"user_name" form:"user_name" valid:"Required"` // 用户名
	Password string `json:"password" form:"password" valid:"MinSize(8)"` // 密码
}

type AddUserForm struct {
	UserName   string `json:"user_name" form:"user_name" valid:"Required"` // 用户名
	Password   string `json:"password" form:"password" valid:"MinSize(8)"` // 密码
	VerifyCode string `json:"verify_code" form:"verify_code" valid:"Required"`
}

type UpdateUserForm struct {
	ID       int    `json:"-" form:"-" valid:"Min(1)"`                     // ID
	Password string `json:"password" form:"password"`                      // 密码
	RealName string `json:"real_name" form:"real_name" valid:"MinSize(2)"` // 联系人
	Mobile   string `json:"mobile" form:"mobile" valid:"Mobile"`           // 手机号
	Address  string `json:"address" form:"address" valid:"MinSize(6)"`     // 地址
	Gender   string `json:"gender" form:"Address" valid:"Length(1)"`       // 性别
}
