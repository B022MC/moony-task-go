package service

import (
	"fmt"
	"github.com/spf13/cast"
	"math"
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/common/service"
	"moony-task-go/core/common"
	"moony-task-go/core/config"
	"moony-task-go/utils"

	"time"
)

type Base struct {
}

func BaseInstance() *Base {
	return &Base{}
}

// Launch 初始化用户信息
func (b *Base) Launch(session *common.Session, req model.UserConfigReq) model.UserConfigRsp {
	if session.UserHeader.DeviceId == "" {
		session.UserHeader.DeviceId = fmt.Sprintf("%d", time.Now().UnixNano())
	}

	var launch model.UserConfigRsp
	if session.UserToken == nil { //新用户
		var matchRsp *model.AdBackRsp
		matchRsp = service.AdInstance().Match(session, req, false)

		user := new(model.User)
		user.Status = model.UserStatusEnable
		user.AppId = session.UserHeader.AppId

		user.CreateTime = time.Now().Unix()
		user.ActiveTime = time.Now().Unix()
		user.HashId = utils.Hash64(session.UserHeader.DeviceId)
		hashNum := int(uint64(user.HashId/uint64(math.Pow(100, float64(1-1)))) % 100)
		user.HashNum = hashNum

		if matchRsp != nil && matchRsp.Data.AdId != 0 {

		}

		if err := dao.UserInstance().Create(user); err != nil {
			panic(config.ErrDb.New().Append(err))
		}
		launch.Token = common.CreateToken(user)
		launch.UserId = cast.ToString(user.Id)
		launch.HashId = cast.ToString(user.HashId)
		launch.Guest = user.IsGuest()
		//b.UserDevice(session.UserHeader.AppId, user.Id, session.UserHeader.RemoteIp, req)
	} else { //老用户
		user, err := dao.UserInstance().Get(session.UserToken.UserId)
		if err != nil {
			panic(config.ErrDb.New().Append(err))
		}
		if user == nil {
			panic(config.ErrParam.New().Append("user not exist"))
		}
		user.ActiveTime = time.Now().Unix()
		if err := dao.UserInstance().Update(user); err != nil {
			panic(config.ErrDb.New().Append(err))
		}
		launch.UserId = cast.ToString(session.UserToken.UserId)
		launch.HashId = cast.ToString(session.UserToken.HashId)
		launch.Guest = user.IsGuest()
	}

	launch.Config = session.Experiment.GetClient()

	return launch
}

// UserDevice 用户设备信息记录
//func (b *Base) UserDevice(appId int64, userId int64, ip string, req model.UserConfigReq) {
//	device, err := dao.UserDeviceInstance().GetByUserId(userId)
//	if err != nil {
//		log.Errorf("UserDevice.dao.UserDeviceInstance().GetByUserId err=[%s]", err.Error())
//		return
//	}
//	if device == nil {
//		device = new(model.UserDevice)
//	}
//	device.AppId = appId
//	device.UserId = userId
//	device.Oaid = req.Oaid
//	device.Idfa = req.Idfa
//	device.Imei = req.Imei
//	device.Ua = req.Ua
//	device.Ip = ip
//
//	if device.Id > 0 {
//		if err := dao.UserDeviceInstance().Update(device); err != nil {
//			log.Errorf("UserDevice.dao.UserDeviceInstance().Update err=[%s]", err.Error())
//		}
//	} else {
//		if err := dao.UserDeviceInstance().Create(device); err != nil {
//			log.Errorf("UserDevice.dao.UserDeviceInstance().Create err=[%s]", err.Error())
//		}
//	}
//}
