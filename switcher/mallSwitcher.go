// Package switcher cmd的函数
package switcher

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MallDispatch 根据cmd分发到不同的接口函数
func MallDispatch(db *sql.DB) Xl {
	return Xl{
		// 测试接口
		"test1": func(c *gin.Context) (string, interface{}) {
			id := GetParam(c, "id")
			log.Println(id)
			data, _ := testDao(id, db)
			return "test success", data
		},

		// 查询用户收藏信息（收藏的商品和收藏的场景）
		"getCollection": func(c *gin.Context) (string, interface{}) {
			// 用户ID
			userid, _ := strconv.Atoi(GetParam(c, "userid"))
			// 店铺ID
			shopid, _ := strconv.Atoi(GetParam(c, "shopid"))
			data, err := getCollection(userid, shopid, db)
			Perror(err, "查询用户收藏信息失败")
			return "查询用户收藏信息成功", data
		},

		// 用户增加收藏
		"insertCollection": func(c *gin.Context) (string, interface{}) {
			// 用户ID
			userid, _ := strconv.Atoi(GetParam(c, "userid"))
			// 店铺ID
			shopid, _ := strconv.Atoi(GetParam(c, "shopid"))
			// 商品ID
			goodsid, _ := strconv.Atoi(GetParam(c, "goodsid"))
			// 收藏类型
			collectionType := GetParam(c, "collectionType")
			// 场景链接
			scenesURL := GetParam(c, "scenesURL")

			data, err := insertCollection(userid, shopid, goodsid, collectionType, scenesURL, db)
			Perror(err, "增加收藏失败")
			return "增加收藏成功", data
		},

		// 用户取消收藏
		"delCollection": func(c *gin.Context) (string, interface{}) {
			// 用户ID
			userid, _ := strconv.Atoi(GetParam(c, "userid"))
			// 收藏ID
			collectionid, _ := strconv.Atoi(GetParam(c, "collectionid"))

			err := delCollection(userid, collectionid, db)
			Perror(err, "取消收藏失败")
			return "取消收藏成功", ""
		},

		// 查询模拟场景菜单
		"getMallScenesMenu": func(c *gin.Context) (string, interface{}) {
			// 父菜单id
			superid := GetParam(c, "superid")

			data, err := getMallScenesMenuList(superid, db)
			Perror(err, "查询模拟场景菜单失败")
			return "查询模拟场景菜单成功", data
		},

		// 实体机器状态检查
		"machineCheck": func(c *gin.Context) (string, interface{}) {
			// 机器ID
			machineid, _ := strconv.Atoi(GetParam(c, "machineid"))
			// 用户IP地址
			userIP := GetParam(c, "userip")

			data, err := machineCheck(machineid, userIP, db)
			Perror(err, "实体机器状态检查失败")
			if data == "3" {
				panic("用户ip与机器ip不同")
			}
			if data == "4" {
				panic("该机器目前未联网")
			}
			if data == "5" {
				panic("该机器目前正在运行")
			}
			return "该实体机器可用", "该实体机器可用"
		},

		// 增加机器展示任务
		"insertMachineTask": func(c *gin.Context) (string, interface{}) {
			// 机器ID
			machineid, _ := strconv.Atoi(GetParam(c, "machineid"))
			// 用户ID
			userid, _ := strconv.Atoi(GetParam(c, "userid"))
			// 店铺ID
			shopid, _ := strconv.Atoi(GetParam(c, "shopid"))
			// 商品ID
			goodsid, _ := strconv.Atoi(GetParam(c, "goodsid"))
			// 机器槽位号
			slotid, _ := strconv.Atoi(GetParam(c, "slotid"))

			err := insertMachineTask(machineid, userid, shopid, goodsid, slotid, db)
			Perror(err, "新增机器展示任务失败")
			return "新增机器展示任务成功", ""
		},

		// 增加模拟场景任务
		"insertScenesTask": func(c *gin.Context) (string, interface{}) {
			// 店铺ID
			shopid, _ := strconv.Atoi(GetParam(c, "shopid"))
			// 用户ID
			userid, _ := strconv.Atoi(GetParam(c, "userid"))
			// 模拟场景菜单场景id
			scenesmenuid1, _ := strconv.Atoi(GetParam(c, "scenesmenuid1"))
			// 模拟场景场景图片URL
			scenesmenuurl1 := GetParam(c, "scenesmenuurl1")
			// 模拟场景菜单房间id
			scenesmenuid2, _ := strconv.Atoi(GetParam(c, "scenesmenuid2"))
			// 模拟场景房间图片URL
			scenesmenuurl2 := GetParam(c, "scenesmenuurl2")
			// 商品Aid
			goodsAid, _ := strconv.Atoi(GetParam(c, "goodsAid"))
			// 商品A图片URL
			goodsAurl := GetParam(c, "goodsAurl")
			// 商品A宽度
			goodsAwidth := GetParam(c, "goodsAwidth")
			// 商品A高度
			goodsAheight := GetParam(c, "goodsAheight")
			// 商品A拼接方式
			goodsAspliceType := GetParam(c, "goodsAspliceType")
			// 商品Bid
			goodsBid, _ := strconv.Atoi(GetParam(c, "goodsBid"))
			// 商品B图片URL
			goodsBurl := GetParam(c, "goodsBurl")
			// 商品B宽度
			goodsBwidth := GetParam(c, "goodsBwidth")
			// 商品B高度
			goodsBheight := GetParam(c, "goodsBheight")
			// 商品B拼接方式
			goodsBspliceType := GetParam(c, "goodsBspliceType")

			data, err := insertScenesTask(shopid, userid, scenesmenuid1, scenesmenuid2, goodsAid, goodsBid,
				scenesmenuurl1, scenesmenuurl2, goodsAurl, goodsAwidth, goodsAheight, goodsAspliceType, goodsBurl,
				goodsBwidth, goodsBheight, goodsBspliceType, db)
			Perror(err, "增加模拟场景任务失败")
			return "增加模拟场景任务成功", data
		},

		// 查询模拟场景URL
		"getScenesTaskURL": func(c *gin.Context) (string, interface{}) {
			// 模拟场景任务id
			scenesTaskid, _ := strconv.Atoi(GetParam(c, "scenesTaskid"))

			data, err := getScenesTaskURL(scenesTaskid, db)
			Perror(err, "查询模拟场景URL失败")
			return "查询模拟场景URL成功", data
		},

		// 根据url获取店铺id name
		"getShopId": func(c *gin.Context) (string, interface{}) {
			// 获取 店铺url
			shopurl := GetParam(c, "shopurl")
			UTI, err := getShopID(shopurl, db)
			Perror(err, "查询失败")
			return "查询成功", UTI
		},

		// 用户登录
		"userLogin": func(c *gin.Context) (string, interface{}) {
			// 获取 手机号 密码
			phoneNum := GetParam(c, "phoneNum")
			password := GetParam(c, "password")
			msg, err, data := getUserLogin(phoneNum, password, db)
			Perror(err, msg)
			return msg, data
		},

		// 获取首页信息
		"getIndexInfo": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id
			shopid := GetParam(c, "shopid")
			// 查询店铺 名称 logo 热门商品等信息
			data, err := getIndexInfo(shopid, db)
			Perror(err, "查询失败")
			return "获取店铺联系方式成功", data
		},

		// 店铺联系方式查询
		"getShopContactInfo": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id
			shopid := GetParam(c, "shopid")
			data, err := getShopContactInfo(shopid, db)
			Perror(err, "查询失败")
			return "获取店铺联系方式成功", data
		},

		// 获取全部选择分类
		"getAllClass": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id
			shopid := GetParam(c, "shopid")
			data, err := getAllClass(shopid, db)
			Perror(err, "查询失败")
			return "获取全部分类成功", data
		},

		// 获取二级菜单下的所有商品
		"getGoods": func(c *gin.Context) (string, interface{}) {
			// 获取菜单id
			menuid := GetParam(c, "menuid")
			data, err := getAllGoods(menuid, db)
			Perror(err, "查询失败")
			return "获取商品成功", data
		},

		// 获取商品详细信息
		"getGoodsDetail": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 用户id
			goodsid := GetParam(c, "goodsid")
			userid := GetParam(c, "userid")
			data, err := getGoodsDetail(goodsid, userid, db)
			Perror(err, "查询失败")
			return "获取商品详细信息成功", data
		},

		// 店铺访问记录
		"insertShopAccess": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id 用户id 用户ip
			shopid := GetParam(c, "shopid")
			userid := GetParam(c, "userid")
			ip := GetParam(c, "ip")
			msg, err := insertShopAccess(shopid, userid, ip, db)
			Perror(err, msg)
			return msg, nil
		},

		// 商品访问记录
		"insertGoodsAccess": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id 用户id 用户ip
			goodsid := GetParam(c, "goodsid")
			userid := GetParam(c, "userid")
			ip := GetParam(c, "ip")
			msg, err := insertGoodsAccess(goodsid, userid, ip, db)
			Perror(err, msg)
			return msg, nil
		},
	}
}
