package initialize

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhenghuajing/fresh_shop/global"
	"log"
)

func Mysql() {
	var err error
	mysqlCfg := global.Config.Mysql
	args := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Name,
	)
	global.DB, err = gorm.Open(mysqlCfg.Type, args)
	if err != nil {
		log.Fatal("数据库连接失败: " + err.Error())
	}

	// 表前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return mysqlCfg.TablePrefix + defaultTableName
	}

	global.DB.DB().SetMaxIdleConns(10)
	global.DB.DB().SetMaxOpenConns(100)

	// 是否显示 SQL 语句
	global.DB.LogMode(false)
}
