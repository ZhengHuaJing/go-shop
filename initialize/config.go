package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/zhenghuajing/fresh_shop/global"
	"time"
)

const defaultConfigFile = "config/config.yaml"

func Config() {
	global.Viper = viper.New()
	global.Viper.SetConfigFile(defaultConfigFile)
	err := global.Viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}
	global.Viper.WatchConfig()

	global.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := global.Viper.Unmarshal(&global.Config); err != nil {
			fmt.Println(err)
		}
	})
	if err := global.Viper.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}

	// 初始化配置
	serverCfg := &global.Config.Server
	serverCfg.ReadTimeout = serverCfg.ReadTimeout * time.Second
	serverCfg.WriteTimeout = serverCfg.WriteTimeout * time.Second

	fileCfg := &global.Config.File
	fileCfg.ImageMaxSize = fileCfg.ImageMaxSize * 1024 * 1024

	redisCfg := &global.Config.Redis
	redisCfg.IdleTimeout = redisCfg.IdleTimeout * time.Second
}
