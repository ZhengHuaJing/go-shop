package initialize

import (
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
)

// 注册数据库表专用
func DBTables() {
	// 因为同步最新的api接口到数据库，所以要删除旧表
	if global.DB.HasTable("apis") {
		global.DB.DropTable("apis")
	}

	global.DB.AutoMigrate(
		model.User{},
		model.Api{},
		model.Ad{},
		model.Address{},
		model.Brand{},
		model.Category{},
		model.Collect{},
		model.HotSearch{},
		model.IndexBanner{},
		model.LeavingMessage{},
		model.Product{},
		model.ShoppingCart{},
		model.VerifyCode{},
		model.OrderInfo{},
		model.OrderProduct{},
	)
}
