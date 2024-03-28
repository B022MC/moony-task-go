package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogap/errors"
	log "github.com/sirupsen/logrus"
	"moony-task-go/common/dao"
	"moony-task-go/core/common"
	"moony-task-go/core/config"
	"moony-task-go/utils"
	"net/http"
	"runtime"
	"time"
)

type Api struct {
}

func ApiInstance() *Api {
	return &Api{}
}

// ApiBefore api前置监测
func (a *Api) ApiBefore(ctx *gin.Context) {
	// 获取头部参数
	var header common.UserHeader
	if err := ctx.BindHeader(&header); err != nil {
		panic(config.ErrParam.New().Append(err))
	}
	if ctx.Keys == nil {
		ctx.Keys = make(map[string]interface{})
	}
	header.AppId = config.GetConfig().Server.AppId

	session := new(common.Session)
	session.UserHeader = &header
	ctx.Keys[common.ContextSession] = session

	// 验证是否忽略token
	if !utils.StrInContains(ctx.Request.URL.Path, config.GetConfig().Server.TokenIgnore) {
		//if header.Token == "" {
		//	panic(config.ErrParam.New().Append("token不能为空"))
		//}
		if header.AppId == 0 {
			panic(config.ErrInternal.New().Append(fmt.Sprintf("IP:=[%s]", header.RemoteIp)))
		}
		if utils.StrInSlice(header.RemoteIp, []string{"127.0.0.1"}) {
			panic(config.ErrInternal.New().Append(fmt.Sprintf("IP:=[%s]", header.RemoteIp)))
		}
	}

	if header.Token != "" {
		session.UserToken = common.ParseToken(header.Token)
		user, err := dao.UserInstance().Get(session.GetUserId())
		if err != nil {
			panic(config.ErrDb.New().Append(err))
		}
		if user == nil {
			panic(config.ErrNoLogin.New().Append(err))
		}
		//tokenMd5 := utils.Md5(session.UserHeader.Token)
		//if tokenMd5 != user.TokenMd5 && !user.IsGuest() {
		//panic(config.ErrNoLogin.New().Append("未登陆，请先登陆"))
		//}
		session.User = user
	}

	session.Context = ctx
	//session.Experiment = new(common.Experiment)
	//if err := session.Experiment.LoadExperiment(session); err != nil {
	//	log.Errorf("load session error %s", err.Error())
	//}
}

// Recovery api中间件
func (a *Api) Recovery(c *gin.Context) {
	defer func() {
		if err := recover(); err != nil {
			var rsp common.BaseResponse
			if e, ok := err.(errors.ErrCode); ok {
				rsp.Code = int(e.Code())
				rsp.Message = e.Error()
				c.JSON(http.StatusOK, rsp)
				log.Errorf("[%s][%s][%s] %s", c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery, e.Error())
			} else {
				var buf [2 << 10]byte
				stack := string(buf[:runtime.Stack(buf[:], true)])
				log.Errorf("[%s][%s][%s] Internal ERROR:::%v stack:%s", c.Request.Method, c.Request.URL.Path, c.Request.URL.RawQuery, err, stack)
				rsp.Code = -1
				rsp.Message = fmt.Sprintf("%v", err)
				c.JSON(http.StatusOK, rsp)
			}
			c.Abort()
		}
	}()
	start := time.Now()
	c.Next()
	//a.Event(c) //接口行为记录

	session := c.Keys[common.ContextSession].(*common.Session)
	log.Infof("[%s][%s] cost[%s] header[%s] query[%s]",
		c.Request.Method, c.Request.URL.Path, time.Since(start), utils.EncodeJSON(session.UserHeader), c.Request.URL.RawQuery)
}

// 会员验证
//func (a *Api) VipFilter(c *gin.Context) {
//	// 详情只能够会员看
//	if strings.HasSuffix(c.Request.URL.Path, "detail") {
//		session := c.Keys[common.ContextSession].(*common.Session)
//		isPay := cast.ToBool(session.Experiment.Get(common.ClientPayEnableKey))
//		if isPay && session.User.VipExpireTime < time.Now().Unix() {
//			panic(config.ErrNoMember.New().Append("无权限查看"))
//		}
//	}
//}

// Event 接口行为记录
//func (a *Api) Event(c *gin.Context) {
//	session := c.Keys[common.ContextSession].(*common.Session)
//	// 请求埋点
//	if c.Request.URL.Path == "/api/user/event" {
//		return
//	}
//	_ = service.UserInstance().CreateServerUserEvent(session, c)
//}
