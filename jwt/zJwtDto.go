// Package jwt 数据结构定义
package jwt

import (
    "github.com/dgrijalva/jwt-go"
)

// 自定义的claim
type MyClaims struct {
    User        string  `json:"user"`
    IpAddr      string  `json:"ipaddr"`
    jwt.StandardClaims
}

// Rtn 返回参数
type Rtn struct {
    Success bool        `json:"success"`
    Errmsg  string      `json:"errmsg"`
    Data    interface{} `json:"data"`
}

// JwtToken jwt token
type JwtToken struct {
    Token       string  `json:"token"`  // 生成的token
    Expired     int64   `json:"expired"`// token过期时间 unix时间戳
}