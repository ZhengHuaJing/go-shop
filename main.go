package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/initialize"
	"github.com/zhenghuajing/fresh_shop/service/timing_tasks_service"
	"net/http"
)

func init() {
	initialize.Config()
	initialize.Log()
	initialize.Mysql()
	initialize.Redis()
	initialize.DBTables()
	initialize.Casbin()
	initialize.Alipay()
}

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	// 定时任务
	timing_tasks_service.StartTasks()

	serverCfg := global.Config.Server
	gin.SetMode(serverCfg.RunMode)
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", serverCfg.HttpPort),
		Handler:        initialize.Router(),
		ReadTimeout:    serverCfg.ReadTimeout,
		WriteTimeout:   serverCfg.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

	// 程序结束前关闭数据库链接
	defer global.DB.Close()
}
