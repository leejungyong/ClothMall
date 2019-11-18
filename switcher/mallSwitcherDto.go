// Package switcher 数据结构定义
package switcher

import "github.com/gin-gonic/gin"

// Xl Dispatch的返回函数
type Xl map[string]func(*gin.Context) (string, interface{})

// AName 测试接口所需要的字段
type AName struct {
	Name string `json:"name"`
}

// AllCollection 用户收藏（收藏的商品和收藏的场景）
type AllCollection struct {
	ShopCollection   interface{} `json:"shopCollection"`   // 收藏的商品
	ScenesCollection interface{} `json:"scenesCollection"` // 收藏的场景
}

// Collection 收藏信息
type Collection struct {
	ID             int    `json:"id"`             // 收藏ID
	GoodsID        int    `json:"goodsID"`        // 商品ID
	CollectionType string `json:"collectionType"` // 收藏类型
	GoodsName      string `json:"goodsName"`      // 商品名称
	Style          string `json:"style"`          // 商品风格
	Unit           string `json:"unit"`           // 商品规格
	Price          string `json:"price"`          // 商品价格
	Width          string `json:"width"`          // 商品宽度
	Height         string `json:"height"`         // 商品高度
	CompressPic    string `json:"compressPic"`    // 商品缩略图
	ScenesURL      string `json:"scenesURL"`      // 场景链接
}

// ScenesMenu 模拟场景菜单
type ScenesMenu struct {
	ID       int           `json:"id"`       // 模拟场景菜单ID
	Name     string        `json:"name"`     // 名称
	SuperID  int           `json:"superid"`  // 父菜单id,为0即为1级菜单
	Pic      string        `json:"pic"`      // 图片URL
	Children []*ScenesMenu `json:"children"` // 二级菜单
}

// Machine 实体机器信息
type Machine struct {
	ID        int    `json:"id"`        // 机器ID
	MachineID string `json:"machineid"` // 机器序列号
	ShopID    int    `json:"shopid"`    // 店铺ID
	SlotNum   int    `json:"slotnum"`   // 机器槽位数量
	RunState  string `json:"runstate"`  // 运行状态
	NetState  string `json:"netstate"`  // 联网状态
	MachineIP string `json:"machineip"` // 机器IP
	State     string `json:"state"`     // 3种状态：1为有效，0为无效，2为删除
}

// UrlToId 根据url获取店铺id name
type UrlToId struct {
	ID       int    `json:"id"`       // 店铺id
	ShopName string `json:"shopname"` // 店铺名称
}

// 用户登陆
type UserInfo struct {
	ID int `json:"userid"` // 用户id
}

// IndexInfo 首页信息查询
type IndexInfo struct {
	ID        int        `json:"id"`        // 店铺id
	ShopName  string     `json:"shopname"`  // 店铺名
	LogoImg   string     `json:"logoimg"`   // 店铺logo
	ShopShow  string     `json:"shopshow"`  // 店铺展示图
	BannerImg string     `json:"bannerImg"` // 店铺banner
	HotGoods  []HotGoods `json:"hotgoods"`  // 热门商品查询
}

// HotCondition 首页爆款条件查询
type HotCondition struct {
	PopularLimit    int `json:"popularlimit"`    // 爆款点击量门限
	PopularQuantity int `json:"popularquantity"` // 爆款分类数量
}

// HotGoods 首页热门商品查询
type HotGoods struct {
	ID       int    `json:"id"`       // 商品id
	MenuName string `json:"menuname"` // 菜单名称
	ClickNum int    `json:"clicknum"` // 点击量
	ColorNum int    `json:"colornum"` // 颜色数量
	SmallPic string `json:"smallpic"` // 缩略图
}

// ShopContactInfo 店铺联系方式
type ShopContactInfo struct {
	Telnum    string `json:"telnum"`    // 座机号码
	Phonenum  string `json:"phonenum"`  // 手机号码
	Wechat    string `json:"wechat"`    // 微信号
	Wechaturl string `json:"wechaturl"` // 微信二维码图片
	Location  string `json:"location"`  // 店铺详细地址
	Lng       string `json:"lng"`       // 店铺经度
	Lat       string `json:"lat"`       // 店铺纬度
}

// AllClassOne 店铺全部分类查询(一级分类查询)
type AllClassOne struct {
	ID       string `json:"id"`       // 菜单ID
	MenuName string `json:"menuname"` // 菜单名称
	SuperID  string `json:"superid"`  // 父id
}

// AllClassTwo 店铺全部分类查询(二级分类查询)
type AllClassTwo struct {
	ID       string `json:"id"`       // 菜单ID
	MenuName string `json:"menuname"` // 菜单名称
	SuperID  string `json:"superid"`  // 父id
	Num      int    `json:"num"`      // 数量
}

// AllClass 店铺全部分类查询(一二级分类合并查询)
type AllClass struct {
	ID       string        `json:"id"`       // 菜单ID
	MenuName string        `json:"menuname"` // 菜单名称
	SuperID  string        `json:"superid"`  // 父id
	Children []AllClassTwo `json:"children"` // 子菜单
}

// AllClassChoice 全部菜单选择(不带颜色数量)
type AllClassChoice struct {
	ID        string      `json:"id"`        // 菜单ID
	MenuName  string      `json:"menuname"`  // 菜单名称
	ClassList []ClassList `json:"classList"` // 子菜单
}

// ClassList 菜单选择子集(不带颜色数量)
type ClassList struct {
	ID       string `json:"id"`       // 菜单ID
	MenuName string `json:"menuname"` // 菜单名称
}

// GoodsSearch 二级分类下商品查询
type GoodsSearch struct {
	Total     int         `json:"total"`     // 二级菜单全部数量
	GoodsList []GoodsList `json:"goodslist"` // 子菜单
}

// GoodsList 二级分类下商品查询
type GoodsList struct {
	ID       int    `json:"id"`       // 商品id
	ClickNum int    `json:"clicknum"` // 点击量
	SmallPic string `json:"smallpic"` // 商品图片
	ColorNum int    `json:"colornum"` // 颜色数量
}

// GoodsDetail 商品详细信息
type GoodsDetail struct {
	ID            int           `json:"id"`            // 商品id
	GoodsName     string        `json:"goodsname"`     // 商品名称
	Brand         string        `json:"brand"`         // 商品品牌
	Style         string        `json:"style"`         // 商品风格
	Unit          string        `json:"unit"`          // 规格
	Material      string        `json:"material"`      // 材质
	MadeIn        string        `json:"madein"`        // 产地
	Price         string        `json:"price"`         // 价格
	Width         string        `json:"width"`         // 宽度
	Height        string        `json:"height"`        // 高度
	MachineDetail MachineDetail `json:"machinedetail"` // 机器详情
	PicList       []PicList     `json:"picList"`       // 图片数组
	IsCollection  bool          `json:"iscollection"`  // 是否收藏
	CollectionID  int           `json:"collectionid"`  // 收藏id
}

// PicList 商品图片信息 子集
type PicList struct {
	ID          int    `json:"id"`          // 商品图片id
	Model       string `json:"model"`       // 型号
	Pic         string `json:"pic"`         // 原图URL
	Compresspic string `json:"compresspic"` // 缩略图小图URL
}

// MachineDetail 机器的详情 槽位的详情
type MachineDetail struct {
	MachineID   int    `json:"machineid"`   // 机器id
	MachineName string `json:"machinename"` // 机器名称
	SlotID      int    `json:"slotid"`      // 槽位id
	SlotNum     int    `json:"slotnum"`     // 槽位表序号
}
