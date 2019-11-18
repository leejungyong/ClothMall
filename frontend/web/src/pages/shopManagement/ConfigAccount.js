import React, {
    PureComponent, Fragment
  } from 'react';
  import {
    connect
  } from 'dva';
  import router from 'umi/router';
  import qs from "qs";
  import axios from "axios";
  import { Table, Divider, Tag, Pagination, Modal,Button, Form, Row, Col, Select, Input, Icon ,message} from 'antd';
  import PageHeaderWrapper from '@/components/PageHeaderWrapper';
  import styles from './ConfigAccount.less';
  const FormItem = Form.Item;
  const { Option } = Select;
  @Form.create()
  class ConfigAccount extends PureComponent {
    state = {
      listData: [],
      total: 0,//列表总数
      pageSize: 10,//当前翻页条数
      pageNo: 1,//当前页码
      state: "",//当前选择的店铺状态
      shopname: "" ,//搜索点名名称
      visible:"",//是否展示
      isEdit:'',//是否是编辑
      editData:""//编辑数据

    };
    componentDidMount() {
      console.log(this.props.location.query.shopid)
      this.getList();
    }
    componentWillUnmount() { }
    async getList() {
      let res = await axios.post('/api2', qs.stringify({
        shopid:this.props.location.query.shopid,
        token: 'Jh2044695',
        cmd: 'getShopManagerList',
        pageNo: this.state.pageNo,
        pageSize: this.state.pageSize
      }));
      console.log("xxx", res);
      this.setState({
        listData: res.data.data.list,
        total: res.data.data.listCount
      })
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
     //删除
     handleDelete=async(record)=>{
      const responseDelete = await axios.post('/api2', qs.stringify({
          token: 'Jh2044695',
          cmd: 'delShopManager',
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
    handleModalVisible = () => {
      this.setState({
        visible: true,
        isEdit:false
      });
    };
      //新增确定
    handleAddOk = () => {
        const { form: { validateFields } } = this.props;
        validateFields(['name','password','remark'], (err, values) => {
            if (err) return;
            this.handleAdd(values);
        });
    }
      //新增和编辑
      handleAdd = async (fields) => {
        //编辑
        if(this.state.isEdit){
          const responseEdit = await axios.post('/api2', qs.stringify({
            token: 'Jh2044695',
            cmd: 'updateShopManagerInfo',
            id:this.state.editData.id,
            password:fields.password,
            remark:fields.remark,
        }));
        if (responseEdit.data.success) {
            message.success('编辑成功');
            this.setState({ visible: false});
            //更新列表
            this.getList();
        }
        //添加失败
        !responseEdit.data.success && message.error(`${responseEdit.data.errmsg}`) && this.setState({ visible: false});
        }
        if(!this.state.isEdit){
            const responseAdd = await axios.post('/api2', qs.stringify({
                token: 'Jh2044695',
                cmd: 'insertNewShopManager',
                phonenum:fields.name,
                password:fields.password,
                remark:fields.remark,
                shopid:this.props.location.query.shopid

            }));
            if (responseAdd.data.success) {
                message.success('添加成功');
                this.setState({ visible: false});
                //更新列表
                this.getList();
            }
            //添加失败
            !responseAdd.data.success && message.error(`${responseAdd.data.errmsg}`) && this.setState({ visible: false});
        }
    }
      //取消新增
      handleAddCancel = () => {
        this.setState({
        visible: false,
        });
    }
    //展示编辑
    showEdit(record){
     this.setState({visible:true,
        editData:record,
        isEdit:true})
    }
    render() {
      const { getFieldDecorator } = this.props.form;
    
      const columns = [
        {
          title: '用户名',
          dataIndex: 'phonenum',
          key: 'phonenum',
        },
        {
          title: '密码',
          dataIndex: 'password',
          key: 'password',
        },
        {
          title: '备注',
          dataIndex: 'remark',
          key: 'remark',
        },{
          title: '操作',
          render: (text, record) => (
            <Fragment>
              <a onClick={() => this.handleDelete(record)}>删除</a>
              <Divider type="vertical" />
              <a onClick={() => this.showEdit(record)}>编辑</a>
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
      return  <Fragment><PageHeaderWrapper title="配置账号">
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
           <Pagination
            total={this.state.total}
            defaultCurrent={1}
            showTotal={total => `共 ${total} 条`}
            showSizeChanger={true}
            onShowSizeChange={this.onShowSizeChange}
            showQuickJumper
            onChange={this.onChange}
            style={{ margin: '10px 0', textAlign: 'right' }}
          />
        </div>
      </PageHeaderWrapper>
       <Modal
       destroyOnClose={true}
       title={`${this.state.isEdit ? '编辑账号' : '新增账号'}`}
       visible={(this.state.visible)}
       onOk={this.handleAddOk}
       onCancel={this.handleAddCancel}
       maskClosable={false}
       bodyStyle={{ 'color': 'red' }}
   >
       <Form>
           <FormItem label={'用户名'} {...formLayout}>
               {
                   getFieldDecorator('name',{initialValue:this.state.isEdit?this.state.editData.phonenum:"", rules: [{
                       required: true,
                       message: '用户名不能为空',
                   }]})(
                     this.state.isEdit?<Input disabled placeholder={'请输入用户名'} />:
                     <Input  placeholder={'请输入用户名'} />
                   )
               }
           </FormItem>
           <FormItem label={'密码'} {...formLayout}>
               {
                   getFieldDecorator('password',{initialValue:this.state.isEdit?this.state.editData.password:"", rules: [{
                       required: true,
                       message: '密码不能为空',
                   }]})(
                       <Input placeholder={'请输入密码'} />
                   )
               }
           </FormItem>
           <FormItem label={'备注'} {...formLayout}>
               {
                   getFieldDecorator('remark',{initialValue:this.state.isEdit?this.state.editData.remark:"" })(
                       <Input placeholder={'请输入备注'} />
                   )
               }
           </FormItem>
       </Form>
   </Modal>
     </Fragment>
    }
  }
  export default ConfigAccount;
  