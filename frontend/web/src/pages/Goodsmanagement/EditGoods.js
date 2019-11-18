import React,{PureComponent} from 'react'
import {connect} from 'dva'
import axios from 'axios'
import qs from 'qs'
import {
    Input,
    Form,
    Icon,
    Button,
    Cascader ,
    message   
} from 'antd'

import PageHeaderWrapper from '@/components/PageHeaderWrapper';

import { router } from 'umi';

const formItemLayout = {
    labelCol: {
      xs: { span: 0 },
      sm: { span: 3 },
    },
    wrapperCol: {
      xs: { span: 10 },
      sm: { span: 10 },
    },
  };
  const tailFormItemLayout = {
    wrapperCol: {
      xs: {
        span: 24,
        offset: 0,
      },
      sm: {
        span: 16,
        offset: 8,
      },
    },
  };
@Form.create()
class NewProduct extends PureComponent{
    state={
        menuoptions:[],
        machineoptions:[],
        shopid:1,
        goodsid:'',
        goodsinfo:{}
    }

  async  componentDidMount(){
        let goodsid = this.props.location.query.goodsid
       await this.setState({
            goodsid:goodsid,
            shopid:localStorage.getItem('shopid')
        })
        this.getMenuAndmachine()
        this.getInfo()
    }
    //获取商品信息
    getInfo(){
        axios.post('/api2',qs.stringify({cmd:'selectGoodsDetail',token:'Jh2044695',goodsid:this.state.goodsid})).then(res=>{
            if(res.data.success){

                this.setState({
                    goodsinfo:res.data.data
                })
                console.log(this.state.goodsinfo)
            }else{
                message.error(res.data.errmsg)
            }
        })
    }
    //获取菜单和机器信息
    getMenuAndmachine(){
        axios.post('/api2',qs.stringify({cmd:'getMenuAndmachine',token:'Jh2044695',shopid:this.state.shopid})).then(res=>{
           console.log(res.data.data)
           if (res.data.success) {
            let menuarr = res.data.data.menuList,
                machinearr = res.data.data.machineList,
                menuoptions = [],
                machineoptions = []
            if (res.data.data.menuList) {
                for (let i = 0; i < menuarr.length; i++) {
                    let oneobj = {}
                    oneobj.value = menuarr[i].id
                    oneobj.label = menuarr[i].menuname
                    let children = []
                    if (menuarr[i].twoLevelClass) {
                        for (let j = 0; j < menuarr[i].twoLevelClass.length; j++) {
                            let twoobj = {}
                            twoobj.value = menuarr[i].twoLevelClass[j].id
                            twoobj.label = menuarr[i].twoLevelClass[j].menuname
                            children.push(twoobj)
                        }
                    }
                    oneobj.children = children
                    menuoptions.push(oneobj)
                }
            }
            if (res.data.data.machineList) {
                for (let i = 0; i < machinearr.length; i++) {
                    let oneobj = {}
                    oneobj.value = machinearr[i].machineid
                    oneobj.label = machinearr[i].machinename
                    let children = []
                    if (machinearr[i].slotlist) {
                        for (let j = 0; j < machinearr[i].slotlist.length; j++) {
                            let twoobj = {}
                            twoobj.value = machinearr[i].slotlist[j].slotid
                            twoobj.label = machinearr[i].slotlist[j].slotnum
                            children.push(twoobj)
                        }
                    }
                    oneobj.children = children
                    machineoptions.push(oneobj)
                }
            }

            this.setState({
                menuoptions: menuoptions,
                machineoptions: machineoptions
            })
            //    console.log(this.state.)

        }
        })
    }

    //确认编辑产品
    handleSubmit = e => {
        e.preventDefault();
        this.props.form.validateFields((err, values) => {
            console.log(err)
          if (!err) {
            console.log(values);
            axios.post('/api2',qs.stringify({
                cmd:'updateGoods',
                token:'Jh2044695',
                goodsid:this.state.goodsid,
                menuid:values.menuid[1],
                goodsname:values.goodsname,
                brand:values.brand,
                style:values.style,
                material:values.material,
                unit:values.unit,
                madein:values.madein,
                price:values.price,
                width:values.width,
                height:values.height,
                machineid:values.machine[0],
                slotnum:values.machine[1]
            })).then(res=>{
                if(res.data.success){
                    console.log(res.data)
                    message.success(res.data.errmsg)
                    router.push({
                        pathname:'/goodsmanagement/productmanage/list',
                        query:{}
                    })
                }else{
                    message.error(res.data.errmsg)
                }
            })
          }
        });
      };
    
      render() {
        const { getFieldDecorator } = this.props.form;
        const {goodsinfo}=this.state
        return (
            <PageHeaderWrapper title="编辑产品">
    <Form {...formItemLayout} onSubmit={this.handleSubmit} className="" >
                <Form.Item label='名称:' >
                {getFieldDecorator('goodsname', {
                    initialValue:goodsinfo.goodsname,
                    rules: [{ required: true, message: '请输入产品名称!' }],
                })(
                    <Input
                    placeholder="请输入产品名称"
                    />,
                )}
                </Form.Item>
                <Form.Item label='品牌:'>
                {getFieldDecorator('brand', {
                    initialValue:goodsinfo.brand,
                    rules: [{ required: true, message: '请输入产品品牌!' }],
                })(
                    <Input
                    placeholder="请输入产品品牌"
                    />,
                )}
                </Form.Item>
                <Form.Item label='风格:'>
                {getFieldDecorator('style', {
                    initialValue:goodsinfo.style,
                    rules: [{ required: true, message: '请输入产品风格!' }],
                })(
                    <Input
                    placeholder="请输入产品风格"
                    />,
                )}
                </Form.Item>
                <Form.Item label='材质:'>
                {getFieldDecorator('material', {
                    initialValue:goodsinfo.material,
                    rules: [{ required: true, message: '请输入产品材质!' }],
                })(
                    <Input
                    placeholder="请输入产品材质"
                    />,
                )}
                </Form.Item>
                <Form.Item label='规格:'>
                {getFieldDecorator('unit', {
                    initialValue:goodsinfo.unit,
                    rules: [{ required: true, message: '请输入产品规格!' }],
                })(
                    <Input
                    placeholder="请输入产品规格"
                    />,
                )}
                </Form.Item>
                <Form.Item label='产地:'>
                {getFieldDecorator('madein', {
                    initialValue:goodsinfo.madein,
                    rules: [{ required: true, message: '请输入产品产地!' }],
                })(
                    <Input
                    placeholder="请输入产品产地"
                    />,
                )}
                </Form.Item>
                <Form.Item label='价格:'>
                {getFieldDecorator('price', {
                    initialValue:goodsinfo.price,
                    rules: [{ required: true, message: '请输入产品价格!' }],
                })(
                    <Input
                    placeholder="请输入产品价格"
                    />,
                )}
                </Form.Item>
                <Form.Item label='宽度:'>
                {getFieldDecorator('width', {
                    initialValue:goodsinfo.width,
                    rules: [{ required: true, message: '请输入产品宽度!' }],
                })(
                    <Input
                    placeholder="请输入产品宽度"
                    />,
                )}
                </Form.Item>
                <Form.Item label='高度:'>
                {getFieldDecorator('height', {
                    initialValue:goodsinfo.height,
                    rules: [{ required: true, message: '请输入高度!' }],
                })(
                    <Input
                    placeholder="请输入产品高度"
                    />,
                )}
                </Form.Item>
                <Form.Item label="所属菜单分类">
            {getFieldDecorator('menuid', {
                initialValue: [goodsinfo.classoneid,goodsinfo.classtwoid],
                rules: [
                { type: 'array', required: true, message: '请选择分类' },
                ],
            })(<Cascader options={this.state.menuoptions} />)}
            </Form.Item>
            <Form.Item label="机器位置">
            {getFieldDecorator('machine', {
                initialValue: [parseInt(goodsinfo.machineid),parseInt(goodsinfo.machineadder)],
                // initialValue:[1,1],
                rules: [
                { type: 'array', required: true, message: '请选择机器' },
                ],
            })(<Cascader options={this.state.machineoptions} />)}
            </Form.Item>
                <Form.Item {...tailFormItemLayout} >
            <Button type="primary" htmlType="submit">
                确认修改
            </Button>
            </Form.Item>

            </Form>

            </PageHeaderWrapper>
         
        );
      }

    
}


export default NewProduct