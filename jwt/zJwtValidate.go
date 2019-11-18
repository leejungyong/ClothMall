// Package jwt gin的中间件
package jwt

import (
	sw "ClothMall/switcher"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// JwtValidate jwt token 认证
func JwtValidate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenstr string
		// 获取token
		tokenstr = getjwtTokenFromGin(c)
		// token解析并检查
		claim, err := parseAndValidateClaims(tokenstr, c)
		// 有token超时之外的错误
		if nil != err && TokenExpired != err {
			log.Println(err)
			// token塞回header中
			c.Header(JWTTokenKey, tokenstr)
			c.AbortWithStatusJSON(http.StatusForbidden, Rtn{false, err.Error(), nil})
		}
		// token超时
		if TokenExpired == err {
			// 刷新token
			token := claim.refleshToken()
			tokenstr = token.Token
		}
		// token塞回header中
		c.Header(JWTTokenKey, tokenstr)
		c.Next()
	}
}

// 获取jwt的token
func getjwtTokenFromGin(c *gin.Context) string {
	var token string
	// 先查找token是否写入body或者url中
	token = sw.GetParam(c, JWTTokenKey)
	// 如果为空
	if "" == token {
		// 从header中查找
		token = c.GetHeader(JWTTokenKey)
	}
	// 如果还是为空
	if "" == token {
		// cookie中查找
		token, _ = c.Cookie(JWTTokenKey)
	}
	return token
}
