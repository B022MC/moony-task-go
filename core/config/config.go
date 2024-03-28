package config

import (
	"github.com/mitchellh/mapstructure"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"moony-task-go/utils"
	"os"
	"strings"
)

var (
	config    *Config
	configEnv string
)

// Mysql 数据库配置
type Mysql struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
	User string `toml:"user"`
	Pass string `toml:"pass"`
	Db   string `toml:"db"`
}

// Server 服务连接配置
type Server struct {
	Host            string   `toml:"host"`
	AppId           int64    `toml:"app_id"`
	Address         string   `toml:"address"`
	LogLevel        int      `toml:"log_level"`
	EnableDoc       bool     `toml:"enable_doc"`
	TokenIgnore     []string `toml:"token_ignore"`
	RequestLimitNum int64    `toml:"request_limit_num"`
}

// Redis 配置
type Redis struct {
	Addr     string `toml:"addr"`
	Db       int    `toml:"db"`
	Password string `toml:"password"`
}

// WxLogin 微信登录配置
type WxLogin struct {
	AppId     string `toml:"app_id"`
	AppSecret string `toml:"app_secret"`
}

// Qiniu 七牛云配置
//type Qiniu struct {
//	Path      string `toml:"path"`
//	Bucket    string `toml:"bucket"`
//	Domain    string `toml:"domain"`
//	SavePath  string `toml:"save_path"`
//	AccessKey string `toml:"access_key"`
//	SecretKey string `toml:"secret_key"`
//}

// Qywx 企业微信配置
//type Qywx struct {
//	Name      string   `toml:"name"`
//	CorpId    string   `toml:"corp_id"`
//	AgentId   int      `toml:"agent_id"`
//	Secret    string   `toml:"secret"`
//	Receivers []string `toml:"receivers"`
//}

// Pay 支付配置
//type Pay struct {
//	WxpayAppid  string `toml:"wxpay_appid"`
//	AlipayAppid string `toml:"alipay_appid"`
//}

// Elastic 链接配置
//type Elastic struct {
//	Address       string `toml:"address"`
//	Username      string `toml:"username"`
//	Password      string `toml:"password"`
//	IndexName     string `toml:"index_name"`      //接单信息索引名称
//	IndexUserCard string `toml:"index_user_card"` //用户名片索引名称
//}

// AliCloudSms 阿里云短信配置
//type AliCloudSms struct {
//	SignName        string `toml:"sign_name"`         //签名
//	TemplateCode    string `toml:"template_code"`     //模版编号
//	AccessKeyId     string `toml:"access_key_id"`     //授权ID
//	AccessKeySecret string `toml:"access_key_secret"` //授权secret
//}

// Ad 广告请求配置
type Ad struct {
	Url       string `toml:"url"`
	DataUrl   string `toml:"data_url"`
	SignParam string `toml:"sign_param"`
}

// Config 配置参数
type Config struct {
	Ad      *Ad      `toml:"ad"`
	Server  *Server  `toml:"server"`
	Mysql   *Mysql   `toml:"mysql"`
	Redis   *Redis   `toml:"redis"`
	WxLogin *WxLogin `toml:"wx_login"`
	//Qiniu       *Qiniu       `toml:"qiniu"`
	//Pay         *Pay         `toml:"pay"` //支付配置
	//Qywx        *Qywx        `toml:"qywx"`
	//Elastic     *Elastic     `toml:"elastic"`
	//AliCloudSms *AliCloudSms `toml:"ali_cloud_sms"`
}

func GetEnv() string {
	return configEnv
}

func IsProdEnv() bool {
	return configEnv == "prod"
}

func IsDevEnv() bool {
	return configEnv == "dev"
}

func IsTestEnv() bool {
	return configEnv == "test"
}

func GetConfig() *Config {
	return config
}

// InitConfig 初始化系统配置
func InitConfig() {
	configEnv = os.Getenv("CONFIG_ENV")
	if configEnv == "" {
		configEnv = "dev"
	}

	var envConfig Config

	viper.SetConfigFile("config/server.conf." + configEnv)
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("conf")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&envConfig, func(decoderConfig *mapstructure.DecoderConfig) {
		decoderConfig.TagName = "toml"
	}); err != nil {
		panic(err)
	}

	config = &envConfig

	// 支付配置
	//if config.Pay.WxpayAppid == "" || config.Pay.AlipayAppid == "" {
	//	panic("wxpay appid nil or alipay appid nil")
	//}

	// 初始化支付配置
	//InitPayConfig(config.Pay.WxpayAppid, config.Pay.AlipayAppid)

	log.Infof("load real config[%s] ", utils.EncodeJSONIndent(config))
}
