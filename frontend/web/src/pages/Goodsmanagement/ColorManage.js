import React, { PureComponent } from 'react'
import {
    Button,
    Table,
    Divider,
    Pagination,
    message,
    Modal,
    Form,
    Input,
    Icon,
    Upload
} from 'antd'

import axios from 'axios'
import qs from 'qs'
import { router } from 'umi'


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
@Form.create()
class ColorManage extends PureComponent {
    state = {
        goodsid: '',
        page: 1,
        pageSize: 10,
        list: [],
        total: 0,
        visible: false,
        ModalText: '确认删除该颜色吗？',
        confirmLoading: false,
        colorid: '',
        isShowAddModal:false,   //是否显示新增弹框
        colorLoading:false,
        isShowEditModal:false,  //是否显示编辑弹框
        colorinfo:{},       //某一行颜色的信息
        imageUrl:'',               //图像链接
        loading:false,
        pic:'',
        compresspic:'',
        picarr:[]
    }
    async componentDidMount() {
        let goodsid = this.props.location.query.goodsid
        console.log(goodsid)
        await this.setState({
            goodsid: goodsid
        })
        this.getList()
    }

    //获取颜色列表
    getList = () => {
        axios.post('/api2', qs.stringify({
            cmd: 'selectColorList',
            token: 'Jh2044695',
            goodsid: this.state.goodsid,
            page: this.state.page,
            count: this.state.pageSize
        })).then(res => {
            console.log(res.data.data)
            if (res.data.success) {
                this.setState({
                    total: res.data.data.total,
                    list: res.data.data.colorlist
                })
            }
        })
    }
    //切换每页的数量
    onShowSizeChange = (current, pageSize) => {
        this.setState({
            pageSize: pageSize,
            page: current
        }, () => {
            this.getList();
        })
    };
    //切换到第几页
    changePage = (page, pageSize) => {
        this.setState({
            pageSize: pageSize,
            page: page
        }, () => {
            this.getList();
        })
    }
    //显示删除弹框
    showDeleteModal = (record) => {
        this.setState({
            visible: true,
            colorid: record.id
        })
    }
    //确认删除
    handleOk = () => {
        this.setState({
            confirmLoading: true
        });
        axios.post('/api2', qs.stringify({ cmd: 'deleteColor', token: 'Jh2044695', colorid: this.state.colorid })).then(res => {
            console.log(res.data)
            if (res.data.success) {
                this.setState({
                    visible: false,
                    confirmLoading: false,
                });
                message.success('删除成功！')
                this.getList()
            } else {
                message.error(res.data.errmsg)
            }
        })
    }
    //取消删除
    handleCancel = () => {
        this.setState({
            visible: false
        })
    }
    //编辑
    edit = (record) => {
        
        this.setState({
            isShowEditModal:true,
            colorinfo:record,
            imageUrl:record.compresspic
        })
    }
    //点击新增按钮
    showAddModal=()=>{
        this.setState({
            isShowAddModal:true
        })
    }
    //确认新增颜色
    sureAddColor=(e)=>{
        e.preventDefault()
        this.props.form.validateFields((err,values)=>{
            if(!err){
                this.setState({
                    colorLoading:true
                })
                axios.post('/api2',qs.stringify({
                    cmd:'insertColor',
                    token:'Jh2044695',
                    goodsid:this.state.goodsid,
                    model:values.model,
                    pic:this.state.pic,
                    compresspic:this.state.compresspic
                })).then(res=>{
                    console.log(res.data)
                    if(res.data.success){
                        this.setState({
                            isShowAddModal:false,
                            colorLoading:false
                        })
                        message.success(res.data.errmsg)
                        this.getList()
                    }else{
                        message.error(res.data.errmsg)
                    }
                })
            }
        })
        
    }
    //取消新增
    cancelAdd=()=>{
        this.setState({
            isShowAddModal:false
        })
    }
    //确认编辑颜色
    sureEditColor=(e)=>{
        e.preventDefault()
        this.props.form.validateFields((err,values)=>{
            if(!err){
                this.setState({
                    colorLoading:true
                })
                axios.post('/api2',qs.stringify({
                    cmd:'updateColor',
                    token:'Jh2044695',
                    colorid:this.state.colorinfo.id,
                    model:values.model,
                    pic:this.state.pic,
                    compresspic:this.state.compresspic
                })).then(res=>{
                    console.log(res.data)
                    if(res.data.success){
                        this.setState({
                            isShowEditModal:false,
                            colorLoading:false
                        })
                        message.success(res.data.errmsg)
                        this.getList()
                    }else{
                        message.error(res.data.errmsg)
                    }
                })
            }
        })
        
    }
    //取消编辑
    cancelEdit=()=>{
        this.setState({
            isShowEditModal:false
        })
    }
    getBase64(img, callback) {
        const reader = new FileReader();
        reader.addEventListener('load', () => callback(reader.result));
        reader.readAsDataURL(img);
    }
    beforeUpload(file) {
        const isJpgOrPng = file.type === 'image/jpeg' || file.type === 'image/png';
        if (!isJpgOrPng) {
            message.error('You can only upload JPG/PNG file!');
        }
        const isLt2M = file.size / 1024 / 1024 < 2;
        if (!isLt2M) {
            message.error('Image must smaller than 2MB!');
        }
        return isJpgOrPng && isLt2M;
    }

      handleChange = info => {
        if (info.file.status === 'uploading') {
          this.setState({ loading: true });
          return;
        }
        if (info.file.status === 'done') {
          // Get this url from response in real world.
          
          this.getBase64(info.file.originFileObj, imageUrl =>
            this.setState({
              imageUrl:imageUrl,
              loading: false,
              pic:info.file.response.data[0].image,
              compresspic:info.file.response.data[0].thumb
            }),
          );
        }
      };
    render() {
        const uploadButton = (
            <div>
              <Icon type={this.state.loading ? 'loading' : 'plus'} />
              <div className="ant-upload-text">Upload</div>
            </div>
          );
          const { imageUrl } = this.state;
        const { visible, confirmLoading, ModalText ,isShowAddModal,colorLoading,isShowEditModal} = this.state
        const {getFieldDecorator}=this.props.form
        const columns = [
            {
                title: '型号',
                dataIndex: 'model',
                key: 'model'
            },
            {
                title: '缩略图',
                dataIndex: 'compresspic',
                render: (val) => <img style={{ width: '40px', height: '40px' }} src={val} />
            },
            {
                title: '操作',
                key: 'action',
                render: (text, record) => (
                    <span>
                        <a onClick={() => this.showDeleteModal(record)}>删除</a>
                        <Divider type="vertical" />
                        <a onClick={() => this.edit(record)}>编辑</a>
                    </span>
                )
            }
        ]

        return (
            <div>
                <Button type='primary' onClick={this.showAddModal}> 新增颜色</Button>
                <Table
                    columns={columns}
                    dataSource={this.state.list}
                    pagination={false}
                    style={{margin:'10px 0'}}
                />
                <Pagination
                    total={this.state.total}
                    defaultCurrent={1}
                    showTotal={total => `共${total}条`}
                    showSizeChanger={true}
                    onShowSizeChange={this.onShowSizeChange}
                    showQuickJumper
                    onChange={this.changePage}
                    style={{ margin: '10px 0', textAlign: 'right' }}
                />
                <Modal
                    title="删除确认"
                    visible={visible}
                    onOk={this.handleOk}
                    confirmLoading={confirmLoading}
                    onCancel={this.handleCancel}
                >
                    <p>{ModalText}</p>
                </Modal>
                <Modal
                    title="新增颜色"
                    visible={isShowAddModal}
                    onOk={this.sureAddColor}
                    confirmLoading={colorLoading}
                    onCancel={this.cancelAdd}
                >
                    <Form {...formItemLayout} onSubmit={this.sureAddColor}>
                        <Form.Item label='型号'>
                            {getFieldDecorator('model',{
                                rules:[{required:true,message:'请输入型号'}]
                            })(
                                <Input placeholder='请输入型号'/>
                            )
                            }
                        </Form.Item>
                        上传图片
                        <Upload
                        data={{token:'Jh2044695',
                        module:'goodspic',
                        kresize:'{"key": "k", "Thumbonly":0, "Width":100, "Height":100}'
                    }}
                        name="k"
                        listType="picture-card"
                        className="avatar-uploader"
                        showUploadList={false}
                        action="http://localhost:8000/multiupload"
                        beforeUpload={this.beforeUpload}
                        onChange={this.handleChange}
                        
                    >
                        {imageUrl ? <img src={imageUrl} alt="avatar" style={{ width: '100%' }} /> : uploadButton}
                    </Upload>
                    </Form>
                </Modal>
                <Modal
                    title="编辑颜色"
                    visible={isShowEditModal}
                    onOk={this.sureEditColor}
                    confirmLoading={colorLoading}
                    onCancel={this.cancelEdit}
                >
                    <Form {...formItemLayout} onSubmit={this.sureEditColor}>
                        <Form.Item label='型号'>
                            {getFieldDecorator('model',{
                                initialValue:this.state.colorinfo.model,
                                rules:[{required:true,message:'请输入型号'}]
                            })(
                                <Input placeholder='请输入型号'/>
                            )
                            }
                        </Form.Item>
                        上传图片
                    <Upload
                        data={{token:'Jh2044695',
                        module:'goodspic',
                        kresize:'{"key": "k", "Thumbonly":0, "Width":100, "Height":100}'
                    }}
                        name="k"
                        listType="picture-card"
                        className="avatar-uploader"
                        showUploadList={false}
                        action="http://localhost:8000/multiupload"
                        beforeUpload={this.beforeUpload}
                        onChange={this.handleChange}
                    >
                        {imageUrl ? <img src={imageUrl} alt="avatar" style={{ width: '100%' }} /> : uploadButton}
                    </Upload>
                    </Form>
                </Modal>
            </div>
        )
    }
}

export default ColorManage