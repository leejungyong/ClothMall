import React,{PureComponent} from 'react'
import {connect} from 'dva'
import axios from 'axios'
import qs from 'qs'
import {
    Input,
    Form,
    Icon,
    Button,
    Cascader    
} from 'antd'

import PageHeaderWrapper from '@/components/PageHeaderWrapper';

import { router } from 'umi';

import styles from './ShopInfo.less'

const formItemLayout = {
    labelCol: {
      xs: { span: 0 },
      sm: { span: 3 },
    },
    wrapperCol: {
      xs: { span: 10 },
      sm: { span: 10 },
    }
  };
@Form.create()
class ShopInfo extends PureComponent{
    state={
        shopInfo:{},
        shopid:'1'
    }
  async  componentDidMount(){
      await this.setState({
          shopid:localStorage.getItem('shopid')
      })
        this.getInfo()
    }
    getInfo(){
        axios.post('/api2',qs.stringify({cmd:'getShopInfo',token:'Jh2044695',id:this.state.shopid})).then(res=>{
            if(res.data.success){
                console.log(res.data.data)
                this.setState({
                    shopInfo:res.data.data
                })
            }
        })
    }
    render(){
        const {getFieldDecorator}=this.props.form
        const {shopInfo}=this.state
        return(
                <div className={styles.shopinfo}>
                    <div className={styles.com}><span className={styles.star}>*</span>店铺URL标识:  {shopInfo.shopurl}</div>
                    <div className={styles.com}><span className={styles.star}>*</span>店铺名称:  {shopInfo.shopname}</div>
                    <div className={styles.com}><span className={styles.star}>*</span>老板姓名:  {shopInfo.bossname}</div>
                    <div className={styles.com}><span className={styles.star}>*</span>座机号码:  {shopInfo.telnum}</div>
                    <div className={styles.com}><span className={styles.star}>*</span>手机号码:  {shopInfo.phonenum}</div>
                    <div className={styles.com}><span className={styles.star}>*</span>地址:  {shopInfo.location}</div>
                    <div className={styles.imgbox}><span className={styles.star}>*</span>店铺logo: <img src={shopInfo.logoimg} className={styles.img}/> </div>
                    <div className={styles.imgbox}><span className={styles.star}>*</span>店铺信息:  <img src={shopInfo.shopshow} className={styles.img}/></div>
                    <div className={styles.imgbox}><span className={styles.star}>*</span>广告图:  <img src={shopInfo.bannerimg} className={styles.img}/></div>
                </div>
        )
    }
}

export default ShopInfo 
