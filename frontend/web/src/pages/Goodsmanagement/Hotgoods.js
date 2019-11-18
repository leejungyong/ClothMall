import React, {PureComponent,Fragment} from 'react'
import {
    Form,
    Input,
    Button,
    message
} from 'antd'

import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import {styles} from './Hotgoods.less'
import axios from 'axios';
import qs from 'qs'

// const formItemLayout = {
//     labelCol: {
//       xs: { span: 0 },
//       sm: { span: 3 },
//     },
//     wrapperCol: {
//       xs: { span: 10 },
//       sm: { span: 10 },
//     },
//   };
@Form.create()
class Hotgoods extends PureComponent{
    state={
        popularlimit:'',
        popularquantity:'',
        minclick:'',
        num:'',
        isShowEdit:false,
        shopid:1
    }
    async componentDidMount(){
        await this.setState({
            shopid:localStorage.getItem('shopid')
        })
        this.getInfo()
    }
    //获取热门配置信息
    getInfo(){
        axios.post('/api2',qs.stringify({cmd:'selectHotGoodsShow',token:'Jh2044695',shopid:this.state.shopid})).then(res=>{
            console.log(res.data)
            if(res.data.success){
                this.setState({
                    minclick:res.data.data.popularlimit,
                    num:res.data.data.popularquantity
                })
            }else{
                message.error(res.data.msg)
            }
        })
    }
    //提交
    submitEdit=(e)=>{
        e.preventDefault()
        this.props.form.validateFields((err,values)=>{
            console.log(values)
            if(!err){
                axios.post('/api2',qs.stringify({
                    cmd:'updateHotGoodsShow',
                    token:'Jh2044695',
                    shopid:this.state.shopid,
                    popularlimit:parseInt(values.popularlimit) ,
                    popularquantity:parseInt(values.popularquantity) 
                })).then(res=>{
                    if(res.data.success){
                        message.success('修改成功')
                        this.setState({
                            isShowEdit:false
                        })
                        this.getInfo()
                    }else{
                        message.error(res.data.errmsg)
                    }
                })
            }
        })
        // console.log(values)
    }
    //取消修改
    cancelEdit=()=>{
        this.setState({
            isShowEdit:false
        })
    }
    //显示修改
    showEdit=()=>{
        this.setState({
            isShowEdit:true
        })
    }
    render(){
        const { getFieldDecorator } = this.props.form;
        return (
            <PageHeaderWrapper title='热门商品展示配置'>
                
                {
                    this.state.isShowEdit?(   <div >
                    <Form layout='inline'  onSubmit={this.submitEdit}>
                    <span style={{display: 'inline-block',height: '40px',lineHeight: '40px'}}> 热门商品最低点击量：</span>
                    <Form.Item  >
                    {getFieldDecorator('popularlimit', {
                        rules: [{ required: true, message: '请输入' }],
                    })(
                        <Input
                        placeholder="请输入"
                        style={{width:'300px'}}
                        />,
                    )}
                    </Form.Item>
                    <span  style={{display: 'inline-block',height: '40px',lineHeight: '40px'}}> 热门商品展示数量：</span>
                    <Form.Item>
                    {getFieldDecorator('popularquantity', {
                        rules: [{ required: true, message: '请输入' }],
                    })(
                        <Input
                        placeholder="请输入"
                        style={{width:'300px'}}
                        />,
                    )}
                    </Form.Item>
                    <Form.Item  >
                    <Button type="primary" htmlType="submit">
                        提交
                    </Button>
                    <Button onClick={this.cancelEdit}  style={{margin:'0 10px'}}>
                        取消
                    </Button>
                    </Form.Item>
                </Form>
                </div>):''
                }
             <div style={{marginTop:'50px'}}>
                    <span style={{display:'inline-block',width:'40%'}}>热门商品最低点击量：{this.state.minclick}</span>
                    <span style={{display:'inline-block',width:'40%'}}>热门商品展示数量:{this.state.num}</span>
                    <Button type='primary' onClick={this.showEdit}> 修改</Button>
                </div>

                
            </PageHeaderWrapper>
        )
    }
}

export default Hotgoods;