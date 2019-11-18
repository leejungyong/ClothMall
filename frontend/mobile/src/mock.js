import Mock from 'mockjs';
const Random = Mock.Random
Mock.mock(`/userLogin`, 'post', {
    success: true,
    message: "",
    data: null
})

//首页
Mock.mock(`/getIndexInfo`, 'post', {
    success: true,
    errmsg: '',
    data: {
        shopname: Random.csentence(6, 12),
        logoimg: Random.dataImage('30x30', 'logo'),
        shopshow: Random.dataImage('335x130', ''),
        bannerImg: Random.dataImage('335x180', ''),
        hotgoods: [{
            id: Random.dataImage('158x178', '热门商品'),
            menuname: Random.csentence(3, 4),
            clicknum: 100,
            colornum: 16,
            smallpic: Random.dataImage('158x178', 'beijing')
        }]
    }
})

//店铺联系方式

Mock.mock(`/getShopContactInfo`, 'post', {
    success: true,
    errmsg: '',
    data: {
        telnum: Random.natural(11),
        phonenum: Random.natural(11),
        wechat: Random.string(14, 15),
        wechaturl: Random.dataImage('60x60', 'code'),
        location: Random.csentence(6, 12),
        lng: '120',
        lat: '30'
    }
})

let all = [],
    goodslist = []
for (let i = 0; i < 10; i++) {
    let t = {
        id: i,
        menuname: Random.csentence(4)
    }
    let good = {
        id: i,
        clicknum: Random.natural(10, 20),
        smallpic: Random.dataImage('158x178', '热门商品'),
        colornum: Random.natural(10, 20)
    }
    all.push(t)
    goodslist.push(good)
}

// Mock.mock(`/getAllClass`,'post',{
//     success:true,
//     errmsg:'',
//     data:{
//         id:1,
//         menuname:'墙纸类',
//         classList:all
//     }
// })

//获取全部分类
Mock.mock(`/getAllClass`, 'post', {
    success: true,
    errmsg: '',
    data: [{
            id: '1',
            menuname: '墙纸类',
            twoLevelClass: [{
                id: '7',
                menuname: '英式风格',
                pic: '', //？pic是放哪里的？？
                num: 1
            }]
        },
        {
            id: '2',
            menuname: '墙布类',
            twoLevelClass: [{
                id: '10',
                menuname: '现代风格',
                pic: '', //？pic是放哪里的？？
                num: 2
            }, ]
        },
        {
            id: '3',
            menuname: '其他类',
            twoLevelClass: [{
                id: '10',
                menuname: '中式风格',
                pic: '', //？pic是放哪里的？？
                num: 3
            }]
        }
    ]
})

//获取商品分类下的所有商品
Mock.mock(`/getGoods`, 'post', {
    success: true,
    errmsg: '',
    data: {
        total: Random.natural(1, 3),
        goodslist: goodslist
    }
})

//某一商品详情
Mock.mock(`/getGoodsDetail`, 'post', {
    success: true,
    errmsg: '',
    data: {
        id: 1,
        goodsname: Random.csentence(8, 12),
        brand: Random.csentence(5, 8),
        style: Random.csentence(2, 4),
        unit: '400*500',
        material: Random.csentence(2, 4),
        madein: Random.csentence(2, 3),
        price: '价格：88/m2',
        machineid: 'a01',
        machineadder: '030',
        picList: {
            id: 2,
            model: 'JBL-1302A',
            pic: Random.dataImage('350x350', '大图'),
            compresspic: [Random.dataImage('350x350', '小图'),Random.dataImage('350x350', '小图'),Random.dataImage('350x350', '小图'),Random.dataImage('350x350', '小图'), Random.dataImage('350x350', '小图'), Random.dataImage('350x350', '小图'), Random.dataImage('350x350', '小图'), Random.dataImage('350x350', '小图'), Random.dataImage('350x350', '小图')]
        }
    }
})

//获得用户的收藏的所有商品
Mock.mock(`/getCollection`,'post',{
    success:true,
    errmsg:'',
    data:{
        Record:[
            {
                id:1,
                nagoodsIDme:2,
                collectionType:'卧室',
                goodsName:Random.csentence(4,8),
                model:'JBL-1302A',
                compressPic:Random.dataImage('350x307', '大图'),
                price:'80',
                unit:'400*400',
                scenesURL:'???'
            }
        ],
        ShopCollection:[
            {
                id:1,
                nagoodsIDme:2,
                collectionType:'卧室',
                goodsName:Random.csentence(4,8),
                model:'JBL-1302A',
                compressPic:Random.dataImage('350x307', '大图'),
                price:'80',
                unit:'400*400',
                scenesURL:'???'
            }
        ],
        ScenesCollection:[
            {
                id:1,
                nagoodsIDme:2,
                collectionType:'卧室',
                goodsName:Random.csentence(4,8),
                model:'JBL-1302A',
                price:'80',
                unit:'400*400',
                compressPic:Random.dataImage('350x307', '大图'),
                scenesURL:'???'
            }
        ]
    }
})

//取消收藏 
Mock.mock('/delCollection','post',{
    success:true,
    errmsg:'',
    data:''
})