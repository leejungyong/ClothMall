// Package jwt jwt的代码
package jwt

import (
	sw "ClothMall/switcher"
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	// JWTTokenKey token放进去的key
	JWTTokenKey = "X-JWT-TOKEN"
)

// 自定义的jwt token错误
var (
	// TokenExpired token过期
	TokenExpired error = errors.New("Token is expired")
	// TokenNotValidYet token非法
	TokenNotValidYet error = errors.New("Token not active yet")
	// TokenMalformed 根本不是token
	TokenMalformed error = errors.New("That's not even a token")
	// TokenInvalid 无法识别此token
	TokenInvalid error = errors.New("Couldn't handle this token:")
	// TokenInvalidIp token中的IP地址验证失败
	TokenInvalidIp error = errors.New("Invalid ip address")
)

// 新建claim
func newMyClaims(username, ipaddr string) *MyClaims {
	claims := MyClaims{
		User:   username, // 用户名
		IpAddr: ipaddr,   // 发送请求的ip地址
	}
	// 生效时间
	claims.NotBefore = time.Now().Unix()
	// 过期时间
	claims.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()
	return &claims
}

// 生成token
func (c *MyClaims) getToken() *JwtToken {
	// 新建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, c)
	// 加密生成token字符串
	tokenstr, _ := token.SignedString([]byte(sw.ConstVar["tokenKey"]))
	return &JwtToken{
		Token:   tokenstr,
		Expired: c.ExpiresAt,
	}
}

// 刷新token
func (c *MyClaims) refleshToken() *JwtToken {
	// 刷新生效时间和过期时间
	c.NotBefore = time.Now().Unix()
	c.ExpiresAt = time.Now().Add(time.Hour * 1).Unix()
	// 生成token
	return c.getToken()
}

// 解析token并检查错误
func parseAndValidateClaims(tokenstr string, c *gin.Context) (*MyClaims, error) {
	// 解析claims
	claim, err := parseClaims(tokenstr)
	if nil != err && TokenExpired != err {
		log.Println(err)
		return nil, err
	}
	// jwt标准claim解析时没有错误 或者 是token过期的错误
	if nil == err || TokenExpired == err {
		log.Printf("Ip address from token:%s, ip we get:%s", claim.IpAddr, c.ClientIP())
		// 判断自定义错误
		// 请求的ip地址是否一致
		if claim.IpAddr != c.ClientIP() {
			log.Printf("Ip address is not the same, expect:%s, get:%s", claim.IpAddr, c.ClientIP())
			return nil, TokenInvalidIp
		}
	}
	return claim, err
}

// 解析claims
func parseClaims(tokenstr string) (*MyClaims, error) {
	// 解析claims
	token, err := jwt.ParseWithClaims(tokenstr, &MyClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return []byte(sw.ConstVar["tokenKey"]), nil
	})
	// 判断出错信息
	if err != nil {
		log.Println(err)
		// 错误信息断言成jwt包的错误信息
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	// 判断解析出来的claim是否合法
	if t2, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return t2, nil
	} else {
		return nil, TokenInvalid
	}
}
