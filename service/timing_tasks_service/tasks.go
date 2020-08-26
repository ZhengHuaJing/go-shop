package timing_tasks_service

import (
	"github.com/robfig/cron"
	"github.com/zhenghuajing/fresh_shop/global"
	"github.com/zhenghuajing/fresh_shop/service/user_service"
)

// 定时任务
func StartTasks() {
	c := cron.New()

	c.AddFunc("0 0 0 * * *", CleanSoftDeleteUser)

	c.Start()
}

func CleanSoftDeleteUser() {
	global.Log.Info("定时任务: 清空已删除的用户")
	user_service.CleanSoftDeleteUser()
}
