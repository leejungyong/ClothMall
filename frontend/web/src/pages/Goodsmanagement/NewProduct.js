import React, { PureComponent } from 'react'
import { connect } from 'dva'
import axios from 'axios'
import qs from 'qs'
import {
    Input,
    Form,
    Icon,
    Button,
    Cascader,
    message
} from 'antd'

import PageHeaderWrapper from '@/components/PageHeaderWrapper';

import styles from './NewProduct.less'
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
class NewProduct extends PureComponent {
    state = {
        menuoptions: [
            {
                value: '1',
                label: '墙纸类',
                children: [
                    {
                        value: '4',
                        label: '中式风格'
                    },
                ],
            },
        ],
        machineoptions: [],
        newproid: '',
        shopid: 1
    }

    async  componentDidMount() {
        await this.setState({
            shopid: localStorage.getItem('shopid')
        })
        this.getMenuAndmachine()
    }

    //获取菜单和机器信息
    getMenuAndmachine() {
        axios.post('/api2', qs.stringify({ cmd: 'getMenuAndmachine', token: 'Jh2044695', shopid: this.state.shopid })).then(res => {
            console.log(res.data)
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

    //确认新增产品
    handleSubmit = e => {
        e.preventDefault();
        this.props.form.validateFields((err, values) => {
            console.log(err)
            if (!err) {
                console.log(values);
                axios.post('/api2', qs.stringify({
                    cmd: 'insertGoods',
                    token: 'Jh2044695',
                    shopid: this.state.shopid,
                    menuid: values.menuid[1],
                    goodsname: values.goodsname,
                    brand: values.brand,
                    style: values.style,
                    material: values.material,
                    unit: values.unit,
                    madein: values.madein,
                    price: values.price,
                    width: values.width,
                    height: values.height,
                    machineid: parseInt(values.machine[0]),
                    slotnum: values.machine[1]
                })).then(res => {
                    if (res.data.success) {
                        console.log(res.data)
                        message.success(res.data.errmsg)
                        this.setState({
                            newproid: res.data.data.id
                        })

                        router.push({
                            pathname: '/goodsmanagement/productmanage/list',
                            query: {}
                        })
                    }
                })
            }
        });
    };

    render() {
        const { getFieldDecorator } = this.props.form;
        return (
            <PageHeaderWrapper title="新增产品">
                <Form {...formItemLayout} onSubmit={this.handleSubmit} className="" >
                    <Form.Item label='名称:' >
                        {getFieldDecorator('goodsname', {
                            rules: [{ required: true, message: '请输入产品名称!' }],
                        })(
                            <Input
                                placeholder="请输入产品名称"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='品牌:'>
                        {getFieldDecorator('brand', {
                            rules: [{ required: true, message: '请输入产品品牌!' }],
                        })(
                            <Input
                                placeholder="请输入产品品牌"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='风格:'>
                        {getFieldDecorator('style', {
                            rules: [{ required: true, message: '请输入产品风格!' }],
                        })(
                            <Input
                                placeholder="请输入产品风格"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='材质:'>
                        {getFieldDecorator('material', {
                            rules: [{ required: true, message: '请输入产品材质!' }],
                        })(
                            <Input
                                placeholder="请输入产品材质"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='规格:'>
                        {getFieldDecorator('unit', {
                            rules: [{ required: true, message: '请输入产品规格!' }],
                        })(
                            <Input
                                placeholder="请输入产品规格"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='产地:'>
                        {getFieldDecorator('madein', {
                            rules: [{ required: true, message: '请输入产品产地!' }],
                        })(
                            <Input
                                placeholder="请输入产品产地"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='价格:'>
                        {getFieldDecorator('price', {
                            rules: [{ required: true, message: '请输入产品价格!' }],
                        })(
                            <Input
                                placeholder="请输入产品价格"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='宽度:'>
                        {getFieldDecorator('width', {
                            rules: [{ required: true, message: '请输入产品宽度!' }],
                        })(
                            <Input
                                placeholder="请输入产品宽度"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label='高度:'>
                        {getFieldDecorator('height', {
                            rules: [{ required: true, message: '请输入高度!' }],
                        })(
                            <Input
                                placeholder="请输入产品高度"
                            />,
                        )}
                    </Form.Item>
                    <Form.Item label="所属菜单分类">
                        {getFieldDecorator('menuid', {
                            // initialValue: ['zhejiang', 'hangzhou', 'xihu'],
                            rules: [
                                { type: 'array', required: true, message: '请选择分类' },
                            ],
                        })(<Cascader options={this.state.menuoptions} />)}
                    </Form.Item>
                    <Form.Item label="机器位置">
                        {getFieldDecorator('machine', {
                            // initialValue: ['zhejiang', 'hangzhou', 'xihu'],
                            rules: [
                                { type: 'array', required: true, message: '请选择机器' },
                            ],
                        })(<Cascader options={this.state.machineoptions} />)}
                    </Form.Item>
                    <Form.Item {...tailFormItemLayout} >
                        <Button type="primary" htmlType="submit">
                            提交
            </Button>
                    </Form.Item>

                </Form>

            </PageHeaderWrapper>

        );
    }


}


export default NewProduct