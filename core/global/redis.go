package global

import (
	"errors"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"moony-task-go/core/config"
	"time"
)

var RedisClient *redis.Client = nil

// InitRedis 初始化redis
func InitRedis() error {
	cfg := config.GetConfig().Redis
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Password,
		DB:       cfg.Db,
	})
	if client == nil {
		return errors.New("redis连接失败")
	}
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}

	RedisClient = client

	return nil
}

// GetRedis 获取数据
func GetRedis(key string) string {
	result, err := RedisClient.Get(key).Result()
	if errors.Is(err, redis.Nil) {
		return ""
	}
	if err != nil {
		log.Errorf("redis get  err=[%s]", err.Error())
		return ""
	}

	return result
}

// SetRedis 设置数据
func SetRedis(key string, data string, expireTime int) error {
	_, err := RedisClient.Set(key, data, time.Second*time.Duration(expireTime)).Result()
	if err != nil {
		log.Errorf("redis set cache err=[%s]", err.Error())
		return err
	}

	return nil
}

// DeleteRedis 删除数据
func DeleteRedis(key string) error {
	_, err := RedisClient.Del(key).Result()
	if err != nil {
		log.Errorf("redis delete err=[%s]", err.Error())
		return err
	}
	return nil
}
