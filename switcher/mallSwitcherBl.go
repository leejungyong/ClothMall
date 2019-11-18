// Package switcher 业务逻辑函数
package switcher

import (
	"database/sql"
	"log"
	"strings"
)

// 查询用户收藏信息（收藏的商品和收藏的场景）
func getCollection(userid, shopid int, db *sql.DB) (AllCollection, error) {
	var allCollection AllCollection
	var err error
	// 拼sql
	shopCollectionSQL, urlCollectionSQL := getCollectionSQL(userid, shopid)
	log.Println(shopCollectionSQL)
	log.Println(urlCollectionSQL)
	// 收藏的商品
	shopCollection, err := getCollectionDao(shopCollectionSQL, db)
	if nil != err {
		log.Println(err)
		return allCollection, err
	}
	// 收藏的场景
	scenesCollection, err := getCollectionDao(urlCollectionSQL, db)
	if nil != err {
		log.Println(err)
		return allCollection, err
	}
	allCollection.ShopCollection = shopCollection
	allCollection.ScenesCollection = scenesCollection

	return allCollection, err
}

// 生成根据用户ID和店铺ID获取收藏商品和收藏场景的SQL
func getCollectionSQL(userid, shopid int) (string, string) {
	// 查询收藏商品的sql
	sql1 := `select c.id,g.id as goodsid,c.type,g.goodsname,g.style,g.unit,g.price,g.width,g.height,g.compresspic,c.scenesurl from collection c 
	left join (select g.id,g.shopid,g.goodsname,g.style,g.unit,g.price,g.width,g.height,gp.compresspic from goods g
		left join goodspic gp on g.id = gp.goodsid
		where gp.isfirst = 1 and g.state = 1) g on c.goodsid = g.id
	where c.state = 1 and c.scenesurl = '' %condition%`

	// 查询收藏场景的sql
	sql2 := `select c.id,g.id as goodsid,c.type,g.goodsname,g.style,g.unit,g.price,g.width,g.height,g.compresspic,c.scenesurl from collection c 
	left join (select g.id,g.shopid,g.goodsname,g.style,g.unit,g.price,g.width,g.height,gp.compresspic from goods g
		left join goodspic gp on g.id = gp.goodsid
		where gp.isfirst = 1 and g.state = 1) g on c.goodsid = g.id
	where c.state = 1 and c.scenesurl <> '' %condition%`
	// 拼查询条件
	var condition string
	// 查询条件
	condition = " and g.shopid = '" + IntToString(shopid) + "' and c.userid = " + IntToString(userid)

	return strings.Replace(sql1, "%condition%", condition, 1), strings.Replace(sql2, "%condition%", condition, 1)
}

// 用户增加收藏
func insertCollection(userid, shopid, goodsid int, collectionType, scenesURL string, db *sql.DB) (int64, error) {
	var err error
	var rowID int64
	// 事务开启
	tx, err := db.Begin()
	rowID, err = insertCollectionDao(userid, shopid, goodsid, collectionType, scenesURL, tx)
	if nil != err {
		log.Println(err)
		return rowID, err
	}

	tx.Commit()

	return rowID, err
}

// 用户取消收藏
func delCollection(userid, collectionid int, db *sql.DB) error {
	var err error
	err = delCollectionDao(userid, collectionid, db)
	if nil != err {
		log.Println(err)
		return err
	}
	return err
}

// 实体机器状态检查
func machineCheck(machineid int, userIP string, db *sql.DB) (string, error) {
	var machine Machine
	var err error

	// 获取机器状态
	machine, err = getMachineDao(machineid, db)
	if nil != err {
		log.Println(err)
		return "2", err
	}

	// 判断ip地址是否相同
	if userIP != machine.MachineIP {
		return "3", err
	}
	// 判断机器是否联网
	if machine.NetState != "1" {
		return "4", err
	}
	// 判断机器运行状态
	if machine.RunState != "1" {
		return "5", err
	}
	return "1", err
}

// 实体机器商品展示
func insertMachineTask(machineid, userid, shopid, goodsid, slotid int, db *sql.DB) error {
	var err error
	// 事务开启
	tx, err := db.Begin()
	err = insertMachineTaskDao(machineid, userid, shopid, goodsid, slotid, tx)
	if nil != err {
		log.Println(err)
		return err
	}

	tx.Commit()

	return err
}

// 增加模拟场景任务
func insertScenesTask(shopid, userid, scenesmenuid1, scenesmenuid2, goodsAid, goodsBid int,
	scenesmenuurl1, scenesmenuurl2, goodsAurl, goodsAwidth, goodsAheight, goodsAspliceType, goodsBurl,
	goodsBwidth, goodsBheight, goodsBspliceType string, db *sql.DB) (int64, error) {
	var rowid int64
	var err error
	var scenesid string
	// 拼模拟场景标识（由shopid_goodsAid_goodsBid拼接而成）
	scenesid = IntToString(shopid) + "_" + IntToString(goodsAid) + "_" + IntToString(goodsBid)
	// 事务开启
	tx, err := db.Begin()
	rowid, err = insertScenesTaskDao(shopid, userid, scenesmenuid1, scenesmenuid2, goodsAid, goodsBid, scenesid, scenesmenuurl1,
		scenesmenuurl2, goodsAurl, goodsAwidth, goodsAheight, goodsAspliceType, goodsBurl, goodsBwidth, goodsBheight, goodsBspliceType, tx)
	if nil != err {
		log.Println(err)
		return rowid, err
	}

	tx.Commit()

	return rowid, err
}

// 查询模拟场景URL
func getScenesTaskURL(scenesTaskid int, db *sql.DB) (string, error) {
	var err error
	var scenesTaskURL string
	scenesTaskURL, err = getScenesTaskURLDao(scenesTaskid, db)
	if nil != err {
		log.Println(err)
		return scenesTaskURL, err
	}

	return scenesTaskURL, err
}

// 获取首页信息
func getIndexInfo(shopid string, db *sql.DB) (interface{}, error) {
	var IIAll IndexInfo
	// 店铺信息查询 包括 首页展示商品的数量和阀值
	II, HC, err := getIndex(shopid, db)
	if nil != err {
		log.Println(err)
		return IIAll, err
	}
	// 店铺二级菜单查询
	CLS, err := getAllClassifyTwo(shopid, db)
	if nil != err {
		log.Println(err)
		return IIAll, err
	}
	// 店铺爆款查询
	HGS, err := getHotGoods(shopid, HC, CLS, db)
	if nil != err {
		log.Println(err)
		return IIAll, err
	}
	IIAll = II
	IIAll.HotGoods = HGS
	return IIAll, err
}

// 获取全部分类(带颜色、点击量数量信息)
func getAllClass(shopid string, db *sql.DB) (interface{}, error) {
	var ACS []AllClass
	// 获取所有一级分类
	ACOS, err := getAllClassifyOne(shopid, db)
	if nil != err {
		log.Println(err)
		return ACS, err
	}
	for _, v := range ACOS {
		var AC AllClass
		// 获取每一个一级分类下的二级分类
		AC.ID = v.ID
		AC.MenuName = v.MenuName
		AC.SuperID = v.SuperID
		AC.Children, err = getAllClassifyTwoWithPic(shopid, v.ID, db)
		if nil != err {
			log.Println(err)
			return ACS, err
		}
		ACS = append(ACS, AC)
	}
	return ACS, err
}

// 获取二级菜单下的所有商品
func getAllGoods(menuid string, db *sql.DB) (interface{}, error) {
	var GS GoodsSearch
	// 获取商品总数
	total, err := getGoodsNum(menuid, db)
	if nil != err {
		log.Println(err)
		return GS, err
	}
	// 获取商品图片 点击量 颜色种类
	GL, err := getGoods(menuid, db)
	if nil != err {
		log.Println(err)
		return GS, err
	}
	GS.Total = total
	GS.GoodsList = GL
	return GS, err
}

// 获取商品详细信息
func getGoodsDetail(goodsid, userid string, db *sql.DB) (interface{}, error) {
	var GD GoodsDetail
	// 获取商品详细内容
	GDInfo, err := getGoodsInfoDetail(goodsid, db)
	if nil != err {
		log.Println(err)
		return GD, err
	}
	// 获取商品图片
	GDPic, err := getGoodsPicDetail(goodsid, db)
	if nil != err {
		log.Println(err)
		return GD, err
	}
	// 获取商品是否收藏
	collectionid, isCollectionInt, err := getIsCollect(goodsid, userid, db)
	log.Println(collectionid)
	log.Println(isCollectionInt)
	var isCollection bool
	if isCollectionInt == 1 {
		isCollection = true
	}
	GD = GDInfo
	GD.PicList = GDPic
	GD.IsCollection = isCollection
	GD.CollectionID = collectionid
	return GD, nil
}

// 获取模拟场景房间信息
func getMallScenesMenuList(superid string, db *sql.DB) ([]*ScenesMenu, error) {
	var err error
	var selectSQL string
	var scenesMenu []*ScenesMenu

	// superid为0是一级菜单，反之为二级菜单
	if superid == "0" {
		selectSQL = `select id,name,superid,pic from scenesmenu s 
		where s.state = 1 and s.superid = 0`
	} else {
		selectSQL = `select id,name,superid,pic from scenesmenu s 
		where s.state = 1 and s.superid = ` + superid
	}
	log.Println(selectSQL)
	// 获取模拟场景房间信息
	scenesMenu, err = getScenesMenuListDao(selectSQL, db)
	if nil != err {
		log.Println(err)
		return scenesMenu, err
	}
	return scenesMenu, err
}
