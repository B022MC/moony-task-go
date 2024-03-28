package controller

import (
	"github.com/gin-gonic/gin"
	"moony-task-go/app/service"
	"moony-task-go/common/model"
	"moony-task-go/core/common"
	"moony-task-go/core/config"
	"moony-task-go/core/global"
	"net/http"
)

/**
 * @apiDefine Base 基础配置
 *
 */

/**
* @apiDefine baseHeader
* @apiHeader (基础参数（header）) {string} x-app-id 应用ID，每个应用独立分配的id
* @apiHeader (基础参数（header）) {string} x-token 登录的token，不需要登录的接口可以为空
* @apiHeader (基础参数（header）) {String} x-version 客户端版本号，如：1.3.123
* @apiHeader (基础参数（header）) {String} x-platform 客户端平台，如：ios、android、h5、admin
* @apiHeader (基础参数（header）) {String} x-device-id 客户端设备唯一ID，如：IMEI、android_id
* @apiHeader (基础参数（header）) {String} x-channel 用户来源渠道
* @apiHeader (基础参数（header）) {String} x-mobile-brand 手机品牌
* @apiHeader (基础参数（header）) {String} x-mobile-model 手机型号
* @apiHeader (基础参数（header）) {String} x-user-agent UA
 */

type Base struct {
}

func BaseInstance() *Base {
	return &Base{}
}

/**
* @api {get} /health 健康检测
* @apiGroup Base
* @apiName Base-health
* @apiDescription 健康检测
@apiSuccessExample {json} 返回成功
HTTP/1.1 200 OK
{
    "code": 0,
    "message": "OK",
}
*/

func (b *Base) Health(ctx *gin.Context) {
	if err := global.Ping(); err != nil {
		ctx.String(http.StatusOK, "%s", err.Error())
		return
	}
	ctx.String(http.StatusOK, "%s", "OK")
}

/**
* @api {get} /api/launch 配置获取
* @apiGroup Base
* @apiName Base-health
* @apiDescription 健康检测
@apiSuccessExample {json} 返回成功
HTTP/1.1 200 OK
{
    "code": 0,
    "message": "OK",
	"data": {
        "config": {
            "pay.vip.pre": false,
            "rec.count2": 100
        },
        "token":"YWJkZGVhZmZhCg==", // 新用户产的临时token，首次启动的新用户使用
        "userId": "6",
        "hashId": "7179595360673787904",
        "debug": {
        }
    }
}
*/

func (b *Base) Launch(ctx *gin.Context) {
	var req model.UserConfigReq
	if err := ctx.Bind(&req); err != nil {
		panic(config.ErrParam.New().Append(err))
	}
	session := ctx.Keys[common.ContextSession].(*common.Session)

	launch := service.BaseInstance().Launch(session, req)

	ctx.JSON(http.StatusOK, common.NewRsp(launch))
}
