import axios from 'axios'
import '../mock'
//用户登录
export const userLogin=(params)=>{
    return axios.post(`/userLogin`,params)
}
//获取店铺首页信息
export const getIndexInfo=(params)=>{
    return axios.post(`/getIndexInfo`,params)
}
//获取联系方式
export const getShopContactInfo=(params)=>{
    return axios.post(`/getShopContactInfo`,params)
}

//获取全部分类
export const getAllClass=(params)=>{
    return axios.post('/getAllClass',params)
}
// export const getAllClassifyOne=(params)=>{
//     return axios.post('/getAllClassifyOne',params)
// }

//获取分类下的所有商品
export const getGoods=(params)=>{
    return axios.post('/getGoods',params)
}

//商品详情
export const getGoodsDetail=(params)=>{
    return axios.post('/getGoodsDetail',params)
}

//获得用户收藏的所有
export const getCollection=(params)=>{
    return axios.post('/getCollection',params)
}

//取消收藏
export const delCollection=(params)=>{
    return axios.post('/delCollection',params)
}
