// Package jwt main函数中路由指向的句柄函数
package jwt

import (
	sw "ClothMall/switcher"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginHandler 登陆并且获取token
func LoginHandler(c *gin.Context) {
	// 用户名
	username := sw.GetParamWithPanic(c, "username")
	// 请求地址
	ipaddr := c.ClientIP()
	// 生成claim
	claim := newMyClaims(username, ipaddr)
	// 生成token
	token := claim.getToken()
	c.Header(JWTTokenKey, token.Token)
	// 拼返回数据
	rtn := &Rtn{true, "获取token成功", token}
	// 返回请求的数据
	dataResponse(c, rtn)
}

// 返回请求的数据
func dataResponse(c *gin.Context, rtn *Rtn) {
	_, exists := c.GetQuery("callback")
	// 有callback参数时 采用jsonp返回数据
	if exists {
		c.JSONP(http.StatusOK, rtn)
	} else {
		// 没有callback参数时 采用json返回数据
		c.JSON(http.StatusOK, rtn)
	}
}
