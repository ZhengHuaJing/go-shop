package cache_service

import (
	"github.com/zhenghuajing/fresh_shop/pkg/e"
	"strconv"
)

type User struct {
	ID int
}

func (u *User) GetUserKey() string {
	return e.CACHE_USER + "_" + strconv.Itoa(u.ID)
}
