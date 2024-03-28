package service

import (
	"github.com/go-co-op/gocron"
	"time"
)

type Crontab struct {
}

func CrontabInstance() *Crontab {
	return &Crontab{}
}

// InitCrontab 初始化定时任务
func (c *Crontab) InitCrontab() {
	timezone, _ := time.LoadLocation("Asia/Shanghai")
	cron := gocron.NewScheduler(timezone)

	cron.StartAsync()
}
