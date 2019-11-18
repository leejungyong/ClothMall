// Package switcher 业务逻辑函数
package switcher

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

// 获取商品列表
func getGoodsList(shopid, classonemenu, classtwomenu, goodsname, page, count string, db *sql.DB) (interface{}, error) {
	var MGL ManageGoodsList
	selectSQL := "select x.id, x.goodsname, x.menuname as classtwoname, m.menuname as classonename, ifnull(y.clicknum, 0) as clicknum, x.state from (select g.id, g.shopid, g.menuid, g.goodsname, g.state, m.superid, m.menuname from goods g left join menu m on m.id = g.menuid) x left join menu m on x.superid = m.id left join (select goodsid, count(*) as clicknum from goodsaccess group by goodsid) y on x.id = y.goodsid where x.id > 0 and x.state <> 2"
	if goodsname != "" {
		selectSQL += ` and x.goodsname like "%` + goodsname + `%"`
	}
	if classonemenu != "" {
		selectSQL += ` and m.id = ` + classonemenu
	}
	if classtwomenu != "" {
		selectSQL += ` and x.menuid = ` + classtwomenu
	}
	// 获取limit后起始数量
	pageInt := StringToInt(page) - 1
	// 默认每页10条
	countInt := StringToInt(count)
	if count == "" {
		countInt = 10
		count = "10"
	}
	begin := IntToString(pageInt * countInt)
	selectSQL = selectSQL + " limit " + begin + "," + count
	log.Println(selectSQL)
	MGLPS, err := selectGoodsList(selectSQL, db)
	if err != nil {
		log.Println(err)
		return MGL, err
	}
	// 获取需要查询的菜单id
	selectSQL2 := "select count(*) from goods where shopid = " + shopid + " and state <> 2"
	splicingSQL := " and menuid = " + classtwomenu
	// 如果二级菜单未选 获取所有一级菜单下的二级菜单
	if classtwomenu == "" {
		arr, err := getClassTwoMenuId(classonemenu, db)
		if err != nil {
			log.Println(err)
			return MGL, err
		}
		var arrStr string
		for _, k := range arr {
			if arrStr == "" {
				arrStr = IntToString(k)
			} else {
				arrStr += "," + IntToString(k)
			}
		}
		splicingSQL = " and menuid in (" + arrStr + ")"
	}
	selectSQL2 += splicingSQL
	total, err := selectGoodsNum(selectSQL2, db)
	if err != nil {
		log.Println(err)
		return MGL, err
	}
	MGL.ManageGoodsListParam = MGLPS
	MGL.Total = total
	return MGL, err
}

// 商品上|下架
func updateGoodsState(id, stateNow string, db *sql.DB) error {
	if stateNow == "1" {
		return updateGoodsShelfState(0, id, db)
	}
	return updateGoodsShelfState(1, id, db)
}

// 获取菜单和机器信息
func getMenuAndmachine(shopid string, db *sql.DB) (interface{}, error) {
	var MMAM ManageMenuAndMachine
	// 获取店铺菜单信息
	var MMLS []ManageMenuList
	// 获取所有一级分类
	ACOS, err := getAllClassifyOne(shopid, db)
	if nil != err {
		log.Println(err)
		return MMAM, err
	}
	for _, v := range ACOS {
		var MML ManageMenuList
		// 获取每一个一级分类下的二级分类
		MML.ID = v.ID
		MML.MenuName = v.MenuName
		// 获取二级分类
		CLS, err := getAllClassifyTwo(v.ID, db)
		if err != nil {
			return MMAM, err
		}
		MML.ClassList = CLS
		MMLS = append(MMLS, MML)
	}
	// 获取机器槽位详情
	MMLS1, err := getMachine(shopid, db)
	if err != nil {
		return MMAM, err
	}
	for i, k := range MMLS1 {
		log.Println(k.ID)
		MSLS, err := getMachineSlot(k.ID, db)
		if err != nil {
			return MMAM, err
		}
		MMLS1[i].SlotList = MSLS
	}
	MMAM.ManageMenuList = MMLS
	MMAM.ManageMachineList = MMLS1
	// 获取机器槽位
	return MMAM, nil
}

// 新增产品
func insertManageGoods(shopid, menuid, goodsname, brand, style, material, unit, madein, price, width, height, machineid, slotnum string, db *sql.DB) (interface{}, error) {
	var MBGI ManageBackGoodsId
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return MBGI, err
	}
	rowID, err := insertGoods(shopid, menuid, goodsname, brand, style, material, unit, madein, price, width, height, tx)
	if nil != err {
		log.Println(err)
		return MBGI, err
	}
	err = updateMachineSlot(Int64ToString(rowID), machineid, slotnum, tx)
	if err != nil {
		log.Println(err)
		return MBGI, err
	}
	tx.Commit()
	MBGI.ID = int(rowID)
	return MBGI, err
}

// 修改产品
func updateManageGoods(goodsid, menuid, goodsname, brand, style, material, unit, madein, price, width, height, machineid, slotnum string, db *sql.DB) error {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// 修改goods表
	err = updateGoods(goodsid, menuid, goodsname, brand, style, material, unit, madein, price, width, height, tx)
	if nil != err {
		log.Println(err)
		return err
	}
	// 将商品原来机器槽位删除
	err = deleteMachineSlot(goodsid, tx)
	if err != nil {
		log.Println(err)
		return err
	}
	// 修改机器槽位表
	err = updateMachineSlot(goodsid, machineid, slotnum, tx)
	if err != nil {
		log.Println(err)
		return err
	}
	tx.Commit()
	return err
}

// 查询颜色列表
func selectColorList(goodsid, page, count string, db *sql.DB) (interface{}, error) {
	var MCL ManageColorList
	selectSQL := "select id, goodsid, model, compresspic, isfirst from goodspic where goodsid = ? and state = 1 order by isfirst desc"
	// 获取limit后起始数量
	pageInt := StringToInt(page) - 1
	// 默认每页10条
	countInt := StringToInt(count)
	if count == "" {
		countInt = 10
		count = "10"
	}
	begin := IntToString(pageInt * countInt)
	selectSQL = selectSQL + " limit " + begin + "," + count
	CLS, err := selectColor(selectSQL, goodsid, db)
	if err != nil {
		log.Println(err)
		return MCL, err
	}
	// 获取颜色总数
	total, err := selectColorCount(goodsid, db)
	if err != nil {
		log.Println(err)
		return MCL, err
	}
	MCL.ColorList = CLS
	MCL.Total = total
	return MCL, err
}

// 颜色置顶
func updateColorSetTop(goodsid, colorid string, db *sql.DB) error {
	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	// 把该商品的所有颜色isfirst设为0
	err = updateColorZeroing(goodsid, tx)
	if err != nil {
		return err
	}
	// 修改需要置顶的颜色isfirst
	err = updateTopColor(colorid, tx)
	if err != nil {
		return err
	}
	tx.Commit()
	return err
}

// 查询店铺列表
func getShopList(shopname, state string, pageNo, pageSize int, db *sql.DB) (*ListPageData, error) {
	var data = new(ListPageData)
	var shopList []*ShopList
	var cnt int
	var err error
	// 拼sql
	shopCntSQL, shopListSQL := getShopListSQL(shopname, state, pageNo, pageSize)
	log.Println(shopCntSQL)
	log.Println(shopListSQL)

	// 获取店铺数据条数
	cnt, err = getDataCnt(shopCntSQL, db)
	if nil != err {
		log.Println(err)
		return data, err
	}
	// 获取店铺信息
	shopList, err = getShopListDao(shopListSQL, db)
	if nil != err {
		log.Println(err)
		return data, err
	}

	data.PageTotal = cnt
	data.ListCount = cnt
	data.List = shopList

	return data, err
}

// 根据搜索条件生成拼店铺列表SQL
func getShopListSQL(shopname, state string, pageNo, pageSize int) (string, string) {
	// 查询店铺总数量的sql
	sql1 := `select count(*) as total
	from shop s
	left join (select shopid, count(*) as goodsnum from goods where state = 1 group by shopid) x on x.shopid = s.id
	left join (select shopid, count(*) as visitnum from shopaccess group by shopid) y on y.shopid = s.id
	left join (select shopid, count(*) as managenum from manager where state = 1 group by shopid) z on z.shopid = s.id
	left join (select shopid, count(*) as machinenum from machine where state = 1 group by shopid) w on w.shopid = s.id
	where s.state <> 4 %condition%`

	// 查询店铺列表的sql
	sql2 := `select s.id, s.shopurl, s.state, s.shopname, s.bossname, s.phonenum, s.telnum, s.wechat, ifnull(x.goodsnum, 0) as goodsnum, 
	ifnull(y.visitnum, 0) as visitnum, ifnull(z.managenum, 0) as managenum, ifnull(w.machinenum, 0) as machinenum
	from shop s
	left join (select shopid, count(*) as goodsnum from goods where state = 1 group by shopid) x on x.shopid = s.id
	left join (select shopid, count(*) as visitnum from shopaccess group by shopid) y on y.shopid = s.id
	left join (select shopid, count(*) as managenum from manager where state = 1 group by shopid) z on z.shopid = s.id
	left join (select shopid, count(*) as machinenum from machine where state = 1 group by shopid) w on w.shopid = s.id
	where s.state <> 4 %condition% %paging%`
	// 拼查询条件
	var condition string
	if "" != shopname {
		condition += " and s.shopname like '%" + shopname + "%'"
	}
	if "" != state {
		condition += " and s.state = " + state
	}
	sql2 = strings.Replace(sql2, "%condition%", condition, -1)
	// 分页条件
	paging := "limit " + strconv.Itoa((pageNo-1)*pageSize) + ", " + strconv.Itoa(pageSize)
	return strings.Replace(sql1, "%condition%", condition, 1), strings.Replace(sql2, "%paging%", paging, 1)
}

// 新增店铺
func insertNewShop(shopurl, shopname, logoimg, shopshow, bannerimg, bossname, telnum,
	phonenum, wechat, wechaturl, location, lng, lat string, db *sql.DB) (int64, error) {
	var shopid int64
	var err error
	// 新增店铺
	shopid, err = insertNewShopDao(shopurl, shopname, logoimg, shopshow, bannerimg, bossname,
		telnum, phonenum, wechat, wechaturl, location, lng, lat, db)
	if nil != err {
		log.Println(err)
		return shopid, err
	}
	// 根据店铺ID生成默认菜单栏
	err = insertDefaultMenu(shopid, db)
	if nil != err {
		log.Println(err)
		return shopid, err
	}
	return shopid, err
}

// 根据店铺ID生成默认菜单栏
func insertDefaultMenu(shopid int64, db *sql.DB) error {
	superMenu := [3]string{"墙纸类", "墙布类", "其他类"}
	childMenu := [10]string{"中式风格", "欧式风格", "美式风格", "现代风格", "田园风格", "卡通风格", "素色风格", "复古风格", "壁画风格", "地中海风格"}
	otherMenu := [5]string{"墙纸墙布辅料", "窗帘", "软包", "硬包定制", "施工服务"}
	var superMenuID int64
	// 开启事务
	tx, err := db.Begin()

	for _, s := range superMenu {
		// 循环增加一级菜单
		superMenuID, err = insertDefaultSuperMenuDao(shopid, s, tx)
		if nil != err {
			log.Println(err)
		}
		if s != "其他类" {
			for _, c := range childMenu {
				// 循环增加二级菜单
				err = insertDefaultChildMenuDao(shopid, superMenuID, c, tx)
			}
		} else {
			// 循环增加二级菜单
			for _, c := range otherMenu {
				err = insertDefaultChildMenuDao(shopid, superMenuID, c, tx)
			}
		}
	}
	tx.Commit()
	if nil != err {
		log.Println(err)
		return err
	}
	return err
}

// 查询机器列表
func getMachineList(shopid, machineid, state string, pageNo, pageSize int, db *sql.DB) (*ListPageData, error) {
	var data = new(ListPageData)
	var machine []*Machine
	var cnt int
	var err error
	// 拼sql
	machineCntSQL, machineListSQL := getMachineListSQL(shopid, machineid, state, pageNo, pageSize)
	log.Println(machineCntSQL)
	log.Println(machineListSQL)

	// 获取机器数据条数
	cnt, err = getDataCnt(machineCntSQL, db)
	if nil != err {
		log.Println(err)
		return data, err
	}
	// 获取机器信息
	machine, err = getMachineListDao(machineListSQL, db)
	if nil != err {
		log.Println(err)
		return data, err
	}

	data.PageTotal = cnt
	data.ListCount = cnt
	data.List = machine

	return data, err
}

// 根据搜索条件生成拼机器列表SQL
func getMachineListSQL(shopid, machineid, state string, pageNo, pageSize int) (string, string) {
	// 查询店铺总数量的sql
	sql1 := `select count(*) as total
	from machine
	where state <> 2 %condition%`

	// 查询店铺列表的sql
	sql2 := `select id,machineid,shopid,slotnum,runstate,netstate,machineip,state 
	from machine
	where state <> 2 %condition% %paging%`
	// 拼查询条件
	var condition string
	if "" != machineid {
		condition += " and machineid like '%" + machineid + "%'"
	}
	if "" != state {
		condition += " and state = " + state
	}
	condition += " and shopid = " + shopid
	sql2 = strings.Replace(sql2, "%condition%", condition, -1)
	// 分页条件
	paging := "limit " + strconv.Itoa((pageNo-1)*pageSize) + ", " + strconv.Itoa(pageSize)
	return strings.Replace(sql1, "%condition%", condition, 1), strings.Replace(sql2, "%paging%", paging, 1)
}

// 新增机器
func insertNewMachine(machineid, shopid, machineip string, slotnum int, db *sql.DB) error {
	var machineID int64
	var err error
	// 开启事务
	tx, err := db.Begin()

	// 新增机器
	machineID, err = insertNewMachineDao(machineid, shopid, machineip, slotnum, tx)
	if nil != err {
		log.Println(err)
		return err
	}

	// 根据机器槽位数量，批量新增机器槽位数
	for i := 1; i <= slotnum; i++ {
		err = insertNewMachineSlotDao(Int64ToString(machineID), shopid, i, tx)
		if nil != err {
			log.Println(err)
		}
	}
	tx.Commit()
	if nil != err {
		log.Println(err)
		return err
	}

	return err
}

// 查询店铺管理员列表
func getShopManagerList(shopid string, pageNo, pageSize int, db *sql.DB) (*ListPageData, error) {
	var data = new(ListPageData)
	var manager []*Manager
	var cnt int
	var err error
	// 拼sql
	managerCntSQL, managerListSQL := getShopManagerListSQL(shopid, pageNo, pageSize)
	log.Println(managerCntSQL)
	log.Println(managerListSQL)

	// 获取店铺管理员数据条数
	cnt, err = getDataCnt(managerCntSQL, db)
	if nil != err {
		log.Println(err)
		return data, err
	}
	// 获取店铺管理员信息
	manager, err = getShopManagerListDao(managerListSQL, db)
	if nil != err {
		log.Println(err)
		return data, err
	}

	data.PageTotal = cnt
	data.ListCount = cnt
	data.List = manager

	return data, err
}

// 根据搜索条件生成拼店铺管理员列表SQL
func getShopManagerListSQL(shopid string, pageNo, pageSize int) (string, string) {
	// 查询店铺总数量的sql
	sql1 := `select count(*) as total
	from manager
	where state <> 2 %condition%`

	// 查询店铺列表的sql
	sql2 := `select id,phonenum,password,remark,shopid,state 
	from manager
	where state <> 2 %condition% %paging%`
	// 拼查询条件
	var condition string
	condition += " and shopid = " + shopid
	sql2 = strings.Replace(sql2, "%condition%", condition, -1)
	// 分页条件
	paging := "limit " + strconv.Itoa((pageNo-1)*pageSize) + ", " + strconv.Itoa(pageSize)
	return strings.Replace(sql1, "%condition%", condition, 1), strings.Replace(sql2, "%paging%", paging, 1)
}

// 查询模拟场景房间列表
func getScenesMenuList(db *sql.DB) (*ListPageData, error) {
	var data = new(ListPageData)
	var superMenu []*ScenesMenu
	var childMenu []*ScenesMenu
	var cnt int
	var err error
	// 拼sql
	scenesMenuCntSQL, getScenesMenuList := getScenesMenuListSQL()
	log.Println(scenesMenuCntSQL)
	log.Println(getScenesMenuList)

	// 获取模拟场景房间数据条数
	cnt, err = getDataCnt(scenesMenuCntSQL, db)
	if nil != err {
		log.Println(err)
		return data, err
	}
	// 获取模拟场景房间信息
	superMenu, err = getScenesMenuListDao(getScenesMenuList, db)
	if nil != err {
		log.Println(err)
		return data, err
	}

	// 开启事务
	tx, err := db.Begin()
	// 查询二级房间菜单
	for _, s := range superMenu {
		// 循环增加一级菜单
		childMenu, err = getScenesMenuChildListDao(s.ID, tx)
		if nil != err {
			log.Println(err)
		}
		s.Children = childMenu
	}
	tx.Commit()
	if nil != err {
		log.Println(err)
		return data, err
	}
	data.PageTotal = cnt
	data.ListCount = cnt
	data.List = superMenu

	return data, err
}

// 根据搜索条件生成模拟场景菜单列表SQL
func getScenesMenuListSQL() (string, string) {
	// 查询模拟场景菜单总数量的sql
	sql1 := `select count(*) as total from scenesmenu s 
	where s.state = 1 and s.superid = 0`

	// 查询模拟场景菜单列表的sql
	sql2 := `select id,name,superid,pic from scenesmenu s 
	where s.state = 1 and s.superid = 0`

	return sql1, sql2
}

// 删除菜单
func delMenu(menuid, superid string, db *sql.DB) (string, error) {
	var num int
	var err error
	// 检查菜单是否可以删除
	if superid == "0" {
		//如果是一级菜单栏，检查是否还有二级菜单
		num, err = get2levelByMenuIDDao(menuid, db)
		if nil != err {
			log.Println(err)
			return "0", err
		}
		if num != 0 {
			return "1", err
		}
	} else {
		//如果是二级菜单栏，检查是否还有商品存在
		num, err = getGoodsByMenuIDDao(menuid, db)
		if nil != err {
			log.Println(err)
			return "0", err
		}
		if num != 0 {
			return "2", err
		}
	}
	// 执行删除菜单操作
	err = delMenuDao(menuid, db)
	if nil != err {
		log.Println(err)
		return "0", err
	}
	return "99", err
}

// 删除模拟场景或房间
func delScenesMenu(id string, db *sql.DB) (string, error) {
	var err error
	var cnt int
	// 查询判断是否还存在子菜单
	shopCntSQL := "select count(*) as cnt from scenesmenu where superid = " + id
	cnt, err = getDataCnt(shopCntSQL, db)
	if nil != err {
		log.Println(err)
		return "删除模拟场景或房间失败", err
	}
	if cnt == 0 {
		// 没有子菜单，删除模拟场景或房间
		err = delScenesMenuDao(id, db)
		if nil != err {
			log.Println(err)
			return "删除模拟场景或房间失败", err
		}
	} else {
		return "删除模拟场景或房间失败，该场景下还有房间存在", err
	}

	return "删除模拟场景或房间成功", err
}
