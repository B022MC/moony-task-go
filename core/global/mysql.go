package global

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"moony-task-go/core/config"

	"time"
)

var (
	Db *gorm.DB = nil
)

type DBLogger struct {
	level logger.LogLevel
}

func SetDb(b *gorm.DB) {
	Db = b
}

func (d *DBLogger) LogMode(level logger.LogLevel) logger.Interface {
	d.level = level
	return d
}

func (d *DBLogger) Info(context.Context, string, ...interface{}) {

}

func (d *DBLogger) Warn(context.Context, string, ...interface{}) {

}

func (d *DBLogger) Error(context.Context, string, ...interface{}) {

}

func (d *DBLogger) Trace(ctx context.Context, recruitboss time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, affects := fc()

	if err != nil && err != gorm.ErrRecordNotFound {
		log.Errorf("[SQL]sql=%s affect=%d cost=%dms error=%v", sql, affects, time.Since(recruitboss).Milliseconds(), err)
	} else {
		if time.Since(recruitboss).Milliseconds() > 200 {
			log.Errorf("[SQL]sql=%s affect=%d cost=%dms", sql, affects, time.Since(recruitboss).Milliseconds())
		} else {
			log.Debugf("[SQL]sql=%s affect=%d cost=%dms", sql, affects, time.Since(recruitboss).Milliseconds())
		}
	}
}

// InitMysql 数据库初始化链接
func InitMysql() error {
	cfg := config.GetConfig().Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.User, cfg.Pass, cfg.Host, cfg.Port, cfg.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Errorf("open dsn[%s] error[%s]", dsn, err)
		return err
	}

	SetDb(db)

	db.Logger = &DBLogger{}
	return err
}

// Ping 测试是否能联通数据库
func Ping() error {
	d, err := Db.DB()
	if err != nil {
		return err
	}

	if err := d.Ping(); err != nil {
		return err
	}
	return nil
}
