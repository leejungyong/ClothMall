package upload

import (
    "crypto/md5"
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "encoding/json"
    "io"
    "net/http"

    "github.com/gin-gonic/gin"
)

// Rtn 返回数据
type Rtn struct {
    Success bool        `json:"success"`
    Errmsg  string      `json:"errmsg"`
    Data    interface{} `json:"data"`
}

// Pictures 返回的图片数据结构
type Pictures struct {
    Key   string `json:"key"`
    Image string `json:"image"`
    Thumb string `json:"thumb"`
}

// Size 接收的图片缩放等的参数
type Size struct {
    Key       string `json:"key"`
    Thumbonly int    `json:"thumbonly"`
    Width     int    `json:"width"`
    Height    int    `json:"height"`
}

// 返回请求的数据
func dataResponse(c *gin.Context, rtn *Rtn) {
    callback, exists := c.GetQuery("callback")
    if exists {
        c.String(http.StatusOK, GetJsonpResult(callback, rtn))
    } else {
        c.JSON(http.StatusOK, rtn)
    }
}

// GetJsonpResult 返回数据处理成jsonp
func GetJsonpResult(callback string, rtn *Rtn) string {
    respstr, err := json.Marshal(rtn)
    if nil != err {
        panic(err)
    }
    return callback + "(" + string(respstr) + ")"
}

// GetGUID 随机文件名
func GetGUID() string {
    b := make([]byte, 48)
    if _, err := io.ReadFull(rand.Reader, b); err != nil {
        panic("创建文件名出错")
    }
    return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

// GetMd5String md5字符串
func GetMd5String(s string) string {
    h := md5.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}