package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"moony-task-go/core/config"
	"moony-task-go/core/global"
)

type Request struct {
}

func RequestInstance() *Request {
	return &Request{}
}

// RequestLimit 请求限制中间件
func (r *Request) RequestLimit(ctx *gin.Context) {
	cfg := config.GetConfig().Server
	RequestLimitKey := fmt.Sprintf("%d_request_limit", cfg.AppId)
	result := global.GetRedis(RequestLimitKey)
	requestNum := cast.ToInt64(result)
	if requestNum >= cfg.RequestLimitNum {
		panic(config.ErrRequestLimit.New().Append("请求超出限制"))
	} else {
		requestNum = requestNum + 1
		if err := global.SetRedis(RequestLimitKey, cast.ToString(requestNum), 1); err != nil {
			panic(config.ErrInternal.New().Append(err))
		}
	}
	ctx.Next()
}
