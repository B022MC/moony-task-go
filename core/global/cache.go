package global

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/os/gcache"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"time"
)

// GetCache 获取缓存
func GetCache(ctx *gin.Context) (string, error) {
	cache := new(gcache.Cache)
	cacheKey := fmt.Sprintf("%s_%s_%s", ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.URL.Query().Encode())
	result, err := cache.Get(cacheKey)
	if err != nil {
		log.Errorf("get cache err=[%s]", err.Error())
		return "", err
	}
	return cast.ToString(result), nil
}

// SetCache 设置缓存
func SetCache(ctx *gin.Context, value string, cacheTime int) error {
	cache := new(gcache.Cache)
	cacheKey := fmt.Sprintf("%s_%s_%s", ctx.Request.Method, ctx.Request.URL.Path, ctx.Request.URL.Query().Encode())
	err := cache.Set(cacheKey, value, time.Second*time.Duration(cacheTime))
	if err != nil {
		log.Errorf("set cache err=[%s]", err.Error())
		return err
	}

	return nil
}
