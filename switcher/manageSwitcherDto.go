// Package switcher 数据结构定义
package switcher

// ListPageData 列表页面数据
type ListPageData struct {
	PageTotal int         `json:"pageTotal"`        // 数据数量(分页控件内)
	ListCount int         `json:"listCount"`        // 数据数量(右上角)
	List      interface{} `json:"list"`             // 列表数据
	Extra1    interface{} `json:"extra1,omitempty"` // 额外数据1
	Extra2    interface{} `json:"extra2,omitempty"` // 额外数据2
	Extra3    interface{} `json:"extra3,omitempty"` // 额外数据3
}

// ManageLogin 管理员返回参数
type ManageLogin struct {
	ManageID int    `json:"manageid"` // 管理员id
	ShopID   int    `json:"shopid"`   // 对应的店铺id，0为超级管理员
	NickName string `json:"nickname"` // 昵称
}

// ShopInfo 店铺信息
type ShopInfo struct {
	ID              int    `json:"id"`              // 店铺id
	ShopURL         string `json:"shopurl"`         // 店铺url名称
	ShopName        string `json:"shopname"`        // 店铺名称
	LogoImg         string `json:"logoimg"`         // 店铺图标
	ShopShow        string `json:"shopshow"`        // 店铺展示
	BannerImg       string `json:"bannerimg"`       // 广告图
	BossName        string `json:"bossname"`        // 老板姓名
	TelNum          string `json:"telnum"`          // 座机号码
	PhoneNum        string `json:"phonenum"`        // 手机号码
	Wechat          string `json:"wechat"`          // 微信号
	WechatURL       string `json:"wechaturl"`       // 微信号二维码图片
	Location        string `json:"location"`        // 店铺详细地址
	Lng             string `json:"lng"`             // 店铺经度
	Lat             string `json:"lat"`             // 店铺纬度
	PopularLimit    string `json:"popularlimit"`    // 爆款点击量门限
	PopularQuantity string `json:"popularquantity"` // 爆款分类数量
	State           string `json:"state"`           // 状态
}

// ShopList 店铺列表
type ShopList struct {
	ShopID     int    `json:"shopid"`     // 店铺id
	ShopURL    string `json:"shopurl"`    // 店铺url名称
	ShopName   string `json:"shopname"`   // 店铺名称
	BossName   string `json:"bossname"`   // 老板姓名
	TelNum     string `json:"telnum"`     // 座机号码
	PhoneNum   string `json:"phonenum"`   // 手机号码
	Wechat     string `json:"wechat"`     // 微信号
	GoodsNum   string `json:"goodsnum"`   // 产品数量
	ManageNum  string `json:"managenum"`  // 管理员数量
	MachineNum string `json:"machinenum"` // 机器数量数量
	VisitNum   string `json:"visitnum"`   // 店铺访问量
	State      string `json:"state"`      // 4种状态：1为已生效，2为新建店铺还没加入管理员账号未生效，3为已停用，4为已删除
}

// Menu 店铺列表
type Menu struct {
	ShopID   int    `json:"shopid"`   // 店铺id
	SuperID  int    `json:"superid"`  // 父菜单id，为0即为1级菜单
	MenuName string `json:"menuname"` // 菜单名称
	State    string `json:"state"`    // 3种状态：1为有效，0为无效，2为删除
}

// Manager 店铺管理员列表
type Manager struct {
	ID       int    `json:"id"`       // 管理员id
	PhoneNum string `json:"phonenum"` // 账号
	Password string `json:"password"` // 密码
	Remark   string `json:"remark"`   // 备注
	ShopID   int    `json:"shopid"`   // 店铺id
	State    string `json:"state"`    // 3种状态：1为有效，0为无效，2为删除
}

// 商品列表查询
type ManageGoodsList struct {
	Total                int                    `json:"total"`     // 总量
	ManageGoodsListParam []ManageGoodsListParam `json:"goodsList"` // 商品列表
}

// 商品列表查询子集
type ManageGoodsListParam struct {
	ID        int    `json:"id"`        // 商品id
	GoodsName string `json:"goodsname"` // 商品名称
	ClassOne  string `json:"classone"`  // 一级菜单
	ClassTwo  string `json:"classtwo"`  // 二级菜单
	ClickNum  int    `json:"clicknum"`  // 点击量
	State     int    `json:"state"`     // 状态
}

// 新增产品时获取菜单分类和实物机器位置
type ManageMenuAndMachine struct {
	ManageMenuList    []ManageMenuList    `json:"menuList"`    // 菜单列表
	ManageMachineList []ManageMachineList `json:"machineList"` // 机器列表
}

// 新增产品时候的菜单子菜单
type ManageMenuList struct {
	ID        string      `json:"id"`            // 菜单ID
	MenuName  string      `json:"menuname"`      // 菜单名称
	ClassList []ClassList `json:"twoLevelClass"` // 子菜单
}

// 新增产品时候的机器子菜单
type ManageMachineList struct {
	ID       int              `json:"machineid"`   // 机器ID
	Name     string           `json:"machinename"` // 机器名
	SlotList []ManageSlotList `json:"slotlist"`    // 槽位列表
}

// 商品详情
type ManageGoodsDetail struct {
	ID           int    `json:"id"`           // 商品id
	ShopId       int    `json:"shopid"`       // 店铺id
	ClassOneId   int    `json:"classoneid"`   // 菜单id
	ClassTwoId   int    `json:"classtwoid"`   // 菜单id
	GoodsName    string `json:"goodsname"`    // 商品名称
	Brand        string `json:"brand"`        // 产品品牌
	Style        string `json:"style"`        // 风格
	Material     string `json:"material"`     // 材质
	Unit         string `json:"unit"`         // 规格
	Madein       string `json:"madein"`       // 产地
	Price        string `json:"price"`        // 价格
	Width        string `json:"width"`        // 宽度
	Height       string `json:"height"`       // 高度
	MachineId    string `json:"machineid"`    // 机器id
	MachineAdder string `json:"machineadder"` // 机器槽位号
}

// 新增产品返回ID
type ManageBackGoodsId struct {
	ID int `json:"id"` // 新增产品返回id
}

// 机器槽位子菜单
type ManageSlotList struct {
	ID  int `json:"slotid"`  // 槽位表ID
	Num int `json:"slotnum"` // 槽位表序列
}

// 颜色列表
type ManageColorList struct {
	Total     int         `json:"total"`     // 颜色总数
	ColorList []ColorList `json:"colorlist"` // 颜色列表
}

// 颜色子菜单
type ColorList struct {
	ID          int    `json:"id"`          // 颜色id
	GoodsID     int    `json:"goodsid"`     // 商品id
	Model       string `json:"model"`       // 型号
	CompressPic string `json:"compresspic"` // 缩略图
	IsFirst     int    `json:"isFirst"`     // 是否置顶
}

// 热门商品展示配置查询
type HotGoodsShow struct {
	PopularLimit    int `json:"popularlimit"`    // 爆款点击门限
	PopularQuantity int `json:"popularquantity"` // 爆款分类数量
}
