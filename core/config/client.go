package config

import (
	"encoding/json"
	"os"
	"strings"
)

var (
	clientIosConfig     map[string]interface{} //ios 客户端配置
	clientAndroidConfig map[string]interface{} //android 客服端配置
)

// InitClientConfig 初始化客户端配置
func InitClientConfig() {
	clientIosConfig = make(map[string]interface{})
	clientAndroidConfig = make(map[string]interface{})

	body, err := os.ReadFile("config/client/client.ios.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &clientIosConfig); err != nil {
		panic(err)
	}

	body, err = os.ReadFile("config/client/client.android.json")
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &clientAndroidConfig); err != nil {
		panic(err)
	}

	// 检查环境变量
	for k, _ := range clientIosConfig {
		envValue := os.Getenv(k)
		if envValue != "" {
			var inf interface{}
			_ = json.Unmarshal([]byte(envValue), &inf)
			clientIosConfig[k] = inf
		}
	}
	for k, _ := range clientAndroidConfig {
		envValue := os.Getenv(k)
		if envValue != "" {
			var inf interface{}
			_ = json.Unmarshal([]byte(envValue), &inf)
			clientAndroidConfig[k] = inf
		}
	}
}

// GetClientConfig 获取客服端配置信息
func GetClientConfig(platform string) map[string]interface{} {
	clientConfig := clientIosConfig

	if strings.ToLower(platform) == AndroidPlatform {
		clientConfig = clientAndroidConfig
	}
	result := make(map[string]interface{})
	for k, v := range clientConfig {
		result[k] = v
	}
	return result
}
