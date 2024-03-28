package utils

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"io/ioutil"
	"math"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func BytesToTips(bytes uint64) string {
	switch {
	case bytes < 1024:
		return fmt.Sprintf("%dB", bytes)
	case bytes < 1024*1024:
		return fmt.Sprintf("%.2fK", float64(bytes)/1024)
	case bytes < 1024*1024*1024:
		return fmt.Sprintf("%.2fM", float64(bytes)/1024/1024)
	default:
		return fmt.Sprintf("%.2fG", float64(bytes)/1024/1024/1024)
	}
}

func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func CopyStruct(dst interface{}, src interface{}) {

	dtype := reflect.TypeOf(dst)
	stype := reflect.TypeOf(src)

	if stype.Kind() != reflect.Ptr || stype.Kind() != dtype.Kind() {
		panic(errors.New("src/dst must ptr"))
	}
	if reflect.ValueOf(dst).IsNil() || reflect.ValueOf(src).IsNil() {
		panic(errors.New("src/dst is nil"))
	}

	dval := reflect.ValueOf(dst).Elem()
	sval := reflect.ValueOf(src).Elem()

	for i := 0; i < sval.NumField(); i++ {
		sValue := sval.Field(i)

		dValue := dval.FieldByName(sval.Type().Field(i).Name)
		if sValue.IsZero() || dValue.IsValid() == false || !dValue.CanSet() {
			continue
		}
		if sValue.Kind() != dValue.Kind() {
			continue
		}

		switch sValue.Type().Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			dValue.SetInt(sValue.Int())

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			dValue.SetUint(sValue.Uint())

		case reflect.Float32, reflect.Float64:
			dValue.SetFloat(sValue.Float())

		case reflect.String:
			dValue.SetString(sValue.String())

		case reflect.Bool:
			dValue.SetBool(sValue.Bool())

		case reflect.Ptr:
			CopyStruct(dValue.Interface(), sValue.Interface())
		}
	}
}

// ParseModelByUa 从ua里获取机型
func ParseModelByUa(ua string) string {
	var model string
	if strings.Contains(ua, "Windows") {
		return model
	}
	ua = strings.Replace(ua, "; wv", "", -1)
	re, _ := regexp.Compile(";\\s?(\\S*?\\s?\\S*?)\\s?(Build)?/")
	res := re.FindString(ua)
	if res != "" {
		res = strings.Replace(res, "; ", "", -1)
		res = strings.Replace(res, "zh-cn", "", -1)
		res = strings.Replace(res, " Build/", "", -1)
		model = res
	}

	return model
}

// TernaryOperator 三元表达式
func TernaryOperator[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// FormatMoney 格式化商品价格
func FormatMoney(number int64) string {
	num1 := float64(number) / 100
	num2 := cast.ToFloat64(number / 100)
	if num1 != num2 {
		return fmt.Sprintf("%.2f", num1)
	}
	return cast.ToString(cast.ToInt64(num1))
}

/*
RandAllString  生成随机字符串([a~zA~Z0~9])

	lenNum 长度
*/
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

func KeyInMapStringValue(m []string, s string) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}

func KeyInMapIntValue(m []int, s int) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}

func KeyInMapInt64Value(m []int64, s int64) bool {
	for _, v := range m {
		if v == s {
			return true
		}
	}
	return false
}

func Decimal(num float64) float64 {
	num, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", num), 64)
	return num
}

func HtmlStrip(src string) string {
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{1,}")
	src = re.ReplaceAllString(src, "")

	//去除&#12345;这类字符
	//re, _ = regexp.Compile("&#\\d*;")
	//src = re.ReplaceAllString(src, "")

	src = strings.ReplaceAll(src, "&nbsp;", "")
	src = strings.ReplaceAll(src, "nbsp;", "")
	src = strings.ReplaceAll(src, "& nbsp;", "")
	src = strings.ReplaceAll(src, "&nbsp", "")
	return strings.TrimSpace(src)
}

// ParseMoney 格式化金额为万元
func ParseMoney(m int64) string {
	if m == 0 {
		return ""
	} else {
		money := float64(m) / 10000
		moneyStr := strconv.FormatFloat(money, 'f', 4, 64)

		for strings.HasSuffix(moneyStr, "0") {
			moneyStr = strings.TrimSuffix(moneyStr, "0")
		}
		if strings.HasSuffix(moneyStr, ".") {
			moneyStr = strings.TrimSuffix(moneyStr, ".")
		}

		return moneyStr
	}
}

// GetDistance 返回单位为：千米
func GetDistance(lat1, lat2, lng1, lng2 float64) float64 {
	radius := 6371000.0 //6378137.0
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	return dist * radius / 1000
}

func KeywordReplace(s string) string {
	s = strings.ReplaceAll(s, "招聘", "接单")
	s = strings.ReplaceAll(s, "直聘", "接单")

	return s
}

// RandomInt 随机数
func RandomInt(start int, end int) int {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(end - start)
	random = start + random
	return random
}

// GenerateRandomString 生成随机字符串
func GenerateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// CheckPhone 手机号验证
func CheckPhone(phone string) bool {
	regular := "^[1]\\d{10}$"
	reg := regexp.MustCompile(regular)
	return reg.MatchString(phone)
}

// GenerateUUID 生成32位无连字符的UUID字符串
func GenerateUUID() string {
	uuidWithHyphens := uuid.New()
	return strings.ReplaceAll(uuidWithHyphens.String(), "-", "")
}

// GetCurrentUnixMillis 获取当前时间的Unix时间戳（毫秒）
func GetCurrentUnixMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

// ImageURLToBase64 从URL下载图片并转换为Base64编码的字符串
func ImageURLToBase64(url string) (string, error) {
	// 发送GET请求获取图片
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 读取响应的字节
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 转换为Base64
	base64Str := base64.StdEncoding.EncodeToString(imageData)

	return base64Str, nil
}

// IsEmpty 检查任何结构体是否为空
func IsEmpty[T any](v T) bool {
	// 使用反射创建类型T的零值实例
	zeroValue := reflect.Zero(reflect.TypeOf(v)).Interface()
	return reflect.DeepEqual(v, zeroValue)
}
