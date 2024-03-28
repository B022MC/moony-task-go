package service

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/common"
	"moony-task-go/core/config"
	"moony-task-go/utils"
)

var (
	apiPay      = "/api/client/pay"
	apiMatch    = "/api/client/match"
	apiActive   = "/api/client/active"
	apiRegister = "/api/client/register"
)

type Ad struct {
}

func AdInstance() *Ad {
	return new(Ad)
}

// Match 匹配
func (a *Ad) Match(session *common.Session, req model.UserConfigReq, isActive bool) *model.AdBackRsp {
	match := model.ClientInfo{
		AppId:   session.UserHeader.AppId,
		Channel: session.UserHeader.Channel,
		Version: session.UserHeader.Version,
		Os:      session.UserHeader.Platform,
		Model:   session.UserHeader.MobileModel,
		Ip:      session.UserHeader.RemoteIp,
		Ua:      req.Ua,
		Idfa:    req.Idfa,
		Oaid:    req.Oaid,
		Imei:    req.Imei,
	}
	if match.Ip == "" {
		match.Ip = session.Context.GetHeader("x-public-ip")
	}
	if match.Ip == "" {
		match.Ip = session.Context.GetHeader("real-address-ip")
	}
	match.Active = isActive
	match.Repair()

	log.Debugf("MatchData param %s", utils.EncodeJSON(match))

	cfg := config.GetConfig().Ad
	reqUrl := cfg.Url + apiMatch
	postJson, err := utils.HttpPostJson(reqUrl, utils.EncodeJSON(match), cfg.SignParam)
	if err != nil {
		log.Errorf("adMatch postParam=[%s] err=[%s]", utils.EncodeJSON(match), err.Error())
		return nil
	}

	result := new(model.AdBackRsp)
	json.Unmarshal(postJson, &result)
	if result.Code != 0 {
		log.Errorf("match http err=[%s]", result.Message)
		return nil
	}

	return result
}

// Register 注册
func (a *Ad) Register(userId int64) {
	user, err := dao.UserInstance().Get(userId)
	if err != nil {
		log.Errorf(" ad Register err %s", err.Error())
		return
	}
	if user == nil {
		log.Errorf("Register user nil")
		return
	}

	reqParam := make(map[string]interface{})
	reqParam["appId"] = user.AppId

	extra := make(map[string]interface{})
	extra["name"] = user.Name
	extra["phone"] = user.Phone
	extra["unionid"] = user.Unionid
	extra["user_id"] = user.Id

	reqParam["extra"] = extra

	log.Debugf("Register param %s", utils.EncodeJSON(reqParam))

	cfg := config.GetConfig().Ad
	reqUrl := cfg.Url + apiRegister
	postJson, err := utils.HttpPostJson(reqUrl, utils.EncodeJSON(reqParam), cfg.SignParam)
	if err != nil {
		log.Errorf("Register postParam=[%s] err=[%s]", utils.EncodeJSON(reqParam), err.Error())
		return
	}

	result := new(model.AdBackRsp)
	json.Unmarshal(postJson, &result)
	if result.Code != 0 {
		log.Errorf("Register http err=[%s]", result.Message)
	}
}

// Active 激活
func (a *Ad) Active(userId int64) {
	user, err := dao.UserInstance().Get(userId)
	if err != nil {
		log.Errorf(" ad Active err %s", err.Error())
		return
	}
	if user == nil {
		log.Errorf("Active user nil")
		return
	}

	reqParam := make(map[string]interface{})
	reqParam["appId"] = user.AppId

	extra := make(map[string]interface{})
	extra["name"] = user.Name
	extra["phone"] = user.Phone
	extra["unionid"] = user.Unionid
	extra["user_id"] = user.Id

	reqParam["extra"] = extra

	log.Debugf("Active param %s", utils.EncodeJSON(reqParam))

	cfg := config.GetConfig().Ad
	reqUrl := cfg.Url + apiActive
	postJson, err := utils.HttpPostJson(reqUrl, utils.EncodeJSON(reqParam), cfg.SignParam)
	if err != nil {
		log.Errorf("Active postParam=[%s] err=[%s]", utils.EncodeJSON(reqParam), err.Error())
		return
	}

	result := new(model.AdBackRsp)
	json.Unmarshal(postJson, &result)
	if result.Code != 0 {
		log.Errorf("Active http err=[%s]", result.Message)
	}
}

// Pay 支付
//func (a *Ad) Pay(orderId int64) {
//	order, err := dao.OrderInstance().Get(orderId)
//	if err != nil {
//		log.Errorf("pay order select error %s", err.Error())
//		return
//	}
//	user, err := dao.UserInstance().Get(order.UserId)
//	if err != nil {
//		log.Errorf(" ad pay err %s", err.Error())
//		return
//	}
//	if user == nil {
//		log.Errorf("pay user nil")
//		return
//	}
//	if user.AdId == 0 {
//		return
//	}
//
//	reqParam := make(map[string]interface{})
//	reqParam["relId"] = user.AdId
//	reqParam["appId"] = user.AppId
//
//	extra := make(map[string]interface{})
//	extra["payFee"] = order.TotalFee
//	extra["payType"] = order.PayType
//	extra["payLocation"] = order.Source
//
//	//var goodsInfo map[string]interface{}
//	//_ = json.Unmarshal([]byte(order.ItemsInfo), &goodsInfo)
//	var itemsInfo []model.OrderItemInfo // 存储订单中每个商品的信息
//	//goosInfo := make(map[string]interface{})
//	json.Unmarshal([]byte(order.ItemsInfo), &itemsInfo)
//	var names []string
//	var goodsIds []string
//	for _, item := range itemsInfo {
//		names = append(names, item.Name)
//		goodsIds = append(goodsIds, item.GoodsId)
//	}
//	extra["goodsId"] = cast.ToString(goodsIds)
//	extra["goodsName"] = cast.ToString(names)
//
//	reqParam["extra"] = extra
//
//	log.Debugf("pay param %s", utils.EncodeJSON(reqParam))
//
//	cfg := config.GetConfig().Ad
//	reqUrl := cfg.Url + apiPay
//	postJson, err := utils.HttpPostJson(reqUrl, utils.EncodeJSON(reqParam), cfg.SignParam)
//	if err != nil {
//		log.Errorf("pay postParam=[%s] err=[%s]", utils.EncodeJSON(reqParam), err.Error())
//		return
//	}
//
//	result := new(model.AdBackRsp)
//	json.Unmarshal(postJson, &result)
//	if result.Code != 0 {
//		log.Errorf("pay http err=[%s]", result.Message)
//	}
//}
