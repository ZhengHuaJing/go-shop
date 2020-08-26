package initialize

import (
	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
	"github.com/google/uuid"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/model"
	"github.com/zhenghuajing/fresh_shop/pkg/util"
	"github.com/zhenghuajing/fresh_shop/service/user_service"
	"log"
)

func Casbin() {
	// 权限管理
	casbinCfg := global.Config.Casbin
	adapter := gormadapter.NewAdapterByDB(global.DB)
	global.Enforcer = casbin.NewEnforcer(casbinCfg.CasbinConfPath, adapter)
	global.Enforcer.EnableLog(true)
	if err := global.Enforcer.LoadPolicy(); err != nil {
		log.Fatalf("加载权限规则失败: %s", err)
	}

	newUUID, _ := uuid.NewUUID()
	userModel := model.User{
		UserName: casbinCfg.DefaultAdminUserName,
		Password: util.MD5(casbinCfg.DefaultAdminPassword + global.Config.MD5.Salt),
		UUID:     newUUID,
	}
	ok, err := user_service.ExistUserByUserName(userModel)
	if err != nil {
		log.Fatalf("验证管理员账户存在失败: %s", err)
	}
	if !ok {
		_, err := user_service.AddUser(userModel)
		if err != nil {
			log.Fatalf("创建管理员账户失败: %s", err)
		}
	}

	// 权限表初始化两个角色，超级管理员和普通用户
	global.Enforcer.AddRoleForUser(casbinCfg.DefaultAdminUserName, casbinCfg.RoleAdmin)
}
