package request

type LeavingMessageForm struct {
	ID          int    `json:"-" form:"-"`
	UserID      int    `json:"user_id" form:"user_id" valid:"Min(1)"`
	MessageType int    `json:"message_type" form:"message_type" valid:"Min(1)"`
	Title       string `json:"title" form:"title" valid:"Required"`
	Content     string `json:"content" form:"content" valid:"Required"`
}
