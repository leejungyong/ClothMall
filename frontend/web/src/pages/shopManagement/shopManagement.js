import React, {
  PureComponent, Fragment
} from 'react';
import {
  connect
} from 'dva';
import router from 'umi/router';
import qs from "qs";
import axios from "axios";
import { Table, Divider, Tag, Pagination, Button, Form, Row, Col,message, Select, Input, Icon } from 'antd';
import PageHeaderWrapper from '@/components/PageHeaderWrapper';
import styles from './ShopManagement.less';
const FormItem = Form.Item;
const { Option } = Select;
@Form.create()
class shopManagement extends PureComponent {
  state = {
    listData: [],
    total: 0,//列表总数
    pageSize: 10,//当前翻页条数
    pageNo: 1,//当前页码
    state: "",//当前选择的店铺状态
    shopname: "" //搜索点名名称
  };
  componentDidMount() {
    this.getList();
  }
  componentWillUnmount() { }
  async getList() {
    let res = await axios.post('/api2', qs.stringify({
      shopname: this.state.shopname,
      token: 'Jh2044695',
      cmd: 'getShopList',
      state: this.state.state,
      pageNo: this.state.pageNo,
      pageSize: this.state.pageSize
    }));
    console.log("xxx", res);
    this.setState({
      listData: res.data.data.list,
      total: res.data.data.listCount
    })
  }
  //查询
  handleSearch = e => {
    e.preventDefault();
    const { dispatch, form } = this.props;
    form.validateFields((err, fieldsValue) => {
      if (err) return;
      console.log(fieldsValue)
      this.setState({
        shopname: fieldsValue.shopname,
        state: fieldsValue.state
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
            <FormItem label="店铺名称">
              {getFieldDecorator('shopname')(<Input placeholder="请输入" />)}
            </FormItem>
          </Col>
          <Col md={8} sm={24}>
            <FormItem label="店铺状态">
              {getFieldDecorator('state')(
                <Select placeholder="请选择" style={{ width: '100%' }}>
                  <Option value="1">已生效</Option>
                  <Option value="2">未生效</Option>
                  <Option value="3">已停用</Option>
                  <Option value="4">已删除</Option>
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
    //跳转
    router.push({
      pathname: '/shopmanagement/newShop',
    });
  };
  editShop(record){
  //跳转
  router.push({
    pathname: '/shopmanagement/editShop',
    query:{
      id:record.shopid
    }
  });
  }
  editImg(record){
    //跳转
    router.push({
      pathname: '/shopmanagement/editImg',
      query:{
        id:record.shopid
      }
    });
    }
   stopShop=async(record)=>{
    const responseDelete = await axios.post('/api2', qs.stringify({
      token: 'Jh2044695',
      cmd: 'updateShopStatusToDisable',
      shopid:record.shopid
  }));
  if (responseDelete.data.success) {
      message.success('已停用');
      //更新列表
      this.getList();
  }
  //添加失败
  !responseDelete.data.success && message.error(`${responseDelete.data.errmsg}`) ;
   }
  deleteShop=async(record)=>{
    const responseDelete = await axios.post('/api2', qs.stringify({
        token: 'Jh2044695',
        cmd: 'delShop',
        shopid:record.shopid
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
    router.push({
      pathname: '/shopmanagement/configAccount',
      query: {
        shopid: item.shopid,
      },
    });
  }
  //配置机器
  configMachine(item){
    router.push({
      pathname: '/shopmanagement/ConfigMachine',
      query: {
        shopid: item.shopid,
      },
    });
  }
  render() {
    const stateEum = {
      '1': '已生效',
      '2': '未生效',
      '3': '已停用',
      '4': '已删除'
    }
    const columns = [
      {
        title: 'URL标识',
        dataIndex: 'shopurl',
        key: 'shopurl',
      },
      {
        title: '店铺状态',
        dataIndex: 'state',
        key: 'state',
        render: (text, record) => {
          return stateEum[text]
        }
      },
      {
        title: '店铺名称',
        dataIndex: 'shopname',
        key: 'shopname',
      }, {
        title: '产品数量',
        dataIndex: 'goodsnum',
        key: 'goodsnum',
      }, {
        title: '总访问量',
        dataIndex: 'visitnum',
        key: 'visitnum',
      }, {
        title: '注册用户数',
        dataIndex: 'managenum',
        key: 'managenum',
      },
      //  {
      //   title: '产品总访问量',
      //   dataIndex: 'visitnum',
      //   key: 'visitnum',
      // },
      {
        title: '机器数量',
        dataIndex: 'machinenum',
        key: 'machinenum',
      }, {
        title: '操作',
        render: (text, record) => (
          <Fragment>
            <a onClick={() => this.editShop(record)}>修改店铺信息</a>
            <Divider type="vertical" />
            <a onClick={() => this.editImg(record)}>修改店铺广告图</a>
            <Divider type="vertical" />
            <a onClick={() => this.configAccount(record)}>配置账号</a>
            <Divider type="vertical" />
            <a onClick={() => this.configMachine(record)}>配置机器</a>
            <Divider type="vertical" />
            <a onClick={() => this.deleteShop(record)}>删除</a>
            <Divider type="vertical" />
            <a onClick={() => this.stopShop(record)}>停用</a>
          </Fragment>
        ),
      },
    ];


    return <PageHeaderWrapper title="店铺管理">
      <div>
        <div className={styles.tableList}>
          <div className={styles.tableListForm}>{this.renderSimpleForm()}</div>
          <div className={styles.tableListOperator}>
            <Button icon="plus" type="primary" onClick={() => this.handleModalVisible()} >
              新建
              </Button>
          </div>
        </div>
        <Table
          dataSource={this.state.listData}
          columns={columns}
          pagination={false}
          rowKey={row => row.shopid}
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
  }
}
export default shopManagement;
