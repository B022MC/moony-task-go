package main

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	log "github.com/sirupsen/logrus"
	"moony-task-go/common/service"
	"moony-task-go/core/config"
	"moony-task-go/core/global"
	"moony-task-go/routers"
	"moony-task-go/utils"
	"os"
	"time"
)

func main() {
	//初始化配置文件
	config.InitConfig()
	// 初始化客户端配置
	config.InitClientConfig()
	// 初始化服务端配置
	config.InitServerConfig()
	// 配置文件获取
	cfg := config.GetConfig()
	// 初始化日志
	initLog(cfg)

	// 数据库初始化
	if err := global.InitMysql(); err != nil {
		panic(err)
	}
	// 初始化redis链接
	//if err := global.InitRedis(); err != nil {
	//	panic(err)
	//}

	engine := gin.New()

	//初始化路由
	routers.InitRouter(engine)

	service.CrontabInstance().InitCrontab() //定时任务

	// 运行程序
	if err := engine.Run(cfg.Server.Address); err != nil {
		panic(config.ErrInternal.New().Append(err))
	}
}

// 日志格式化
func initLog(cfg *config.Config) {
	logfile := "log/server.log"
	writer, err := rotatelogs.New(
		logfile+".%Y%m%d",
		rotatelogs.WithLinkName(logfile),
		rotatelogs.WithMaxAge(time.Duration(86400)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(86400*7)*time.Second),
	)
	if err != nil {
		panic(err)
	}

	pathMap := lfshook.WriterMap{
		log.TraceLevel: writer,
		log.DebugLevel: writer,
		log.InfoLevel:  writer,
		log.WarnLevel:  writer,
		log.ErrorLevel: writer,
		log.FatalLevel: writer,
		log.PanicLevel: writer,
	}

	log.AddHook(lfshook.NewHook(pathMap, new(utils.LogFile)))

	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	log.SetFormatter(new(utils.LogFile))
	log.SetLevel(log.Level(cfg.Server.LogLevel))
}
