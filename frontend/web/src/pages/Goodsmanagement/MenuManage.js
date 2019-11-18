import React,{ PureComponent,Fragment} from 'react'
import {
    Table,
    Divider,
    Button,
    Pagination,
    Form,
    Modal,
    message,
    Upload,
    Input
} from 'antd'
import router from 'umi/router'
import axios from 'axios'
import qs from 'qs'
import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import styles from './MenuManage.less'


@Form.create()
class MenuManage extends PureComponent{
    state={
        shopid:1,
        listData:[],
        tital:'',
        loading:false,
        addVisible:false,  
        editVisible:false, 
        imageUrl:'', //图像链接
        superid:'',         //当前行的一级菜单id
        menuobj:{},         //当前行的菜单对象信息
        newmenuname:''
    }
   async componentDidMount(){
       await this.setState({
           shopid:localStorage.getItem('shopid')
       })
        this.getList()
    }
    async getList(){
        let res=await axios.post('/api2',qs.stringify({
            token:'Jh2044695',
            cmd:'getAllClass',
            shopid:this.state.shopid
        }))
        console.log(res.data)
        if(res.data.success){
           this.setState({
            listData:res.data.data
        }) 
        }else{
            message.error(res.data.errmsg)
        }
        
    }

    //新建分类弹框
    handleModalVisible=()=>{
        this.setState({
            addVisible:true,
            superid:0,
            menuobj:{}
        })
    }
    //编辑分类
    showEdit=(record)=>{
        console.log(record)
        this.setState({
            editVisible:true,
            menuobj:record
        })
        console.log(this.state.menuobj)
    }
    //确认编辑
    handleEditOk=(e)=>{
        let that=this
        e.preventDefault()
        this.props.form.validateFields((err,values)=>{
            if(!err){
                // console.log(values)
                axios.post('/api2',qs.stringify({
                    cmd:'updateMenu',
                    token:'Jh2044695',
                    menuid:this.state.menuobj.id,
                    menuname:values.menuname
                })).then(res=>{
                        console.log(res.data)
                    if(res.data.success){
                        message.success(res.data.errmsg)
                        that.getList()
                    }else{
                        message.error(res.data.errmsg)
                    }
                    this.setState({
                        editVisible:false
                    })
                })
            }
        })
    }
    //取消编辑
    handleEditCancel=()=>{
        this.setState({
            editVisible:false
        })
    }
    //删除一级分类
    handleDeleteProject=(record)=>{
        console.log(record)
        Modal.confirm({
            title: '删除确认',
            content: '确认删除该分类吗？',
            okText: '确认',
            cancelText: '取消',
            onOk(){
                axios.post('/api2',qs.stringify({
                    cmd:'delMenu',
                    token:'Jh2044695',
                    menuid:record.id,
                    superid:record.superid
            })).then(res=>{
                if(res.data.success){
                    message.success(res.data.errmsg)
                    this.getList()
                }else{
                    message.error(res.data.errmsg)
                }
                })
            },
            onCancel(){}
          });
    }
    //删除二级分类
    handleDeleteGroup=(record)=>{
        console.log(record)
        let that=this
        Modal.confirm({
            title: '删除确认',
            content: '确认删除该分类吗？',
            okText: '确认',
            cancelText: '取消',
            onOk(){
                axios.post('/api2',qs.stringify({
                    cmd:'delMenu',
                    token:'Jh2044695',
                    menuid:record.id,
                    superid:record.superid
            })).then(res=>{
                console.log(res.data)
                if(res.data.success){
                    message.success(res.data.errmsg)
                    that.getList()
                }else{
                    message.error(res.data.errmsg)
                }
                })
            },
            onCancel(){}
          });
    }
    //添加子分类
    showAddChild=(record)=>{
        this.setState({
            addVisible:true,
            superid:record.id
        })
    }
    //确认添加
    handleAddOk=()=>{
        this.setState({
            loading:true
        })
        axios.post('/api2',qs.stringify({
            cmd:'insertMenu',
            token:'Jh2044695',
            shopid:this.state.shopid,
            superid:this.state.superid,
            menuname:this.state.newmenuname
        })).then(res=>{
            console.log(res.data)
            if(res.data.success){
              message.success(res.data.errmsg)  
              this.setState({
                    loading:false,
                    addVisible:false,
                    newmenuname:''
                })
                this.getList()
            }else{
                message.error(res.data.errmsg)
                this.setState({
                    loading:false
                })
            }       
            
        })
        
        // e.preventDefault()
        // this.props.form.validateFields((err,values)=>{
        //     console.log(values)
        //     if(!err){
        //         // console.log(values)
               
        //     }else{
        //         this.setState({
        //             loading:false
        //         })
        //     }
        // })
    }
    //取消添加
    handleCancel=()=>{
        this.setState({
            addVisible:false
        })
    }
    handleInput=(event)=>{
        console.log(event.target.value)
        this.setState({
            newmenuname:event.target.value
        })
    }
    render(){
        const {getFieldDecorator}=this.props.form
        const formLayout = {
            labelCol: {
                span: 6
            },
            wrapperCol: {
                span: 16
            }
        };
        const columns = [
            {
                title: '名称',
                dataIndex: 'menuname',
                key: 'menuname',
            },
            {
                title: '操作',
                render: (text, record) => (
                    <Fragment>
                        <span>
                            <a target="_blank" onClick={() => this.showEdit(record)}>编辑</a> &nbsp;&nbsp;&nbsp;
            {record.superid==0 ? <span> <a target="_blank" onClick={() => this.handleDeleteProject(record)}>删除</a>&nbsp;&nbsp;&nbsp;<a target="_blank" onClick={() => { this.showAddChild(record) }}>添加子分类 </a> &nbsp;&nbsp;&nbsp;
          </span>
                                : <a target="_blank" onClick={() => this.handleDeleteGroup(record)}>删除</a>
                            }
                        </span>
                    </Fragment>
                ),
            },
        ];
       
        return (
            <Fragment>
                <PageHeaderWrapper title='菜单管理'>
                    <div className={styles.tableList}>
                    <div className={styles.tableListOperator}>
                        <Button icon="plus" type="primary" onClick={() => this.handleModalVisible(true)} >
                            新建
                </Button>
                    </div>
                    <Table
                        dataSource={this.state.listData}
                        columns={columns}
                        pagination={false}
                        rowKey={row=>row.id}
                    />
                    </div>
                </PageHeaderWrapper>
                <Modal 
                    title='添加分类'
                    visible={this.state.addVisible}
                    onOk={this.handleAddOk}
                    onCancel={this.handleCancel}
                    confirmLoading={this.state.loading}
                >
                   <span>分类名称</span> <Input placeholder='请输入分类名称' onChange={event=>this.handleInput(event)} value={this.state.newmenuname}/>
                    {/* <Form {...formLayout} onSubmit={this.handleAddOk} > 
                        <Form.Item label='分类名称'>
                            {getFieldDecorator('newmenuname',{
                                rules:[{required:true,message:'请输入菜单名称！'}]
                            })(<Input placeholder='请输入菜单名称'/>)}
                        </Form.Item>
                    </Form> */}
                </Modal>
                <Modal 
                    title='编辑分类'
                    visible={this.state.editVisible}
                    onOk={this.handleEditOk}
                    onCancel={this.handleEditCancel}
                >
                    <Form {...formLayout} onSubmit={this.handleEditOk} > 
                        <Form.Item label='分类名称'>
                            {getFieldDecorator('menuname',{
                                initialValue:this.state.menuobj.menuname,
                                rules:[{required:true,message:'请输入菜单名称！'}]
                            })(<Input placeholder='请输入菜单名称'/>)}
                        </Form.Item>
                    </Form>
                </Modal>
            </Fragment>
        )
    }
}

export default MenuManage;