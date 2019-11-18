// Package switcher 公共函数定义
package switcher

import (
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetParamWithPanic 获取请求参数 如果没有直接抛异常
func GetParamWithPanic(c *gin.Context, key string) string {
	v := GetParam(c, key)
	if "" == v {
		panic("没有参数" + key)
	} else {
		return v
	}
}

// GetParam 获取请求中的参数
func GetParam(c *gin.Context, key string) string {
	if "POST" == c.Request.Method && "" != c.PostForm(key) {
		return c.PostForm(key)
	} else if "GET" == c.Request.Method || ("POST" == c.Request.Method && "" != c.Query(key)) {
		return c.Query(key)
	} else {
		return ""
	}
}

// GetParamDefault 获取带默认值的请求参数
func GetParamDefault(c *gin.Context, key, defaultValue string) string {
	if "GET" == c.Request.Method {
		return c.DefaultQuery(key, defaultValue)
	} else if "POST" == c.Request.Method {
		return c.DefaultPostForm(key, defaultValue)
	} else {
		return defaultValue
	}
}

// GetParamExist 获取请求参数并判断请求参数是否存在
func GetParamExist(c *gin.Context, key string) (string, bool) {
	if "GET" == c.Request.Method {
		return c.GetQuery(key)
	} else if "POST" == c.Request.Method {
		return c.GetPostForm(key)
	} else {
		return "", false
	}
}

// Perror 打印出错信息并抛出异常
func Perror(err error, errMsg string) {
	if nil != err {
		log.Println(err)
		panic(errMsg)
	}
}

// ReadFile 读取文件
func ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

// IntToString int转为string
func IntToString(i int) (s string) {
	str := strconv.Itoa(i)
	return str
}

// Int64ToString int64转为string
func Int64ToString(i int64) (s string) {
	str := strconv.FormatInt(i, 10)
	return str
}

// FloatToString float64转为string
func FloatToString(f float64) (s string) {
	str := strconv.FormatFloat(f, 'f', -1, 64)
	return str
}

// StringToFloat64 string转为float64
func StringToFloat64(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

// StringToInt string转为int
func StringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

// StringToInt64 string转为int64
func StringToInt64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
