// Package switcher cmd的函数
package switcher

import (
	"database/sql"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ManageDispatch 根据cmd分发到不同的接口函数
func ManageDispatch(db *sql.DB) Xl {
	return Xl{
		// 管理员登陆
		"manageLogin": func(c *gin.Context) (string, interface{}) {
			// 获取 账号 密码
			phoneNum := GetParam(c, "phonenum")
			password := GetParam(c, "password")
			msg, data, err := getManageLogin(phoneNum, password, db)
			Perror(err, msg)
			return msg, data
		},

		// 管理员密码修改
		"updateManagePSWD": func(c *gin.Context) (string, interface{}) {
			// 获取 管理员id 新修改的密码
			manageid := GetParam(c, "manageid")
			password := GetParam(c, "password")
			msg, err := updateManagePSWD(manageid, password, db)
			Perror(err, msg)
			return msg, nil
		},

		// 获取菜单列表
		"getAllClass": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id
			shopid := GetParam(c, "shopid")
			data, err := getAllClass(shopid, db)
			Perror(err, "获取菜单栏失败")
			return "获取菜单栏成功", data
		},

		// 新增店铺产品菜单栏
		"insertMenu": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "shopid")
			// 父菜单id（0为一级菜单）
			superid := GetParam(c, "superid")
			// 菜单名称
			menuname := GetParam(c, "menuname")
			err := insertMenuDao(shopid, superid, menuname, db)
			Perror(err, "新增店铺产品菜单栏失败")
			return "新增店铺产品菜单栏成功", nil
		},

		// 修改店铺产品菜单栏
		"updateMenu": func(c *gin.Context) (string, interface{}) {
			// 获取菜单id
			menuid := GetParam(c, "menuid")
			//新菜单名称
			menuname := GetParam(c, "menuname")
			err := updateMenuDao(menuid, menuname, db)
			Perror(err, "修改店铺产品菜单栏失败")
			return "修改店铺产品菜单栏成功", "修改店铺产品菜单栏成功"
		},

		// 删除店铺产品菜单栏
		"delMenu": func(c *gin.Context) (string, interface{}) {
			// 获取菜单id
			menuid := GetParam(c, "menuid")
			// 获取菜单父id
			superid := GetParam(c, "superid")
			data, err := delMenu(menuid, superid, db)
			Perror(err, "删除菜单失败")
			if data == "1" {
				panic("该菜单下还有二级菜单存在，请检查后再删除")
			}
			if data == "2" {
				panic("该菜单下还有商品存在，请检查后再删除")
			}
			return "删除菜单成功", "删除菜单成功"
		},

		// 获取产品列表
		"getGoodsList": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id  一级菜单id 二级菜单id 产品名称 页数 每页显示条数
			shopid := GetParam(c, "shopid")
			classone := GetParam(c, "classone")
			classtwo := GetParam(c, "classtwo")
			goodsname := GetParam(c, "goodsname")
			page := GetParam(c, "page")
			count := GetParam(c, "count")
			data, err := getGoodsList(shopid, classone, classtwo, goodsname, page, count, db)
			Perror(err, "获取产品列表失败")
			return "获取产品列表成功", data
		},

		// 产品删除
		"deleteGoods": func(c *gin.Context) (string, interface{}) {
			// 获取商品id
			goodsid := GetParam(c, "goodsid")
			err := deleteGoods(goodsid, db)
			Perror(err, "删除商品失败")
			return "删除商品成功", nil
		},

		// 产品下架上架
		"updateGoodsState": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 商品目前状态
			id := GetParam(c, "id")
			state := GetParam(c, "state")
			err := updateGoodsState(id, state, db)
			Perror(err, "更改商品状态失败")
			return "更改商品状态成功", nil
		},

		// 新增产品时获取菜单分类和实物机器位置
		"getMenuAndmachine": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id
			shopid := GetParam(c, "shopid")
			data, err := getMenuAndmachine(shopid, db)
			Perror(err, "获取数据失败")
			return "获取数据成功", data
		},

		// 新增产品
		"insertGoods": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id 菜单id 商品名称 商品品牌 商品风格 商品材质 商品规格 商品产地 商品价格 商品宽度 商品长度 商品关联机器 商品关联机器槽位
			shopid := GetParam(c, "shopid")
			menuid := GetParam(c, "menuid")
			goodsname := GetParam(c, "goodsname")
			brand := GetParam(c, "brand")
			style := GetParam(c, "style")
			material := GetParam(c, "material")
			unit := GetParam(c, "unit")
			madein := GetParam(c, "madein")
			price := GetParam(c, "price")
			width := GetParam(c, "price")
			height := GetParam(c, "price")
			machineid := GetParam(c, "machineid")
			slotnum := GetParam(c, "slotnum")
			data, err := insertManageGoods(shopid, menuid, goodsname, brand, style, material, unit, madein, price,
				width, height, machineid, slotnum, db)
			Perror(err, "新增产品失败")
			return "新增产品成功", data
		},

		// 产品详情查询
		"selectGoodsDetail": func(c *gin.Context) (string, interface{}) {
			// 获取商品id
			goodsid := GetParam(c, "goodsid")
			data, err := selectGoodsDetail(goodsid, db)
			Perror(err, "商品详情查询失败")
			return "商品详情查询成功", data
		},

		// 产品编辑
		"updateGoods": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 菜单id 商品名称 商品品牌 商品风格 商品材质 商品规格 商品产地 商品价格 商品关联机器 商品关联机器槽位
			goodsid := GetParam(c, "goodsid")
			menuid := GetParam(c, "menuid")
			goodsname := GetParam(c, "goodsname")
			brand := GetParam(c, "brand")
			style := GetParam(c, "style")
			material := GetParam(c, "material")
			unit := GetParam(c, "unit")
			madein := GetParam(c, "madein")
			price := GetParam(c, "price")
			width := GetParam(c, "price")
			height := GetParam(c, "price")
			machineid := GetParam(c, "machineid")
			slotnum := GetParam(c, "slotnum")
			err := updateManageGoods(goodsid, menuid, goodsname, brand, style, material, unit, madein, price,
				width, height, machineid, slotnum, db)
			Perror(err, "产品编辑失败")
			return "产品编辑成功", nil
		},

		// 产品颜色查询
		"selectColorList": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 页数 页面总数
			goodsid := GetParam(c, "goodsid")
			page := GetParam(c, "page")
			count := GetParam(c, "count")
			data, err := selectColorList(goodsid, page, count, db)
			Perror(err, "颜色查询失败")
			return "颜色查询成功", data
		},

		// 产品颜色修改
		"updateColor": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 型号 大图 缩略图
			colorid := GetParam(c, "colorid")
			model := GetParam(c, "model")
			pic := GetParam(c, "pic")
			compresspic := GetParam(c, "compresspic")
			err := updateColor(colorid, model, pic, compresspic, db)
			Perror(err, "修改颜色信息失败")
			return "修改颜色信息成功", nil
		},

		// 产品颜色删除
		"deleteColor": func(c *gin.Context) (string, interface{}) {
			// 获取颜色id
			colorid := GetParam(c, "colorid")
			err := deleteColor(colorid, db)
			Perror(err, "删除颜色失败")
			return "删除颜色成功", nil
		},

		// 产品颜色新增
		"insertColor": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 型号 大图 缩略图
			goodsid := GetParam(c, "goodsid")
			model := GetParam(c, "model")
			pic := GetParam(c, "pic")
			compresspic := GetParam(c, "compresspic")
			err := insertColor(goodsid, model, pic, compresspic, db)
			Perror(err, "新增颜色失败")
			return "新增颜色成功", nil
		},

		// 产品颜色设置首选
		"updateColorSetTop": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 商品图片表id
			goodsid := GetParam(c, "goodsid")
			colorid := GetParam(c, "colorid")
			err := updateColorSetTop(goodsid, colorid, db)
			Perror(err, "颜色置顶失败")
			return "颜色置顶成功", nil
		},

		// 产品颜色取消首选
		"updateColorRemoveTop": func(c *gin.Context) (string, interface{}) {
			// 获取商品id 商品图片表id
			colorid := GetParam(c, "colorid")
			err := updateColorRemoveTop(colorid, db)
			Perror(err, "取消颜色置顶失败")
			return "取消颜色置顶成功", nil
		},

		// 热门商品展示配置查询
		"selectHotGoodsShow": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id
			shopid := GetParam(c, "shopid")
			data, err := selectHotGoodsShow(shopid, db)
			Perror(err, "查询热门商品展示配置失败")
			return "查询热门商品展示配置成功", data
		},

		// 热门商品展示配置修改
		"updateHotGoodsShow": func(c *gin.Context) (string, interface{}) {
			// 获取店铺id 爆款点击门限 爆款分类数量
			shopid := GetParam(c, "shopid")
			popularlimit := GetParam(c, "popularlimit")
			popularquantity := GetParam(c, "popularquantity")
			err := updateHotGoodsShow(shopid, popularlimit, popularquantity, db)
			Perror(err, "修改热门商品展示配置失败")
			return "修改热门商品展示配置成功", err
		},

		// 查询店铺列表
		"getShopList": func(c *gin.Context) (string, interface{}) {
			// 要查询的店铺名称
			shopname := GetParam(c, "shopname")
			// 要查询的店铺状态
			state := GetParam(c, "state")
			// 页号
			pageNo, _ := strconv.Atoi(GetParam(c, "pageNo"))
			// 每页数量
			pageSize, _ := strconv.Atoi(GetParam(c, "pageSize"))
			data, err := getShopList(shopname, state, pageNo, pageSize, db)
			Perror(err, "查询店铺列表失败")
			return "查询店铺列表成功", data
		},

		// 店铺URL查重
		"checkShopURL": func(c *gin.Context) (string, interface{}) {
			// 要查重的店铺URL
			shopURL := GetParam(c, "shopurl")

			msg, err := checkShopURLDao(shopURL, db)
			Perror(err, msg)
			return msg, msg
		},

		// 新增店铺
		"insertNewShop": func(c *gin.Context) (string, interface{}) {
			// 店铺url名称
			shopurl := GetParam(c, "shopurl")
			// 店铺名称
			shopname := GetParam(c, "shopname")
			// 店铺图标
			logoimg := GetParam(c, "logoimg")
			// 店铺展示
			shopshow := GetParam(c, "shopshow")
			// 广告图
			bannerimg := GetParam(c, "bannerimg")
			// 老板姓名
			bossname := GetParam(c, "bossname")
			// 座机号码
			telnum := GetParam(c, "telnum")
			// 手机号码
			phonenum := GetParam(c, "phonenum")
			// 微信号
			wechat := GetParam(c, "wechat")
			// 微信号二维码图片
			wechaturl := GetParam(c, "wechaturl")
			// 店铺详细地址
			location := GetParam(c, "location")
			// 店铺经度
			lng := GetParam(c, "lng")
			// 店铺纬度
			lat := GetParam(c, "lat")

			data, err := insertNewShop(shopurl, shopname, logoimg, shopshow, bannerimg, bossname,
				telnum, phonenum, wechat, wechaturl, location, lng, lat, db)
			Perror(err, "新增店铺失败")
			return "新增店铺成功", data
		},

		// 修改店铺信息
		"updateShopInfo": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "shopid")
			// 店铺名称
			shopname := GetParam(c, "shopname")
			// 店铺图标
			logoimg := GetParam(c, "logoimg")
			// 店铺展示
			shopshow := GetParam(c, "shopshow")
			// 老板姓名
			bossname := GetParam(c, "bossname")
			// 座机号码
			telnum := GetParam(c, "telnum")
			// 手机号码
			phonenum := GetParam(c, "phonenum")
			// 微信号
			wechat := GetParam(c, "wechat")
			// 微信号二维码图片
			wechaturl := GetParam(c, "wechaturl")
			// 店铺详细地址
			location := GetParam(c, "location")
			// 店铺经度
			lng := GetParam(c, "lng")
			// 店铺纬度
			lat := GetParam(c, "lat")

			err := updateShopInfoDao(shopid, shopname, logoimg, shopshow, bossname,
				telnum, phonenum, wechat, wechaturl, location, lng, lat, db)
			Perror(err, "修改店铺信息失败")
			return "修改店铺信息成功", nil
		},

		// 修改店铺广告图
		"updateShopBannerImg": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "shopid")
			// 广告图
			bannerimg := GetParam(c, "bannerimg")
			err := updateShopBannerImgDao(shopid, bannerimg, db)
			Perror(err, "修改店铺广告图失败")
			return "修改店铺广告图成功", nil
		},

		// 停用店铺
		"updateShopStatusToDisable": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "shopid")

			err := updateShopStatusToDisableDao(shopid, db)
			Perror(err, "停用店铺失败")
			return "停用店铺成功", nil
		},

		// 删除店铺
		"delShop": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "shopid")

			err := delShopDao(shopid, db)
			Perror(err, "删除店铺失败")
			return "删除店铺成功", nil
		},

		// 查询机器机器列表
		"getMachineList": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "shopid")
			// 要查询的机器序列号
			machineid := GetParam(c, "machineid")
			// 要查询的机器状态
			state := GetParam(c, "state")
			// 页号
			pageNo, _ := strconv.Atoi(GetParam(c, "pageNo"))
			// 每页数量
			pageSize, _ := strconv.Atoi(GetParam(c, "pageSize"))
			data, err := getMachineList(shopid, machineid, state, pageNo, pageSize, db)
			Perror(err, "查询机器机器列表失败")
			return "查询机器机器列表成功", data
		},

		// 新增机器
		"insertNewMachine": func(c *gin.Context) (string, interface{}) {
			// 机器序列号
			machineid := GetParam(c, "machineid")
			// 店铺id
			shopid := GetParam(c, "shopid")
			// 机器槽位数量
			slotnum, _ := strconv.Atoi(GetParam(c, "slotnum"))
			// 机器IP
			machineip := GetParam(c, "machineip")

			err := insertNewMachine(machineid, shopid, machineip, slotnum, db)
			Perror(err, "新增机器失败")
			return "新增机器成功", ""
		},

		// 修改机器ip信息
		"updateMachineInfo": func(c *gin.Context) (string, interface{}) {
			// 机器id
			id := GetParam(c, "id")
			// 机器ip
			machineip := GetParam(c, "machineip")

			err := updateMachineInfoDao(id, machineip, db)
			Perror(err, "修改机器ip信息失败")
			return "修改机器ip信息成功", nil
		},

		// 删除机器
		"delMachine": func(c *gin.Context) (string, interface{}) {
			// 机器id
			id := GetParam(c, "id")

			err := delMachineDao(id, db)
			Perror(err, "删除机器失败")
			return "删除机器成功", nil
		},

		// 查询店铺管理员列表
		"getShopManagerList": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "shopid")
			// 页号
			pageNo, _ := strconv.Atoi(GetParam(c, "pageNo"))
			// 每页数量
			pageSize, _ := strconv.Atoi(GetParam(c, "pageSize"))
			data, err := getShopManagerList(shopid, pageNo, pageSize, db)
			Perror(err, "查询店铺管理员列表失败")
			return "查询店铺管理员列表成功", data
		},

		// 新增店铺管理员
		"insertNewShopManager": func(c *gin.Context) (string, interface{}) {
			// 账号
			phonenum := GetParam(c, "phonenum")
			// 密码
			password := GetParam(c, "password")
			// 备注
			remark := GetParam(c, "remark")
			// 昵称
			nickname := GetParam(c, "nickname")
			// 店铺id
			shopid := GetParam(c, "shopid")

			err := insertNewShopManagerDao(phonenum, password, remark, nickname, shopid, db)
			Perror(err, "新增店铺管理员失败")
			return "新增店铺管理员成功", nil
		},

		// 修改店铺管理员信息
		"updateShopManagerInfo": func(c *gin.Context) (string, interface{}) {
			// 管理员id
			id := GetParam(c, "id")
			// 密码
			password := GetParam(c, "password")
			// 备注
			remark := GetParam(c, "remark")
			// 昵称
			nickname := GetParam(c, "nickname")
			err := updateShopManagerInfoDao(id, password, remark, nickname, db)
			Perror(err, "修改店铺管理员信息失败")
			return "修改店铺管理员信息成功", nil
		},

		// 删除店铺管理员
		"delShopManager": func(c *gin.Context) (string, interface{}) {
			// 管理员id
			id := GetParam(c, "id")

			err := delShopManagerDao(id, db)
			Perror(err, "删除店铺管理员失败")
			return "删除店铺管理员成功", ""
		},

		// 查询模拟场景菜单
		"getScenesMenu": func(c *gin.Context) (string, interface{}) {
			data, err := getScenesMenuList(db)
			Perror(err, "查询模拟场景菜单失败")
			return "查询模拟场景菜单成功", data
		},

		// 新增模拟场景或房间
		"insertScenesMenu": func(c *gin.Context) (string, interface{}) {
			// 父菜单id（0为一级菜单）
			superid := GetParam(c, "superid")
			// 菜单名称
			name := GetParam(c, "name")
			// 菜单图片
			pic := GetParam(c, "pic")
			err := insertScenesMenuDao(superid, name, pic, db)
			Perror(err, "新增模拟场景或房间失败")
			return "新增模拟场景或房间成功", nil
		},

		// 修改模拟场景或房间
		"updateScenesMenu": func(c *gin.Context) (string, interface{}) {
			// 模拟场景菜单id
			id := GetParam(c, "id")
			// 父菜单id（0为一级菜单）
			superid := GetParam(c, "superid")
			// 菜单名称
			name := GetParam(c, "name")
			// 菜单图片
			pic := GetParam(c, "pic")

			err := updateScenesMenuDao(id, superid, name, pic, db)
			Perror(err, "修改模拟场景或房间失败")
			return "修改模拟场景或房间成功", nil
		},

		// 删除模拟场景或房间
		"delScenesMenu": func(c *gin.Context) (string, interface{}) {
			// 模拟场景菜单id
			id := GetParam(c, "id")
			data, err := delScenesMenu(id, db)
			Perror(err, "删除模拟场景或房间失败")
			if data != "删除模拟场景或房间成功" {
				panic(data)
			}
			return data, nil
		},

		// 查询店铺详情
		"getShopInfo": func(c *gin.Context) (string, interface{}) {
			// 店铺id
			shopid := GetParam(c, "id")
			data, err := getShopInfoDao(shopid, db)
			Perror(err, "查询店铺详情失败")
			return "查询店铺详情成功", data
		},
	}
}
