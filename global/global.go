package global

import (
	"github.com/casbin/casbin"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/smartwalle/alipay/v3"
	"github.com/spf13/viper"
	"github.com/zhenghuajing/fresh_shop/config"
	"go.uber.org/zap"
)

var (
	DB           *gorm.DB
	Enforcer     *casbin.Enforcer
	RedisPool    *redis.Pool
	Log          *zap.SugaredLogger
	Config       *config.Config
	Viper        *viper.Viper
	AlipayClient *alipay.Client
)
