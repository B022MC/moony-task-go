package service

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gogap/errors"
	"github.com/spf13/cast"
	"io"
	"moony-task-go/common/model"
	"moony-task-go/core/config"
	"net/http"
	"sort"
	"strings"
)

// LocationService provides functionality for interacting with the location cloud API.
type LocationService struct {
}

// NewLocationService creates a new instance of LocationService with the provided API and secret keys.
func NewLocationService() *LocationService {
	return &LocationService{}
}

// GetAddressByCoordinates 根据坐标获取地址信息
func (ls *LocationService) GetAddressByCoordinates(location model.Location) (*model.AddressByCoordinatesResponse, error) {
	// 定义 SecretKey 和请求参数
	cfg := config.GetConfig().TencentLocation
	secretKey := cfg.SecretKey
	params := map[string]string{
		"location": cast.ToString(location.Lat) + "," + cast.ToString(location.Lng),
		"key":      cfg.Key,
	}

	// 对参数进行排序
	sortedParams := make([]string, 0, len(params))
	for k := range params {
		sortedParams = append(sortedParams, k)
	}
	sort.Strings(sortedParams)

	// 构建参数字符串
	var paramString string
	for _, k := range sortedParams {
		paramString += k + "=" + params[k] + "&"
	}
	paramString = strings.TrimSuffix(paramString, "&")

	// 构建待签名字符串
	signatureString := "/ws/geocoder/v1?" + paramString + secretKey

	// 计算 MD5 值
	hash := md5.Sum([]byte(signatureString))
	signature := hex.EncodeToString(hash[:])

	// 构建最终请求 URL
	finalURL := "https://apis.map.qq.com/ws/geocoder/v1?" + paramString + "&sig=" + signature

	// 实际执行 HTTP 请求
	resp, err := http.Get(finalURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("HTTP request failed with status: " + resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response failed: %w", err)
	}
	var response *model.AddressByCoordinatesResponse
	// 在这里处理响应结果
	// 这里只是一个示例，你需要根据实际情况解析响应体并返回相应的结果
	// 假设响应体是一个 JSON 格式的字符串，你可以使用相应的库来解析它
	// 比如，可以使用 encoding/json 包来解析 JSON 字符串
	// 请根据你的实际情况进行修改
	// 示例代码：

	if err := json.Unmarshal(body, &response); err != nil {
		return nil, fmt.Errorf("json unmarshal failed: %w", err)
	}

	if response.Status != 0 {
		return nil, fmt.Errorf("API error: %s", response.Message)
	}

	return response, nil
}
