package config

import (
	"encoding/json"
	"os"
	"strings"
)

var (
	serverIosConfig     map[string]interface{} // ios 服务端配置
	serverAndroidConfig map[string]interface{} // android 服务端配置
)

// InitServerConfig 初始化服务端配置
func InitServerConfig() {
	serverIosConfig = make(map[string]interface{})
	serverAndroidConfig = make(map[string]interface{})

	body, err := os.ReadFile("config/server/server.ios.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &serverIosConfig); err != nil {
		panic(err)
	}

	body, err = os.ReadFile("config/server/server.android.json")
	if err != nil {
		panic(err)
	}
	if err = json.Unmarshal(body, &serverAndroidConfig); err != nil {
		panic(err)
	}

	// 检查环境变量
	for k, _ := range serverIosConfig {
		envValue := os.Getenv(k)
		if envValue != "" {
			var inf interface{}
			_ = json.Unmarshal([]byte(envValue), &inf)
			serverIosConfig[k] = inf
		}
	}

	// 检查环境变量
	for k, _ := range serverAndroidConfig {
		envValue := os.Getenv(k)
		if envValue != "" {
			var inf interface{}
			_ = json.Unmarshal([]byte(envValue), &inf)
			serverAndroidConfig[k] = inf
		}
	}
}

// GetServerConfig 获取配置信息
func GetServerConfig(platform string) map[string]interface{} {
	serverConfig := serverIosConfig

	if strings.ToLower(platform) == AndroidPlatform {
		serverConfig = serverAndroidConfig
	}
	result := make(map[string]interface{})
	for k, v := range serverConfig {
		result[k] = v
	}
	return result
}
