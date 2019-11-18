import React,{PureComponent} from 'react'
import {connect} from 'dva'
import router from 'umi/router'
import qs from 'qs'
import axios from 'axios'
import {
  Table,
  Divider,
  Pagination,
  Button,
  Cascader,
  Input,
  Modal,
  message
} from 'antd'
import PageHeaderWrapper from '@/components/PageHeaderWrapper';



class ProductManage extends PureComponent{
  state={
    list:[],
    total:0,
    pageSize:10,
    page:1,
    options:[],
    classone:'1',
    classtwo:'',
    goodsname:'',
    visible:false,
    confirmLoading:false,
    ModalText:'确认删除该商品吗？',
    goodsid:'',       //当前操作的商品的id
    shopid:''
  }

 async componentDidMount(){
    // console.log(localStorage.getItem('shopid'))
    // console.log(this.state.shopid)
  await  this.setState({
      shopid:localStorage.getItem('shopid')
    })
    this.getAllClass()
    this.getList()
  }
  //获取所有分类
  getAllClass(){
    axios.post('/api2',qs.stringify({
      token:'Jh2044695',
      cmd:'getAllClass',
      shopid:this.state.shopid
    })).then(res=>{
      console.log(res.data.data)
      if(res.data.success){
        if(res.data.data){
          let arr=res.data.data
          let options=[]
          for(let i=0;i<arr.length;i++){
              let oneobj={}
              oneobj.value=arr[i].id
              oneobj.label=arr[i].menuname
              let children=[]
              if(arr[i].children){
                for(let j=0;j<arr[i].children.length;j++){
                let twoobj={}
                twoobj.value=arr[i].children[j].id
                twoobj.label=arr[i].children[j].menuname
                children.push(twoobj)
              }
              }
              oneobj.children=children
              
              options.push(oneobj)
          }
        this.setState({
          options:options
        })
        }
      

      }
    })
  }
  //获取列表
  async getList(){
    await axios.post('/api2',qs.stringify({
      token: 'Jh2044695',
      cmd: 'getGoodsList',
      shopid:this.state.shopid,
      classone:this.state.classone,
      classtwo:this.state.classtwo,
      goodsname:this.state.goodsname,
      page:this.state.page,
      count:this.state.pageSize
    })).then(res=>{
      console.log(res.data)
      if(res.data.success){
        this.setState({
          list:res.data.data.goodsList,
          total:res.data.data.total
        })
      }else{
        this.setState({
          list:[],
          total:0
        })
      }
    })
  }
  //切换每页的数量
  onShowSizeChange = (current, pageSize) => {
    this.setState({
      pageSize: pageSize,
      page: current
    },()=>{
      this.getList();
    })
  };
  //切换到第几页
  changePage = (page, pageSize) => {
    this.setState({
      pageSize: pageSize,
      page: page
    },()=>{
      this.getList();
    })
  }
  //上架商品
  addProduct=()=>{
    router.push({
      pathname:'/goodsmanagement/productmanage/newproduct',
      query:{

      }
    })
  }
  //分类选择
  onChange=(value)=> {
    this.setState({
      classone:value[0],
      classtwo:value[1]?value[1]:''
    })
  }
  //输入框变化
  changeInput=(event)=>{
    event.persist();
    this.setState({
      goodsname:event.target.value
    })
  }
  //点击搜索
  search=()=>{
    // console.log(this.state)
    this.getList()
  }
  //显示删除对话框
  showModal=(record)=>{
    this.setState({
      visible:true,
      goodsid:record.id
    })
    console.log(record)
  }
  //确认删除
  handleOk=()=>{
    this.setState({
      confirmLoading: true
    });
    axios.post('/api2',qs.stringify({cmd:'deleteGoods',token:'Jh2044695',goodsid:this.state.goodsid})).then(res=>{
      console.log(res.data)
      if(res.data.success){
        this.setState({
          visible: false,
          confirmLoading: false,
        });
        message.success('删除成功！')
        this.getList()
      }else{
        message.error(res.data.errmsg)
    }
    })
    // setTimeout(() => {
    //   this.setState({
    //     visible: false,
    //     confirmLoading: false,
    //   });
    // }, 2000);
  }
  //删除弹框 点击取消
  handleCancel = () => {
    this.setState({
      visible: false,
    });
  };
  //跳转至颜色管理
  toColorManage=(record)=>{
    console.log(record)
    router.push({
      pathname:'/goodsmanagement/productmanage/colormanage',
      query:{
        goodsid:record.id
      }
    })
  }
  //跳转至编辑界面
  toEditGoods=(record)=>{
    router.push({
      pathname:'/goodsmanagement/productmanage/editgoods',
      query:{
        goodsid:record.id
      }
    })
  }
  //上架或下架
  goodsState=(record)=>{
    console.log(record)
    let that=this
    let title=record.state==0?'上架':'下架'
    Modal.confirm({
      title:title+'确认',
      content:`确认要改${title}商品吗？`,
      okText:'确认',
      cancelText:'取消',
      onCancel(){},
      onOk(){
            axios.post('/api2',qs.stringify({
              cmd:'updateGoodsState',
              token:'Jh2044695',
              id:record.id,
              state:record.state
            })).then(res=>{
                if(res.data.success){
                  message.success(res.data.errmsg)
                  that.getList()
                }else{
                  message.error(res.data.errmsg)
                }
            })
      }
    })

  }
  render(){
  const { visible, confirmLoading, ModalText } = this.state;
  const  columns=[
      {
        title: '产品名称',
        dataIndex: 'goodsname',
        key:'goodsname'
      },
      {
        title: '一级分类',
        dataIndex: 'classone',
        key: 'classone'
      },
      {
        title: '二级分类',
        dataIndex: 'classtwo',
        key: 'classtwo'
      },
      {
        title: '商品点击量',
        dataIndex: 'clicknum',
        key: 'clicknum'
      },
      {
        title: '操作',
        key: 'action',
        render: (text, record) => (
          <span>
            <a onClick={() => this.showModal(record)}>删除</a>
            <Divider type="vertical" />
            <a onClick={()=>this.toColorManage(record)}>颜色管理</a>
            <Divider type="vertical" />
            <a onClick={()=>this.toEditGoods(record)}>编辑</a>
            <Divider type="vertical" />
            <a onClick={()=>this.goodsState(record)}>{record.state==0?'上架':'下架'}</a>
          </span>
        ),
      },
    ];
    return (
      <PageHeaderWrapper title='产品管理'>
        <div>
          <div> 
            <Button  type='primary' onClick={this.addProduct}>
                上架商品
            </Button>
              </div>
              <div style={{margin:'10px'}}>
                <Cascader 
                options={this.state.options}
                expandTrigger="hover"
                style={{width:'400px'}}
                onChange={this.onChange}/>
                <Input 
                placeholder='请输入产品名称' 
                style={{width:'150px',margin:'0 10px'}} 
                onChange={event=>{this.changeInput(event)}}
                />
                <Button icon='search' type='primary' onClick={this.search}>搜索</Button>
              </div>
            <Table columns={columns} dataSource={this.state.list} pagination={false} />
            <Pagination
              total={this.state.total}
              defaultCurrent={1}
              showTotal={total => `共 ${total} 条`}
              showSizeChanger={true}
              onShowSizeChange={this.onShowSizeChange}
              showQuickJumper
              onChange={this.changePage}
              style={{ margin: '10px 0', textAlign: 'right' }}
        />
      </div>
      <Modal
          title="删除确认"
          visible={visible}
          onOk={this.handleOk}
          confirmLoading={confirmLoading}
          onCancel={this.handleCancel}
        >
          <p>{ModalText}</p>
        </Modal>
      </PageHeaderWrapper>
      
   )
  }
}

export default ProductManage