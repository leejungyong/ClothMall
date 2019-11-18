// Package switcher 数据库交互函数
package switcher

import (
	"database/sql"
	"errors"
	"log"
)

// testDao 测试接口
func testDao(id string, db *sql.DB) ([]AName, error) {
	var aname []AName
	sql := "select name from test"
	log.Println(sql)
	rows, err := db.Query(sql)

	if err != nil {
		log.Println(err)
		return aname, err
	}
	defer rows.Close()

	for rows.Next() {
		var aaname AName
		var name string
		err := rows.Scan(&name)
		if err != nil {
			log.Println(err)
			return aname, err
		}
		aaname.Name = name
		aname = append(aname, aaname)
	}
	return aname, err
}

// 用户收藏查询
func getCollectionDao(selectSQL string, db *sql.DB) ([]Collection, error) {
	var collection []Collection
	rows, err := db.Query(selectSQL)
	if err != nil {
		log.Println(err)
		return collection, err
	}
	defer rows.Close()

	for rows.Next() {
		var c Collection
		var id, goodsID int
		var collectionType, goodsName, style, unit, price, width, height, compressPic, scenesURL string

		err := rows.Scan(&id, &goodsID, &collectionType, &goodsName, &style, &unit, &price, &width, &height, &compressPic, &scenesURL)
		if err != nil {
			log.Println(err)
			return collection, err
		}
		c.ID = id
		c.GoodsID = goodsID
		c.CollectionType = collectionType
		c.GoodsName = goodsName
		c.Style = style
		c.Unit = unit
		c.Price = price
		c.Width = width
		c.Height = height
		c.CompressPic = compressPic
		c.ScenesURL = scenesURL

		collection = append(collection, c)
	}
	return collection, err
}

// 用户增加收藏
func insertCollectionDao(userid, shopid, goodsid int, collectionType, scenesURL string, tx *sql.Tx) (int64, error) {
	var rowid int64
	// 增加收藏sql
	insertSQL := `insert into collection (userid,shopid,goodsid,type,scenesurl) values (?, ?, ?, ?, ?)`
	result, err := tx.Exec(insertSQL, userid, shopid, goodsid, collectionType, scenesURL)
	if nil != err {
		log.Println(err)
		return rowid, err
	}
	rowid, err = result.LastInsertId()
	return rowid, nil
}

// 用户取消收藏
func delCollectionDao(userid, collectionid int, db *sql.DB) error {
	updateSQL := `UPDATE collection SET status = 0 WHERE id = ? and userid = ?`
	_, err := db.Exec(updateSQL, userid, collectionid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 获取机器状态
func getMachineDao(machineid int, db *sql.DB) (Machine, error) {
	var m Machine
	selectSQL := `select id,machineid,shopid,slotnum,runstate,netstate,machineip from machine where state = 1 and id = ` + IntToString(machineid)
	err := db.QueryRow(selectSQL).Scan(&m.ID, &m.MachineID, &m.ShopID, &m.SlotNum, &m.RunState, &m.NetState, &m.MachineIP)
	if err != nil {
		log.Println(err)
		return m, err
	}
	return m, err
}

// 实体机器商品展示
func insertMachineTaskDao(machineid, userid, shopid, goodsid, slotid int, tx *sql.Tx) error {
	// 增加机器任务sql
	insertSQL := `insert into machinetask (machineid,shopid,userid,goodsid,slotid) values (?, ?, ?, ?, ?)`
	_, err := tx.Exec(insertSQL, machineid, shopid, userid, goodsid, slotid)
	if nil != err {
		log.Println(err)
		return err
	}
	return nil
}

// 增加模拟场景任务
func insertScenesTaskDao(shopid, userid, scenesmenuid1, scenesmenuid2, goodsAid, goodsBid int,
	scenesid, scenesmenuurl1, scenesmenuurl2, goodsAurl, goodsAwidth, goodsAheight, goodsAspliceType, goodsBurl,
	goodsBwidth, goodsBheight, goodsBspliceType string, tx *sql.Tx) (int64, error) {
	var rowid int64
	// 增加机器任务sql
	insertSQL := `insert into scenestask (shopid, userid, scenesid, scenesmenuid1, scenesmenuid2, goodsAid, goodsBid,
		scenesmenuurl1, scenesmenuurl2, goodsAurl, goodsAwidth, goodsAheight, goodsAspliceType, goodsBurl, goodsBwidth, goodsBheight, goodsBspliceType) 
		values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	result, err := tx.Exec(insertSQL, shopid, userid, scenesid, scenesmenuid1, scenesmenuid2, goodsAid, goodsBid, scenesmenuurl1,
		scenesmenuurl2, goodsAurl, goodsAwidth, goodsAheight, goodsAspliceType, goodsBurl, goodsBwidth, goodsBheight, goodsBspliceType)
	if nil != err {
		log.Println(err)
		return rowid, err
	}
	rowid, err = result.LastInsertId()
	return rowid, nil
}

// 查询模拟场景URL
func getScenesTaskURLDao(scenesTaskid int, db *sql.DB) (string, error) {
	var scenesTaskURL string
	selectSQL := `select url from scenestask where state = 1 and id = ` + IntToString(scenesTaskid)
	err := db.QueryRow(selectSQL).Scan(&scenesTaskURL)
	if err != nil {
		log.Println(err)
		return scenesTaskURL, err
	}
	return scenesTaskURL, err
}

// 根据url获取店铺id name
func getShopID(shopurl string, db *sql.DB) (UrlToId, error) {
	var UTI UrlToId
	selectSQL := "select id, shopname from shop where shopurl = ?"
	err := db.QueryRow(selectSQL, shopurl).Scan(&UTI.ID, &UTI.ShopName)
	if err != nil {
		return UTI, err
	}
	return UTI, err
}

// 用户登录验证
func getUserLogin(phoneNum, password string, db *sql.DB) (string, error, interface{}) {
	var UI UserInfo
	var num int
	selectSQL := "select ifnull(id, 0), ifnull(count(*), 0) from user where phoneNum = ? and password = ? and state = 1"
	err := db.QueryRow(selectSQL, phoneNum, password).Scan(&UI.ID, &num)
	if err != nil {
		return "查询失败", err, UI
	}
	if num != 0 {
		return "登录成功", err, UI
	}
	return "手机号或密码错误", errors.New("手机账号或密码错误"), UI
}

// 店铺信息查询 包括 首页展示商品的数量和阀值
func getIndex(shopid string, db *sql.DB) (IndexInfo, HotCondition, error) {
	var II IndexInfo
	var HC HotCondition
	selectSQL := "select id, shopname, logoimg, shopshow, bannerimg, popularlimit, popularquantity from shop where id = ?"
	err := db.QueryRow(selectSQL, shopid).Scan(&II.ID, &II.ShopName, &II.LogoImg, &II.ShopShow, &II.BannerImg, &HC.PopularLimit, &HC.PopularQuantity)
	if err != nil {
		return II, HC, err
	}
	return II, HC, err
}

// 爆款查询
func getHotGoods(shopid string, HC HotCondition, CLS []ClassList, db *sql.DB) ([]HotGoods, error) {
	var HGS []HotGoods
	for _, k := range CLS {
		selectSQL := "select g.id, m.menuname, ifnull(x.clicknum, 0) as clicknum, ifnull(y.colornum, 0) as colornum, ifnull(y.smallpic, '') as smallpic from goods g left join menu m on g.menuid = m.id left join (select goodsid, count(*) as clicknum from goodsaccess a group by goodsid) x on g.id = x.goodsid left join (select goodsid, smallpic, count(*) as colornum from goodspic group by goodsid) y on g.id = y.goodsid where g.state = 1 and g.shopid = ? and menuid = ? and ifnull(x.clicknum, 0) > ? order by ifnull(y.colornum, 0) desc limit ?"
		rows, err := db.Query(selectSQL, shopid, k.ID, HC.PopularLimit, HC.PopularQuantity)
		if err != nil {
			log.Println(err)
			return HGS, err
		}
		defer rows.Close()
		for rows.Next() {
			var HG HotGoods
			err := rows.Scan(&HG.ID, &HG.MenuName, &HG.ClickNum, &HG.ColorNum, &HG.SmallPic)
			if err != nil {
				log.Println(err)
				return HGS, err
			}
			HGS = append(HGS, HG)
		}
	}
	return HGS, nil
}

// 店铺联系方式查询
func getShopContactInfo(shopid string, db *sql.DB) (interface{}, error) {
	var SCI ShopContactInfo
	selectSQL := "select telnum, phonenum, wechat, wechaturl, location, lng, lat from shop where id = ?"
	err := db.QueryRow(selectSQL, shopid).Scan(&SCI.Telnum, &SCI.Phonenum, &SCI.Wechat, &SCI.Wechaturl, &SCI.Location, &SCI.Lng, &SCI.Lat)
	if err != nil {
		return "查询失败", err
	}
	return SCI, err
}

// 全部分类查询(一级分类查询)
func getAllClassifyOne(shopid string, db *sql.DB) ([]AllClassOne, error) {
	var ACOS []AllClassOne
	selectSQL := "select id, menuname, superid from menu where state = 1 and superid = 0 and shopid = ?"
	rows, err := db.Query(selectSQL, shopid)
	if err != nil {
		log.Println(err)
		return ACOS, err
	}
	defer rows.Close()
	for rows.Next() {
		var ACO AllClassOne
		err := rows.Scan(&ACO.ID, &ACO.MenuName, &ACO.SuperID)
		if err != nil {
			log.Println(err)
			return ACOS, err
		}
		ACOS = append(ACOS, ACO)
	}
	return ACOS, err
}

// 全部分类查询(二级分类查询 带图片数量)
func getAllClassifyTwoWithPic(shopid, superid string, db *sql.DB) ([]AllClassTwo, error) {
	var ACTS []AllClassTwo
	selectSQL := `select z.id as menuid, z.menuname, z.superid, ifnull(y.num, 0) as num from (select id, menuname, superid from menu where superid = ? and shopid = ? and state = 1) z left join (select id, menuid, count(*) as num from goods where state = 1 and shopid = ? group by menuid) y on z.id = y.menuid group by z.id`
	rows, err := db.Query(selectSQL, superid, shopid, shopid)
	if err != nil {
		log.Println(err)
		return ACTS, err
	}
	defer rows.Close()
	for rows.Next() {
		var ACT AllClassTwo
		err := rows.Scan(&ACT.ID, &ACT.MenuName, &ACT.SuperID, &ACT.Num)
		if err != nil {
			log.Println(err)
			return ACTS, err
		}
		ACTS = append(ACTS, ACT)
	}
	return ACTS, err
}

// 全部选择分类查询
func getAllClassifyTwo(superid string, db *sql.DB) ([]ClassList, error) {
	var CLS []ClassList
	selectSQL := "select id, menuname from menu where superid = ? and state = 1"
	rows, err := db.Query(selectSQL, superid)
	if err != nil {
		log.Println(err)
		return CLS, err
	}
	defer rows.Close()
	for rows.Next() {
		var CL ClassList
		err := rows.Scan(&CL.ID, &CL.MenuName)
		if err != nil {
			log.Println(err)
			return CLS, err
		}
		CLS = append(CLS, CL)
	}
	return CLS, err
}

// 获取二级菜单下的商品总数
func getGoodsNum(menuid string, db *sql.DB) (int, error) {
	var total int
	selectSQL := "select count(*) from goods where menuid = ? and state = 1"
	rows, err := db.Query(selectSQL, menuid)
	if err != nil {
		log.Println(err)
		return total, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&total)
		if err != nil {
			log.Println(err)
			return total, err
		}
	}
	return total, err
}

// 获取商品信息
func getGoods(menuid string, db *sql.DB) ([]GoodsList, error) {
	var GLS []GoodsList
	selectSQL := `select g.id, ifnull(t.num, 0) as clickNum, ifnull(x.smallpic, '') as smallpic, ifnull(x.num, 0) as colorNum 
	from goods g 
	left join (select goodsid, count(goodsid) as num from goodsaccess group by goodsid) t on g.id = t.goodsid 
	left join (select goodsid, smallpic, count(*) as num from goodspic where state = 1 group by goodsid) x on g.id = x.goodsid 
	where g.menuid = ?`
	rows, err := db.Query(selectSQL, menuid)
	if err != nil {
		log.Println(err)
		return GLS, err
	}
	defer rows.Close()
	for rows.Next() {
		var GL GoodsList
		err := rows.Scan(&GL.ID, &GL.ClickNum, &GL.SmallPic, &GL.ColorNum)
		if err != nil {
			log.Println(err)
			return GLS, err
		}
		GLS = append(GLS, GL)
	}
	return GLS, err
}

// 获取商品详细内容
func getGoodsInfoDetail(goodsid string, db *sql.DB) (GoodsDetail, error) {
	var GD GoodsDetail
	selectSQL := `select g.id, g.goodsname, g.brand, g.style, g.unit, g.material, g.madein, g.price, g.width, g.height,
	ifnull(s.machineid, 0) as machineid, ifnull(m.name, '') as machinename, ifnull(s.id, 0) as slotid, 
	ifnull(s.slotnum, 0) as slotnum from goods g 
	left join machineslot s on s.goodsid = g.id 
	left join machine m on m.id = s.machineid 
	where g.id = ?`
	err := db.QueryRow(selectSQL, goodsid).Scan(&GD.ID, &GD.GoodsName, &GD.Brand, &GD.Style, &GD.Unit, &GD.Material, &GD.MadeIn, &GD.Price,
		&GD.Width, &GD.Height, &GD.MachineDetail.MachineID, &GD.MachineDetail.MachineName, &GD.MachineDetail.SlotID, &GD.MachineDetail.SlotNum)
	if err != nil {
		log.Println(err)
	}
	return GD, err
}

// 获取商品是否收藏
func getIsCollect(goodsid, userid string, db *sql.DB) (int, int, error) {
	var num int
	var collectionid int
	selectSQL := "select ifnull(id,0), count(*) from collection where goodsid = ? and userid = ? and state = 1"
	err := db.QueryRow(selectSQL, goodsid, userid).Scan(&collectionid, &num)
	if err != nil {
		log.Println(err)
	}
	return collectionid, num, err
}

// 获取商品图片
func getGoodsPicDetail(goodsid string, db *sql.DB) ([]PicList, error) {
	var PLS []PicList
	selectSQL := "select id, model, pic, compresspic from goodspic where state = 1 and goodsid = ? order by isfirst desc"
	rows, err := db.Query(selectSQL, goodsid)
	if err != nil {
		log.Println(err)
		return PLS, err
	}
	defer rows.Close()
	for rows.Next() {
		var PL PicList
		err := rows.Scan(&PL.ID, &PL.Model, &PL.Pic, &PL.Compresspic)
		if err != nil {
			log.Println(err)
			return PLS, err
		}
		PLS = append(PLS, PL)
	}
	return PLS, err
}

// 店铺访问记录增加
func insertShopAccess(shopid, userid, ip string, db *sql.DB) (string, error) {
	if shopid == "" {
		return "店铺id参数为空", errors.New("店铺id参数为空")
	}
	insertSQL := "insert into shopaccess (shopid, userid, ip) values (?, ?, ?)"
	_, err := db.Exec(insertSQL, shopid, userid, ip)
	if err != nil {
		log.Println(err)
		return "插入店铺访问记录失败", err
	}
	return "插入店铺访问记录成功", err
}

// 商品访问记录增加
func insertGoodsAccess(goodsid, userid, ip string, db *sql.DB) (string, error) {
	if goodsid == "" {
		return "商品id参数为空", errors.New("商品id参数为空")
	}
	insertSQL := "insert into goodsaccess (goodsid, userid, ip) values (?, ?, ?)"
	_, err := db.Exec(insertSQL, goodsid, userid, ip)
	if err != nil {
		log.Println(err)
		return "插入商品访问记录失败", err
	}
	return "插入商品访问记录成功", err
}
