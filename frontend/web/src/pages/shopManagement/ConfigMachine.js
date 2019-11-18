import React, {
    PureComponent, Fragment
  } from 'react';
  import {
    connect
  } from 'dva';
  import router from 'umi/router';
  import qs from "qs";
  import axios from "axios";
  import { Table, Divider, Tag, Pagination, Button, Form, Row, Col, Select, Modal,Input, Icon ,message} from 'antd';
  import PageHeaderWrapper from '@/components/PageHeaderWrapper';
  import styles from './ConfigMachine.less';
  const FormItem = Form.Item;
  const { Option } = Select;
  @Form.create()
  class ConfigMachine extends PureComponent {
    state = {
      listData: [],
      total: 0,//列表总数
      pageSize: 10,//当前翻页条数
      pageNo: 1,//当前页码
      state: "",//状态
      machineid: "", //序列号
      visible:false,//
      isEdit:false,//
      editData:{},
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
        cmd: 'getMachineList',
        pageNo: this.state.pageNo,
        pageSize: this.state.pageSize,
        machineid:this.state.machineid,
        state: this.state.state
      }));
      this.setState({
        listData: res.data.data.list,
        total: res.data.data.listCount
      })
    }
      //新增确定
      handleAddOk = () => {
        const { form: { validateFields } } = this.props;
        validateFields(['machineid','slotnum','machineip'], (err, values) => {
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
            cmd: 'updateMachineInfo',
            id:this.state.editData.id,
            machineip:fields.machineip,
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
                cmd: 'insertNewMachine',
                machineid:fields.machineid,
                machineip:fields.machineip,
                slotnum:fields.slotnum,
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
    handleSearch = e => {
      e.preventDefault();
      const { dispatch, form } = this.props;
      form.validateFields(['machineid1','state1'],(err, fieldsValue) => {
        if (err) return;
        console.log(fieldsValue)
        this.setState({
          machineid: fieldsValue.machineid1,
          state: fieldsValue.state1
        }, () => {
          this.getList();
        })
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
  //查询上半部
  renderSimpleForm() {
    const {
      form: { getFieldDecorator },
    } = this.props;
    return (
      <Form onSubmit={this.handleSearch} layout="inline">
        <Row gutter={{ md: 8, lg: 24, xl: 48 }}>
          <Col md={8} sm={24}>
            <FormItem label="机器序列号">
              {getFieldDecorator('machineid1')(<Input placeholder="请输入" />)}
            </FormItem>
          </Col>
          <Col md={8} sm={24}>
            <FormItem label="机器状态">
              {getFieldDecorator('state1')(
                <Select placeholder="请选择" style={{ width: '100%' }}>
                  <Option value="1">有效</Option>
                  <Option value="0">无效</Option>
                  <Option value="2">已删除</Option>
                </Select>
              )}
            </FormItem>
          </Col>
          <Col md={8} sm={24}>
            <span className={styles.submitButtons}>
              <Button type="primary" htmlType="submit">
                查询
              </Button>
            </span>
          </Col>
        </Row>
      </Form>
    );
  };
  handleModalVisible = () => {
    this.setState({
      visible: true,
      isEdit:false
    });
  };
    configAccount(item) {
      console.log(item);
    }
      //展示编辑
      showEdit(record){
        this.setState({visible:true,
           editData:record,
           isEdit:true})
       }
            //删除
     handleDelete=async(record)=>{
      const responseDelete = await axios.post('/api2', qs.stringify({
          token: 'Jh2044695',
          cmd: 'delMachine',
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
    render() {
      const { getFieldDecorator } = this.props.form;
      const stateEum = {
        '1': '有效',
        '2': '已删除',
        '0': '无效',
      }
      const columns = [
        {
          title: '机器序列号',
          dataIndex: 'machineid',
          key: 'machineid',
        },
        {
          title: '机器槽位数量',
          dataIndex: 'slotnum',
          key: 'slotnum',
        },
        {
          title: '运行状态',
          dataIndex: 'runstate',
          key: 'runstate',
          render: (text, record) => {
          return text?"正在运行":"未运行"
          }
        }, {
          title: '联网状态',
          dataIndex: 'netstate',
          key: 'netstate',
          render: (text, record) => {
            return text?"已联网":"未联网"
          }
        },{
          title: '机器IP',
          dataIndex: 'machineip',
          key: 'netstmachineipate',
        },{
          title: '状态',
          dataIndex: 'state',
          key: 'state',
          render: (text, record) => {
            return stateEum[text]
          }
        },{
          title: '操作',
          render: (text, record) => (
            <Fragment>
            
              <a onClick={() => this.handleDelete(record)}>删除</a>
              <Divider type="vertical" />
              <a onClick={() => this.showEdit(record)} vhref="">编辑</a>
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
  
      return <Fragment><PageHeaderWrapper title="配置机器">
      <div>
        <div className={styles.tableList}>
          <div className={styles.tableListForm}>{this.renderSimpleForm()}</div>
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
        title={`${this.state.isEdit ? '编辑机器序列号' : '新增机器序列号'}`}
        visible={(this.state.visible)}
        onOk={this.handleAddOk}
        onCancel={this.handleAddCancel}
        maskClosable={false}
        bodyStyle={{ 'color': 'red' }}
    >
        <Form>
            <FormItem label={'机器序列号'} {...formLayout}>
                {
                    getFieldDecorator('machineid',{initialValue:this.state.isEdit?this.state.editData.machineid:"", rules: [{
                        required: true,
                        message: '机器序列号不能为空',
                    }]})(
                      this.state.isEdit?<Input disabled placeholder={'请输入机器序列号'} />:
                      <Input  placeholder={'请输入机器序列号'} />
                    )
                }
            </FormItem>
            <FormItem label={'机器槽位数量'} {...formLayout}>
                {
                    getFieldDecorator('slotnum',{initialValue:this.state.isEdit?this.state.editData.slotnum:"", rules: [{
                        required: true,
                        message: '机器槽位数量不能为空',
                    }]})(
                      this.state.isEdit?<Input disabled placeholder={'请输入机器槽位数量'} />:<Input  placeholder={'请输入机器槽位数量'} />
                    )
                }
            </FormItem>
            <FormItem label={'机器IP'} {...formLayout}>
                {
                    getFieldDecorator('machineip',{initialValue:this.state.isEdit?this.state.editData.machineip:"", rules: [{
                      required: true,
                      message: '机器IP不能为空',
                  }] })(
                        <Input placeholder={'请输入机器IP'} />
                    )
                }
            </FormItem>
        </Form>
    </Modal>
      </Fragment>
    }
  }
  export default ConfigMachine;
  