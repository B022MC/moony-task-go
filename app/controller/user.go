package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"io"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
	"moony-task-go/core/common"
	"moony-task-go/core/config"
	"moony-task-go/utils"
	"net/http"
)

type User struct {
}

func UserInstance() *User {
	return &User{}
}

/**
* @api {post} /api/user/login 用户登录
* @apiGroup User
* @apiName user-login
* @apiDescription 用户登录
* @apiUse baseHeader
* @apiParam (业务参数（body）) {string} loginType 登录类型，phone：手机 weixin：微信 account：帐号密码，ios:ios授权登陆
* @apiParam (业务参数（body）) {Object} [weixin] 微信登录确认码
* @apiParam (业务参数（body）) {string} [weixin[code]] 微信登录确认码
* @apiParam (业务参数（body）) {string} [ios[nickname]]  昵称
* @apiParamExample {json} 请求示例
{
	"loginType": "weixin",
	"phone":{
		"phone": "13527408395",
		"code": "314527"
	},
	"weixin": {
    	"code": "123456"
	},
    "account": {
        "username": "test",
        "password": "password"
    },
	"ios":{
		"key":"HJl1kkaslkd",
		"email":"11222@qq.com",
		"nickname":"你好啊"
	},
	"univerify":{
		"openid":"openid",
		"accessToken":"asdasdasdadasdasd"
	}
}
*/

func (u *User) Login(ctx *gin.Context) {
	var request model.LoginUserReq
	if err := ctx.ShouldBind(&request); err != nil {
		panic(config.ErrParam.New().Append(err))
	}
	log.Debugf("loging request: %s", utils.EncodeJSON(request))

	session := ctx.Keys[common.ContextSession].(*common.Session)

	result, err := service.UserInstance().Login(session, request)
	if err != nil {
		panic(config.ErrDb.New().Append(err))
	}
	ctx.JSON(http.StatusOK, common.NewRsp(result))
}

// Logout 退出登录
//func (u *User) Logout(ctx *gin.Context) {
//	session := ctx.Keys[common.ContextSession].(*common.Session)
//
//	result, err := service.UserInstance().Logout(session)
//	if err != nil {
//		panic(config.ErrDb.New().Append(err))
//	}
//	ctx.JSON(http.StatusOK, common.NewRsp(result))
//}

// Cancel 用户注销
func (u *User) Cancel(ctx *gin.Context) {
	if ctx.Request.Method == http.MethodGet {
		type UserCancelReq struct {
			Signature string `form:"signature"`
			Timestamp int64  `form:"timestamp"`
			Nonce     string `form:"nonce"`
			Echostr   string `form:"echostr"`
		}
		var req *UserCancelReq
		if err := ctx.Bind(&req); err != nil {
			panic(config.ErrParam.New().Append(err))
		}
		log.Debugf("usercancelget param=[%s]", utils.EncodeJSON(req))
		ctx.String(http.StatusOK, cast.ToString(req.Echostr))
		return
	}

	bodyBytes, _ := io.ReadAll(ctx.Request.Body)

	log.Debugf("usercancel param =[%s]", string(bodyBytes))
	ctx.String(http.StatusOK, "success")
	return
}

// GetUserInfo 用户信息
func (u *User) GetUserInfo(ctx *gin.Context) {
	session := ctx.Keys[common.ContextSession].(*common.Session)
	result, err := service.UserInstance().GetUserInfo(session)
	if err != nil {
		panic(config.ErrDb.New().Append(err))
	}
	ctx.JSON(http.StatusOK, common.NewRsp(result))
}

// SaveUserInfo 修改用户信息
func (u *User) SaveUserInfo(ctx *gin.Context) {
	request := make(map[string]interface{})
	if err := ctx.BindJSON(&request); err != nil {
		panic(config.ErrParam.New().Append(err))
	}
	session := ctx.Keys[common.ContextSession].(*common.Session)
	err := service.UserInstance().SaveUserInfo(session, request)
	if err != nil {
		panic(config.ErrDb.New().Append(err))
	}
	ctx.JSON(http.StatusOK, common.NewRspOk())
}

/**
* @api {post} /api/user/event 用户埋点
* @apiGroup User
* @apiName user-event
* @apiDescription 用户上报事件相关信息，客户端所有事件（包括错误）均要上报
* @apiUse baseHeader
* @apiParam (业务参数（body）) {string} source 上报来源，client:客户端（包括ios/android/h5/小程序）server:服务端 spider:爬虫
* @apiParam (业务参数（body）) {string} type 上报类型，click:点击事件，error：错误信息
* @apiParam (业务参数（body）) {string} key 事件名称
* @apiParam (业务参数（body）) {string} [value] 事件的值，如点击“返回”可以上报这个页面的停留时间
* @apiParam (业务参数（body）) {string} [extra] 其他扩展信息
* @apiSuccessExample {json} 返回成功
* HTTP/1.1 200 OK
 {
   "code": 0,
   "message": "ok"
 }
**/
