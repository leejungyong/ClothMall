// Package switcher 数据库交互函数
package switcher

import (
	"database/sql"
	"errors"
	"log"
)

// 管理员登陆
func getManageLogin(phoneNum, password string, db *sql.DB) (string, interface{}, error) {
	var ML ManageLogin
	selectSQL := "select ifnull(id, 0), ifnull(shopid, 0), ifnull(count(*), 0) as num, ifnull(nickname,'') from manager where phonenum = ? and password = ? and state = 1"
	var num int
	err := db.QueryRow(selectSQL, phoneNum, password).Scan(&ML.ManageID, &ML.ShopID, &num, &ML.NickName)
	if err != nil {
		return "查询失败", ML, err
	}
	// 没有该账号信息
	if num == 0 {
		return "账号或密码错误", ML, errors.New("账号或密码错误")
	}
	return "登陆成功", ML, err
}

// 管理员账号密码修改
func updateManagePSWD(manageid, password string, db *sql.DB) (string, error) {
	updateSQL := "update manager set password = ? where id = ?"
	_, err := db.Exec(updateSQL, password, manageid)
	if err != nil {
		return "查询失败", err
	}
	return "修改密码成功", err
}

// 获取数据条数
func getDataCnt(sqls string, db *sql.DB) (int, error) {
	var cnt int
	err := db.QueryRow(sqls).Scan(&cnt)
	return cnt, err
}

// 查询店铺列表
func getShopListDao(sqlstr string, db *sql.DB) ([]*ShopList, error) {
	var list []*ShopList
	rows, err := db.Query(sqlstr)
	if nil != err {
		log.Println(err)
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var s ShopList
		rows.Scan(&s.ShopID, &s.ShopURL, &s.ShopName, &s.BossName, &s.TelNum, &s.PhoneNum, &s.Wechat, &s.GoodsNum, &s.ManageNum, &s.MachineNum, &s.VisitNum, &s.State)
		list = append(list, &s)
	}
	return list, err
}

// 店铺URL查重
func checkShopURLDao(shopURL string, db *sql.DB) (string, error) {
	var num int
	selectSQL := "select ifnull(count(*), 0) as num from shop where shopurl = ?"
	err := db.QueryRow(selectSQL, shopURL).Scan(&num)
	if err != nil {
		return "查询失败", err
	}
	// 店铺URL校验
	if num != 0 {
		return "该店铺URL已经存在", errors.New("该店铺URL已经存在")
	}
	return "该店铺URL可以使用", err
}

// 新增店铺
func insertNewShopDao(shopurl, shopname, logoimg, shopshow, bannerimg, bossname, telnum,
	phonenum, wechat, wechaturl, location, lng, lat string, db *sql.DB) (int64, error) {
	insertSQL := `insert into shop (shopurl, shopname, logoimg, shopshow, bannerimg, bossname,telnum, phonenum, wechat, wechaturl, 
		popularlimit, popularquantity, location, lng, lat, state) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, '2','100', ?, ?, ?, '1')`
	var rowid int64
	result, err := db.Exec(insertSQL, shopurl, shopname, logoimg, shopshow, bannerimg, bossname, telnum,
		phonenum, wechat, wechaturl, location, lng, lat)
	if nil != err {
		log.Println(err)
		return rowid, err
	}
	rowid, err = result.LastInsertId()
	return rowid, nil
}

// 新增默认一级菜单
func insertDefaultSuperMenuDao(shopid int64, shopname string, tx *sql.Tx) (int64, error) {
	insertSQL := `insert into menu(shopid,superid,menuname,state) values (?, 0, ?, 1)`
	var rowid int64
	result, err := tx.Exec(insertSQL, shopid, shopname)
	if nil != err {
		log.Println(err)
		return rowid, err
	}
	rowid, err = result.LastInsertId()
	return rowid, nil
}

// 新增默认二级菜单
func insertDefaultChildMenuDao(shopid, superMenuID int64, shopname string, tx *sql.Tx) error {
	insertSQL := `insert into menu(shopid,superid,menuname,state) values (?, ?, ?, 1)`
	_, err := tx.Exec(insertSQL, shopid, superMenuID, shopname)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 修改店铺信息
func updateShopInfoDao(shopid, shopname, logoimg, shopshow, bossname, telnum,
	phonenum, wechat, wechaturl, location, lng, lat string, db *sql.DB) error {
	updateSQL := `UPDATE shop SET shopname = ?, logoimg = ?, shopshow = ?, bossname = ?,telnum = ?, 
	phonenum = ?, wechat = ?, wechaturl = ?, location = ?, lng = ?, lat = ? where id  = ?`
	_, err := db.Exec(updateSQL, shopname, logoimg, shopshow, bossname, telnum,
		phonenum, wechat, wechaturl, location, lng, lat, shopid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 修改店铺广告图
func updateShopBannerImgDao(shopid, bannerimg string, db *sql.DB) error {
	updateSQL := `UPDATE shop SET bannerimg = ? where id  = ?`
	_, err := db.Exec(updateSQL, bannerimg, shopid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 停用店铺
func updateShopStatusToDisableDao(shopid string, db *sql.DB) error {
	updateSQL := `UPDATE shop SET state = 3 where id  = ?`
	_, err := db.Exec(updateSQL, shopid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 删除店铺
func delShopDao(shopid string, db *sql.DB) error {
	updateSQL := `UPDATE shop SET state = 4 where id  = ?`
	_, err := db.Exec(updateSQL, shopid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 查询机器列表
func getMachineListDao(sqlstr string, db *sql.DB) ([]*Machine, error) {
	var list []*Machine
	rows, err := db.Query(sqlstr)
	if nil != err {
		log.Println(err)
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var m Machine
		rows.Scan(&m.ID, &m.MachineID, &m.ShopID, &m.SlotNum, &m.RunState, &m.NetState, &m.MachineIP, &m.State)
		list = append(list, &m)
	}
	return list, err
}

// 新增机器
func insertNewMachineDao(machineid, shopid, machineip string, slotnum int, tx *sql.Tx) (int64, error) {
	var rowid int64
	insertSQL := `insert into machine (machineid, shopid, slotnum, machineip) values (?, ?, ?, ?)`
	result, err := tx.Exec(insertSQL, machineid, shopid, slotnum, machineip)
	if nil != err {
		log.Println(err)
		return rowid, err
	}
	rowid, err = result.LastInsertId()
	return rowid, err
}

// 新增机器槽位
func insertNewMachineSlotDao(machineid, shopid string, slotnum int, tx *sql.Tx) error {
	insertSQL := `insert into machineslot (machineid, shopid, slotnum) values (?, ?, ?)`
	_, err := tx.Exec(insertSQL, machineid, shopid, slotnum)
	if nil != err {
		log.Println(err)
		return err
	}
	return err
}

// 修改机器ip信息
func updateMachineInfoDao(id, machineip string, db *sql.DB) error {
	updateSQL := `UPDATE machine SET machineip = ? where id  = ?`
	_, err := db.Exec(updateSQL, machineip, id)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 删除机器
func delMachineDao(id string, db *sql.DB) error {
	updateSQL := `UPDATE machine SET state = 2 where id  = ?`
	_, err := db.Exec(updateSQL, id)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 查询店铺管理员列表
func getShopManagerListDao(sqlstr string, db *sql.DB) ([]*Manager, error) {
	var list []*Manager
	rows, err := db.Query(sqlstr)
	if nil != err {
		log.Println(err)
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var m Manager
		rows.Scan(&m.ID, &m.PhoneNum, &m.Password, &m.Remark, &m.ShopID, &m.State)
		list = append(list, &m)
	}
	return list, err
}

// 新增店铺管理员
func insertNewShopManagerDao(phonenum, password, remark, nickname, shopid string, db *sql.DB) error {
	insertSQL := `insert into manager (phonenum, password, remark, nickname, shopid) values (?, ?, ?, ?, ?)`
	_, err := db.Exec(insertSQL, phonenum, password, remark, nickname, shopid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 修改店铺管理员信息
func updateShopManagerInfoDao(id, password, remark, nickname string, db *sql.DB) error {
	updateSQL := `UPDATE manager SET password = ?, remark = ?, nickname=? where id  = ?`
	_, err := db.Exec(updateSQL, password, remark, nickname, id)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 删除店铺管理员
func delShopManagerDao(id string, db *sql.DB) error {
	updateSQL := `UPDATE manager SET state = 2 where id  = ?`
	_, err := db.Exec(updateSQL, id)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 查询模拟场景场景列表
func getScenesMenuListDao(sqlstr string, db *sql.DB) ([]*ScenesMenu, error) {
	var list []*ScenesMenu
	rows, err := db.Query(sqlstr)
	if nil != err {
		log.Println(err)
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var s ScenesMenu
		rows.Scan(&s.ID, &s.Name, &s.SuperID, &s.Pic)
		list = append(list, &s)
	}
	return list, err
}

// 查询模拟场景房间列表
func getScenesMenuChildListDao(superid int, tx *sql.Tx) ([]*ScenesMenu, error) {
	var list []*ScenesMenu
	selectSQL := "select id,name,superid,pic from scenesmenu where state = 1 and superid = ?"
	rows, err := tx.Query(selectSQL, superid)
	if nil != err {
		log.Println(err)
		return list, err
	}
	defer rows.Close()
	for rows.Next() {
		var s ScenesMenu
		rows.Scan(&s.ID, &s.Name, &s.SuperID, &s.Pic)
		list = append(list, &s)
	}
	return list, err
}

// 新增模拟场景或房间
func insertScenesMenuDao(superid, name, pic string, db *sql.DB) error {
	insertSQL := `insert into scenesmenu (name, superid, pic) values (?, ?, ?)`
	_, err := db.Exec(insertSQL, name, superid, pic)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 修改模拟场景或房间信息
func updateScenesMenuDao(id, superid, name, pic string, db *sql.DB) error {
	updateSQL := `UPDATE scenesmenu SET superid = ?,name = ?,pic = ? where id  = ?`
	_, err := db.Exec(updateSQL, superid, name, pic, id)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 删除模拟场景或房间
func delScenesMenuDao(id string, db *sql.DB) error {
	updateSQL := `UPDATE scenesmenu SET state = 2 where id  = ?`
	_, err := db.Exec(updateSQL, id)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 新增店铺菜单栏
func insertMenuDao(shopid, superid, menuname string, db *sql.DB) error {
	insertSQL := `insert into menu (shopid, superid, menuname) values (?, ?, ?)`
	_, err := db.Exec(insertSQL, shopid, superid, menuname)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 修改店铺菜单栏名称
func updateMenuDao(menuid, menuname string, db *sql.DB) error {
	updateSQL := "update menu set menuname = ? where id = ?"
	_, err := db.Exec(updateSQL, menuname, menuid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 删除店铺菜单栏
func delMenuDao(menuid string, db *sql.DB) error {
	updateSQL := `UPDATE menu SET state = 2 where id  = ?`
	_, err := db.Exec(updateSQL, menuid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 查找商品列表
func selectGoodsList(sql string, db *sql.DB) ([]ManageGoodsListParam, error) {
	var MGLPS []ManageGoodsListParam
	rows, err := db.Query(sql)
	if err != nil {
		log.Println(err)
		return MGLPS, err
	}
	defer rows.Close()
	for rows.Next() {
		var MGLP ManageGoodsListParam
		err := rows.Scan(&MGLP.ID, &MGLP.GoodsName, &MGLP.ClassOne, &MGLP.ClassTwo, &MGLP.ClickNum, &MGLP.State)
		if err != nil {
			log.Println(err)
			return MGLPS, err
		}
		MGLPS = append(MGLPS, MGLP)
	}
	return MGLPS, err
}

// 获取指定一级菜单下的所有二级菜单
func getClassTwoMenuId(menuid string, db *sql.DB) ([]int, error) {
	var arr []int
	selectSQL := "select id from menu where superid = " + menuid
	rows, err := db.Query(selectSQL)
	if err != nil {
		log.Println(err)
		return arr, err
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return arr, err
		}
		arr = append(arr, id)
	}
	return arr, err
}

// 查找商品数量
func selectGoodsNum(sql string, db *sql.DB) (int, error) {
	var total int
	selectSQL := sql
	err := db.QueryRow(selectSQL).Scan(&total)
	if err != nil {
		log.Println(err)
		return total, err
	}
	return total, err
}

// 删除商品列表
func deleteGoods(id string, db *sql.DB) error {
	updateSQL := "update goods set state = 2 where id = ?"
	_, err := db.Exec(updateSQL, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 商品上|下架 state==1上架 state==0下架
func updateGoodsShelfState(state int, id string, db *sql.DB) error {
	updateSQL := "update goods set state = ? where id = ?"
	_, err := db.Exec(updateSQL, state, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 获取机器详情
func getMachine(shopid string, db *sql.DB) ([]ManageMachineList, error) {
	var MMLS []ManageMachineList
	selectSQL := "select id, name from machine where state = 1 and shopid = ?"
	rows, err := db.Query(selectSQL, shopid)
	if err != nil {
		log.Println(err)
		return MMLS, err
	}
	defer rows.Close()
	for rows.Next() {
		var MML ManageMachineList
		err := rows.Scan(&MML.ID, &MML.Name)
		if err != nil {
			log.Println(err)
			return MMLS, err
		}
		MMLS = append(MMLS, MML)
	}
	return MMLS, err
}

// 获取机器子槽位
func getMachineSlot(machineid int, db *sql.DB) ([]ManageSlotList, error) {
	var MSLS []ManageSlotList
	selectSQL := "select id, slotnum from machineslot where state = 1 and goodsid = 0 and machineid = ?"
	rows, err := db.Query(selectSQL, machineid)
	if err != nil {
		log.Println(err)
		return MSLS, err
	}
	defer rows.Close()
	for rows.Next() {
		var MSL ManageSlotList
		err := rows.Scan(&MSL.ID, &MSL.Num)
		if err != nil {
			log.Println(err)
			return MSLS, err
		}
		MSLS = append(MSLS, MSL)
	}
	return MSLS, err
}

// 新增产品
func insertGoods(shopid, menuid, goodsname, brand, style, material, unit, madein, price, width, height string, tx *sql.Tx) (int64, error) {
	var rowid int64
	insertSQL := "insert into goods (shopid, menuid, goodsname, brand, style, material, unit, madein, price, width, height) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,?)"
	res, err := tx.Exec(insertSQL, shopid, menuid, goodsname, brand, style, material, unit, madein, price, width, height)
	if err != nil {
		log.Println(err)
		return rowid, err
	}
	rowid, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
	}
	return rowid, err
}

// 修改产品
func updateGoods(goodsid, menuid, goodsname, brand, style, material, unit, madein, price, width, height string, tx *sql.Tx) error {
	updateSQL := "update goods set menuid = ?, goodsname = ?, brand = ?, style = ?, material = ?, unit = ?, madein = ?, price = ?, width = ?, height = ? where id = ?"
	_, err := tx.Exec(updateSQL, menuid, goodsname, brand, style, material, unit, madein, price, width, height, goodsid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 查看产品详情
func selectGoodsDetail(goodsid string, db *sql.DB) (interface{}, error) {
	var MGD ManageGoodsDetail
	selectSQL := "select g.id, g.shopid, g.menuid, g.goodsname, g.brand, g.style, g.material, g.unit, g.madein, g.price, g.width, g.height, g.machineid, g.machineadder, m.superid from goods g left join menu m on g.menuid = m.id where g.id = ?"
	err := db.QueryRow(selectSQL, goodsid).Scan(&MGD.ID, &MGD.ShopId, &MGD.ClassTwoId, &MGD.GoodsName, &MGD.Brand, &MGD.Style, &MGD.Material, &MGD.Unit, &MGD.Madein,
		&MGD.Price, &MGD.Width, &MGD.Height, &MGD.MachineId, &MGD.MachineAdder, &MGD.ClassOneId)
	if err != nil {
		log.Println(err)
	}
	return MGD, err
}

// 修改机器槽位状态
func updateMachineSlot(goodsid, machineid, slotnum string, tx *sql.Tx) error {
	updateSQL := "update machineslot set goodsid = ? where machineid = ? and slotnum = ?"
	_, err := tx.Exec(updateSQL, goodsid, machineid, slotnum)
	if nil != err {
		log.Println(err)
	}
	return err
}

// 删除原先机器槽状态
func deleteMachineSlot(goodsid string, tx *sql.Tx) error {
	updateSQL := "update machineslot set goodsid = 0 where goodsid = ?"
	_, err := tx.Exec(updateSQL, goodsid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 修改颜色
func updateColor(colorid, model, pic, compresspic string, db *sql.DB) error {
	updateSQL := `update goodspic set model = ?, pic = ?, compresspic = ? where id = ?`
	_, err := db.Exec(updateSQL, model, pic, compresspic, colorid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 新增颜色
func insertColor(goodsid, model, pic, compresspic string, db *sql.DB) error {
	insertSQL := "insert into goodspic (goodsid, model, pic, compresspic) values (?, ?, ?, ?)"
	_, err := db.Exec(insertSQL, goodsid, model, pic, compresspic)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 获取颜色数量
func selectColorCount(goodsid string, db *sql.DB) (int, error) {
	var num int
	selectSQL := "select count(*) from goodspic where goodsid = ? and state = 1"
	err := db.QueryRow(selectSQL, goodsid).Scan(&num)
	if err != nil {
		log.Println(err)
	}
	return num, err
}

// 获取颜色列表
func selectColor(sql, goodsid string, db *sql.DB) ([]ColorList, error) {
	var CLS []ColorList
	rows, err := db.Query(sql, goodsid)
	if err != nil {
		log.Println(err)
		return CLS, err
	}
	defer rows.Close()
	for rows.Next() {
		var CL ColorList
		err := rows.Scan(&CL.ID, &CL.GoodsID, &CL.Model, &CL.CompressPic, &CL.IsFirst)
		if err != nil {
			log.Println(err)
			return CLS, err
		}
		CLS = append(CLS, CL)
	}
	return CLS, err
}

// 删除颜色
func deleteColor(id string, db *sql.DB) error {
	updateSQL := "update goodspic set state = 0 where id = ?"
	_, err := db.Exec(updateSQL, id)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 把该商品的所有颜色isfirst设为0
func updateColorZeroing(goodsid string, tx *sql.Tx) error {
	updateSQL := "update goodspic set isfirst = 0 where goodsid = ?"
	_, err := tx.Exec(updateSQL, goodsid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 修改需要置顶的颜色isfirst
func updateTopColor(colorid string, tx *sql.Tx) error {
	updateSQL := "update goodspic set isFirst = 1 where id = ?"
	_, err := tx.Exec(updateSQL, colorid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 取消颜色置顶
func updateColorRemoveTop(colorid string, db *sql.DB) error {
	updateSQL := "update goodspic set isFirst = 0 where id = ?"
	_, err := db.Exec(updateSQL, colorid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 热门商品展示配置查询
func selectHotGoodsShow(shopid string, db *sql.DB) (interface{}, error) {
	var HGS HotGoodsShow
	selectSQL := "select popularlimit, popularquantity from shop where id = ?"
	err := db.QueryRow(selectSQL, shopid).Scan(&HGS.PopularLimit, &HGS.PopularQuantity)
	if err != nil {
		log.Println(err)
	}
	return HGS, err
}

// 热门商品展示配置修改
func updateHotGoodsShow(shopid, popularlimit, popularquantity string, db *sql.DB) error {
	updateSQL := "update shop set popularlimit = ?, popularquantity = ? where id = ?"
	_, err := db.Exec(updateSQL, popularlimit, popularquantity, shopid)
	if err != nil {
		log.Println(err)
	}
	return err
}

// 获取一级菜单下二级菜单的数量
func get2levelByMenuIDDao(menuid string, db *sql.DB) (int, error) {
	var num int
	selectSQL := `select count(*) as num from menu m where m.superid = ?`
	err := db.QueryRow(selectSQL, menuid).Scan(&num)
	if err != nil {
		log.Println(err)
	}
	return num, err
}

// 获取二级菜单下商品的数量
func getGoodsByMenuIDDao(menuid string, db *sql.DB) (int, error) {
	var num int
	selectSQL := `select count(*) as num from menu m
	left join goods g on m.id = g.menuid
	where m.id = ? and g.state <> 2`
	err := db.QueryRow(selectSQL, menuid).Scan(&num)
	if err != nil {
		log.Println(err)
	}
	return num, err
}

// 查询店铺详情
func getShopInfoDao(shopid string, db *sql.DB) (interface{}, error) {
	var s ShopInfo
	selectSQL := `select id, shopurl, shopname, logoimg, shopshow, bannerimg, bossname, telnum, phonenum, wechat, wechaturl, 
	popularlimit, popularquantity, location, lng, lat, state from shop where id = ?`
	err := db.QueryRow(selectSQL, shopid).Scan(&s.ID, &s.ShopURL, &s.ShopName, &s.LogoImg, &s.ShopShow, &s.BannerImg, &s.BossName, &s.TelNum,
		&s.PhoneNum, &s.Wechat, &s.WechatURL, &s.PopularLimit, &s.PopularQuantity, &s.Location, &s.Lng, &s.Lat, &s.State)
	if err != nil {
		log.Println(err)
	}
	return s, err
}
