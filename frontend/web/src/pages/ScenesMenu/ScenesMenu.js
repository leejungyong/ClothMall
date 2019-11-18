import React, {
    PureComponent, Fragment
} from 'react';
import {
    connect
} from 'dva';
import router from 'umi/router';
import qs from "qs";
import axios from "axios";
import { Table, Divider, Tag, Pagination, Button, Form, Row, Col, Select, Input, Icon, Modal, Upload, message } from 'antd';
import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import styles from './ScenesMenu.less';
const FormItem = Form.Item;
const { Option } = Select;
@Form.create()
class ScenesMenu extends PureComponent {
    state = {
        listData: [],
        isAddParent: false,//是否添加父
        addVisible: false,//添加弹窗
        loading: false,
        imageUrl:'',//图像链接
        picParams:"",//传给后端的链接
        parentId:"",//父元素的id
        isEdit:false,//是否是编辑
        editData:{},//编辑数据
    };
    componentDidMount() {
        this.getList();
    }
    componentWillUnmount() { }
    async getList() {
        let res = await axios.post('/api2', qs.stringify({
            token: 'Jh2044695',
            cmd: 'getScenesMenu',
        }));
        console.log("xxx", res);
        this.setState({
            listData: res.data.data.list,
            total: res.data.data.listCount,
        })
    }

    //新增确定
    handleAddOk = () => {
        const { form: { validateFields } } = this.props;
        validateFields(['nameadd'], (err, values) => {
            if (err) return;
            this.handleAdd(values);
        });
    }
    //新增和编辑
    handleAdd = async (fields) => {
        //编辑
        if(this.state.isEdit){
            const responseAdd = await axios.post('/api2', qs.stringify({
                token: 'Jh2044695',
                cmd: 'updateScenesMenu',
                id:this.state.editData.id,
                superid:this.state.editData.superid||0,
                pic:this.state.picParams||this.state.editData.pic,
                name:fields.nameadd,
            }));
            if (responseAdd.data.success) {
                message.success('编辑成功');
                this.setState({ addVisible: false,parentId:"" });
                //更新列表
                this.getList();
            }
            //添加失败
            !responseAdd.data.success && message.error(`${responseAdd.data.errmsg}`) && this.setState({ addVisible: false,parentId:"" });

        }
        if(!this.state.isEdit){
            const responseAdd = await axios.post('/api2', qs.stringify({
                token: 'Jh2044695',
                cmd: 'insertScenesMenu',
                superid:this.state.parentId||0,
                name:fields.nameadd,
                pic:this.state.picParams
            }));
            if (responseAdd.data.success) {
                message.success('添加成功');
                this.setState({ addVisible: false,parentId:"" });
                //更新列表
                this.getList();
            }
            //添加失败
            !responseAdd.data.success && message.error(`${responseAdd.data.errmsg}`) && this.setState({ addVisible: false,parentId:"" });
        }
    }
    //取消新增
    handleAddCancel = () => {
        this.setState({
            addVisible: false,
        });
    }
    //页数改变
    onShowSizeChange = (current, pageSize) => {
        this.setState({
            pageSize: pageSize,
            pageNo: current
        }, () => {
            this.getList();
        })
    };
    //切换第一页第二页
    onChange = (page, pageSize) => {
        this.setState({
            rowpageSizes: pageSize,
            pageNo: page
        }, () => {
            this.getList();
        })
    }
   //新建父
    handleModalVisible = flag => {
        this.setState({
            isEdit:false,
            addVisible: true,
            isAddParent: true,
            imageUrl:""
        }
        );

    };
    //新增子
    showAddChild = (record) => {
        this.setState({
            isEdit:false,
            addVisible: true,
            isAddParent: false,
            parentId: record.id,
            imageUrl:""
        }
        );
    }
    //展示编辑
    showEdit(record){
        console.log(record);
     this.setState({addVisible:true,
        editData:record,
        isEdit:true,imageUrl:record.pic})//懒得写都用一个吧
    }
    //删除
     handleDelete=async(record)=>{
        const responseDelete = await axios.post('/api2', qs.stringify({
            token: 'Jh2044695',
            cmd: 'delScenesMenu',
            id:record.id
        }));
        if (responseDelete.data.success) {
            message.success('删除成功');
            //更新列表
            this.getList();
        }
        //添加失败
        !responseDelete.data.success && message.error(`${responseDelete.data.errmsg}`) ;
    }
    //配置账号跳转
    configAccount(item) {
        console.log(item);
        router.push({
            pathname: '/shopmanagement/configAccount',
            query: {
                shopid: item.shopid,
            },
        });
    }
    //配置机器
    configMachine(item) {
        router.push({
            pathname: '/shopmanagement/ConfigMachine',
            query: {
                shopid: item.shopid,
            },
        });
    }
    getBase64(img, callback) {
        const reader = new FileReader();
        reader.addEventListener('load', () => callback(reader.result));
        reader.readAsDataURL(img);
    }
    //上传图片相关
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
              imageUrl,
              loading: false,
              picParams:info.file.response.data[0].image
            }),
          );
        }
      };
    render() {
        const { getFieldDecorator } = this.props.form;
        const uploadButton = (
            <div>
              <Icon type={this.state.loading ? 'loading' : 'plus'} />
              <div className="ant-upload-text">Upload</div>
            </div>
          );
          const { imageUrl } = this.state;
        const stateEum = {
            '1': '已生效',
            '2': '未生效',
            '3': '已停用',
            '4': '已删除'
        }
        const columns = [
            {
                title: '名称',
                dataIndex: 'name',
                key: 'name',
            },
            {
                title: '图片',
                dataIndex: 'pic',
                key: 'pic',
                render: (text, record) => {
                    return <img src={text} style={{ 'width': '60px', 'height': '60px' }}  />
                }
            }, {
                title: '操作',
                render: (text, record) => (
                    <Fragment>
                        <span>
                            <a target="_blank" onClick={() => this.showEdit(record)}>编辑</a> &nbsp;&nbsp;&nbsp;
            {record.superid===0 ? <span> <a target="_blank" onClick={() => this.handleDelete(record)}>删除</a>&nbsp;&nbsp;&nbsp;<a target="_blank" onClick={() => { this.showAddChild(record) }}>添加房间型号</a> &nbsp;&nbsp;&nbsp;
          </span>
                                : <a target="_blank" onClick={() => this.handleDelete(record)}>删除</a>
                            }
                        </span>
                    </Fragment>
                ),
            },
        ];
        const formLayout = {
            labelCol: {
                span: 6
            },
            wrapperCol: {
                span: 16
            }
        };
        return <Fragment><PageHeaderWrapper title="模拟场景菜单配置">
            <div>
                <div className={styles.tableList}>
                    <div className={styles.tableListOperator}>
                        <Button icon="plus" type="primary" onClick={() => this.handleModalVisible(true)} >
                            新建
                </Button>
                    </div>
                </div>
                <Table
                    dataSource={this.state.listData}
                    columns={columns}
                    pagination={false}
                    rowKey={row => row.id}
                />
        </div>
        </PageHeaderWrapper>
        <Modal
                destroyOnClose={true}
                title={`${this.state.isAddParent ? '新增项目' : '新增子项目'}`}
                visible={(this.state.addVisible)}
                onOk={this.handleAddOk}
                onCancel={this.handleAddCancel}
                maskClosable={false}
                bodyStyle={{ 'color': 'red' }}
            >
                <Form>
                    <FormItem label={'名称'} {...formLayout}>
                        {
                            getFieldDecorator('nameadd',{initialValue:this.state.isEdit?this.state.editData.name:"", rules: [{
                                required: true,
                                message: '名称不能为空',
                            }]})(
                                <Input placeholder={'请输入名称'} />
                            )
                        }
                    </FormItem>
                    <span className={styles.tipUpload}><em>*</em>上传图片</span>
                    <div className={styles.UploadContant}>
                    <Upload
                        data={{token:'Jh2044695',
                        module:'goodspic',
                        kresize:'{"key": "k", "Thumbonly":0, "Width":100, "Height":100}'
                    }}
                        name="k"
                        listType="picture-card"
                        className="avatar-uploader"
                        showUploadList={false}
                        action="/multiupload"
                        beforeUpload={this.beforeUpload}
                        onChange={this.handleChange}
                    >
                        {imageUrl ? <img src={imageUrl} alt="avatar" style={{ width: '100%' }} /> : uploadButton}
                    </Upload>
                    </div>
                </Form>
            </Modal>
            
        </Fragment>
        {/* 新增模块结束 */ }
    }
}
export default ScenesMenu;
