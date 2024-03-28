package common

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"math"
	"moony-task-go/common/dao"
	"moony-task-go/common/model"
	"moony-task-go/core/config"
	"moony-task-go/utils"
	"sort"
	"strings"
)

var (
	ExpServerGoodsList        = "server.goods.list"   //商品列表
	ExpServerUserDefault      = "server.user.default" //用户默认信息配置
	ExpServerLoginPhoneVerify = "server.login.phone.verfiy"
	ClientPayEnableKey        = "client.pay.enable" //是否开启付费key
	ClientWxpayCfgKey         = "client.wxpay.cfg"  //微信支付配置
	ServerAlipayCfgKey        = "server.alipay.cfg" //支付宝支付配置
)

type Experiment struct {
	hits   map[string]interface{}
	params map[string]interface{}
}

func ExperimentInstance() *Experiment {
	return &Experiment{}
}

// LoadExperiment 加载试验配置
func (e *Experiment) LoadExperiment(session *Session) error {
	e.hits = make(map[string]interface{})
	e.params = make(map[string]interface{})

	if session == nil {
		return nil
	}

	// 获取系统默认配置
	serverDefaultConfig := config.GetServerConfig(session.UserHeader.Platform)
	clientDefaultConfig := config.GetClientConfig(session.UserHeader.Platform)
	for k, v := range serverDefaultConfig {
		e.params[k] = v
	}
	for k, v := range clientDefaultConfig {
		e.params[k] = v
	}

	exps, err := dao.ConfigDaoInstance().GetAll(session.UserHeader.AppId)
	if err != nil {
		log.Errorf("load error [%s]", err.Error())
		return err
	}
	sort.Sort(sort.Reverse(model.SortExp(exps)))

	// 实验配置
	for _, exp := range exps {
		if exp.Channel != "" && !utils.StrInField(session.UserHeader.Channel, exp.Channel) {
			continue
		}
		if exp.Version != "" && !utils.StrInField(session.UserHeader.Version, exp.Version) {
			continue
		}
		if exp.Platform != "" && strings.ToLower(exp.Platform) != session.UserHeader.Platform {
			continue
		}

		if exp.EndBucket < 0 || exp.EndBucket >= 100 || exp.StartBucket < 0 || exp.StartBucket >= 100 || exp.EndBucket < exp.StartBucket {
			continue
		}

		hashId := session.GetHashId()
		bucketId := e.GetHashNum(hashId, exp.Layer)
		if bucketId > exp.EndBucket || bucketId < exp.StartBucket {
			continue
		}
		var params map[string]interface{}
		if err := json.Unmarshal([]byte(exp.Params), &params); err != nil {
			log.Errorf("params[%s] format error %s", exp.Params, err.Error())
			continue
		}
		e.hits[exp.Name] = params
		for k, v := range params {
			e.params[k] = v
		}
	}

	// 用户个性化配置
	if session.UserToken != nil && session.UserToken.UserId != 0 {
		user, err := dao.UserInstance().Get(session.UserToken.UserId)
		if err != nil {
			log.Errorf("GetUser[%d] error [%s]", session.UserToken.UserId, err.Error())
			return err
		}
		if user != nil {
			var userConfig map[string]interface{}
			json.Unmarshal([]byte(user.Config), &userConfig)
			for k, v := range userConfig {
				e.params[k] = v
			}
		}
	}

	return nil
}

func (e *Experiment) Get(key string) interface{} {
	return e.params[key]
}

func (e *Experiment) GetClient() map[string]interface{} {
	params := make(map[string]interface{})
	for k, v := range e.params {
		if !strings.HasPrefix(k, "client.") {
			continue
		}
		params[k] = v
	}
	return params
}

func (e *Experiment) GetServer() map[string]interface{} {
	params := make(map[string]interface{})
	for k, v := range e.params {
		if !strings.HasPrefix(k, "server.") {
			continue
		}
		params[k] = v
	}
	return params
}

func (e *Experiment) GetHits() map[string]interface{} {
	return e.hits
}

func (e *Experiment) GetAppConfigSubValue(key string, subKey string) interface{} {
	appConfig := cast.ToStringMap(e.Get(key))
	if appConfig != nil {
		if _, ok := appConfig[subKey]; !ok {
			return nil
		}
		return appConfig[subKey]
	}
	return nil
}

func (e *Experiment) GetHashNum(hashId uint64, layer int) int {
	hashNum := uint64(hashId/uint64(math.Pow(100, float64(layer-1)))) % 100

	return cast.ToInt(hashNum)
}
