// Package switcher 常量定义
package switcher

import (
    
)

// 常量部分
const (
    // 产生随机字符串用的常量
    // 需要产生的字符串类型
    // 纯数字
    KCRandKindNum = 0
    // 小写字母
    KCRandKindLower = 1
    // 大写字母
    KCRandKindUpper = 2
    // 数字、大小写字母
    KCRandKindAll = 3
)
// ConstVar 常量定义
var ConstVar map[string]string