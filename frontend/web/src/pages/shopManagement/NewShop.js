import React,{PureComponent} from 'react'
import {connect} from 'dva'
import axios from 'axios'
import qs from 'qs'
import {
    Input,
    Form,
    Icon,
    Button,
    Cascader,
    message,
   Upload,
   Modal 
} from 'antd'

import PageHeaderWrapper from '@/components/PageHeaderWrapper';

import styles from './NewShop.less'
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
@Form.create()
class NewShop extends PureComponent{
    state={
        loading: false,
        wechaturlImageUrl:'',//图像链接微信
        wechaturlPicParams:"",//传给后端的链接微信
        shopshowImageUrl:'',//图像链接店铺展示
        shopshowPicParams:"",//传给后端的链接店铺展示
        bannerimgImageUrl:'',//图像链接广告图
        bannerimgPicParams:[],//传给后端的链接广告图
        logoimgImageUrl:'',//图像链接店铺图标
        logoimgPicParams:"",//传给后端的链接店铺图标
        fileList:"",//动态图
        previewVisible: false,
        previewImage: '',
        fileListParams:new Array,//上传图片数组
    }

    componentDidMount(){
    }
    handleChange = ({ fileList,file }) => {
        this.setState({fileList})
        if (file.status === 'done') {
          this.state.bannerimgPicParams.push(file.response.data[0].image);
          console.log(this.state.bannerimgPicParams);
        }
    };

    handlePreview = async file => {
        if (!file.url && !file.preview) {
          file.preview = await this.getBase64More(file.originFileObj);
        }
        console.log(file);
        this.setState({
          previewImage: file.url || file.preview,
          previewVisible: true,
        });
      };
    handleCancel = () => this.setState({ previewVisible: false });
    //确认新增店铺
    handleSubmit = ()=> {
        console.log(this.state.bannerimgPicParams);
        this.props.form.validateFields((err, values) => {
            console.log(err)
          if (!err) {
            console.log(values);
            axios.post('/api2',qs.stringify({
                cmd:'insertNewShop',
                token:'Jh2044695',
                shopurl:values.shopurl,
                shopname:values.shopname,
                location:values.location,
                bossname:values.bossname,
                telnum:values.telnum,
                phonenum:values.phonenum,
                wechat:values.wechat,
                wechaturl:this.state.wechaturlPicParams,
                shopshow:this.state.shopshowPicParams,
                logoimg:this.state.logoimgPicParams,
                bannerimg:this.state.bannerimgPicParams,
            })).then(res=>{
                if(res.data.success){
                    console.log(res.data)
                    message.success(res.data.errmsg)
                    this.setState({
                        newproid:res.data.data.id
                    })
                    router.push({
                        pathname:'/shopmanagement/ConfigMachine',
                        query:{shopid:res.data.data}
                    })
                }
            })
          }
        });
      };
      getBase64More(file) {
        return new Promise((resolve, reject) => {
          const reader = new FileReader();
          reader.readAsDataURL(file);
          reader.onload = () => resolve(reader.result);
          reader.onerror = error => reject(error);
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
    handleChangeWechaturl = info => {
        if (info.file.status === 'uploading') {
          this.setState({ loading: true });
          return;
        }
        if (info.file.status === 'done') {
          // Get this url from response in real world.
          
          this.getBase64(info.file.originFileObj, imageUrl =>
            this.setState({
                wechaturlImageUrl:imageUrl,
              loading: false,
              wechaturlPicParams:info.file.response.data[0].image
            }),
          );
        }
      };
      handleChangeShopshow = info => {
        if (info.file.status === 'uploading') {
          this.setState({ loading: true });
          return;
        }
        if (info.file.status === 'done') {
          // Get this url from response in real world.
          
          this.getBase64(info.file.originFileObj, imageUrl =>
            this.setState({
                shopshowImageUrl:imageUrl,
              loading: false,
              shopshowPicParams:info.file.response.data[0].image
            }),
          );
        }
      };
      handleChangeLogoimg = info => {
        if (info.file.status === 'uploading') {
          this.setState({ loading: true });
          return;
        }
        if (info.file.status === 'done') {
          // Get this url from response in real world.
          
          this.getBase64(info.file.originFileObj, imageUrl =>
            this.setState({
                logoimgImageUrl:imageUrl,
              loading: false,
              logoimgPicParams:info.file.response.data[0].image
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
          const { wechaturlImageUrl,shopshowImageUrl,bannerimgImageUrl,logoimgImageUrl,previewVisible, previewImage, fileList } = this.state;
        return (
            <PageHeaderWrapper title="新增店铺">
    <Form {...formItemLayout} onSubmit={this.handleSubmit} className="" >
                <Form.Item label='店铺url名称:' >
                {getFieldDecorator('shopurl', {
                    rules: [{ required: true, message: '请输入店铺url名称!' }],
                })(
                    <Input
                    placeholder="请输入店铺url名称"
                    />,
                )}
                </Form.Item>
                <Form.Item label='店铺名称:'>
                {getFieldDecorator('shopname', {
                    rules: [{ required: true, message: '请输入店铺名称!' }],
                })(
                    <Input
                    placeholder="请输入店铺名称"
                    />,
                )}
                </Form.Item>
                <Form.Item label='店铺详细地址:'>
                {getFieldDecorator('location', {
                    rules: [{ required: true, message: '请输入店铺详细地址!' }],
                })(
                    <Input
                    placeholder="请输入店铺详细地址"
                    />,
                )}
                </Form.Item>
                <Form.Item label='老板姓名:'>
                {getFieldDecorator('bossname')(
                    <Input
                    placeholder="请输入老板姓名"
                    />,
                )}
                </Form.Item>
                <Form.Item label='座机号码:'>
                {getFieldDecorator('telnum')(
                    <Input
                    placeholder="请输入座机号码"
                    />,
                )}
                </Form.Item>
                <Form.Item label='手机号码:'>
                {getFieldDecorator('phonenum')(
                    <Input
                    placeholder="请输入手机号码"
                    />,
                )}
                </Form.Item>
                <Form.Item label='微信号:'>
                {getFieldDecorator('wechat')(
                    <Input
                    placeholder="请输入微信号"
                    />,
                )}
                </Form.Item>
                <span className={styles.tipUpload}>上传微信号二维码图片</span>
                    <div className={styles.UploadContant}>
                    <Upload
                        data={{token:'Jh2044695',
                        module:'shoppic',
                        kresize:'{"key": "k", "Thumbonly":0, "Width":100, "Height":100}'
                    }}
                        name="k"
                        listType="picture-card"
                        className="avatar-uploader"
                        showUploadList={false}
                        action="/multiupload"
                        beforeUpload={this.beforeUpload}
                        onChange={this.handleChangeWechaturl}
                    >
                        {wechaturlImageUrl ? <img src={wechaturlImageUrl} alt="avatar" style={{ width: '100%' }} /> : uploadButton}
                    </Upload>
                    </div>
                    <div>
                    <span className={styles.tipUpload}>上传店铺展示</span>
                    <div className={styles.UploadContant}>
                    <Upload
                        data={{token:'Jh2044695',
                        module:'shoppic',
                        kresize:'{"key": "k", "Thumbonly":0, "Width":100, "Height":100}'
                    }}
                        name="k"
                        listType="picture-card"
                        className="avatar-uploader"
                        showUploadList={false}
                        action="/multiupload"
                        beforeUpload={this.beforeUpload}
                        onChange={this.handleChangeShopshow}
                    >
                        {shopshowImageUrl ? <img src={shopshowImageUrl} alt="avatar" style={{ width: '100%' }} /> : uploadButton}
                    </Upload>
                    </div>
                    </div>
                    <div>
                    <span className={styles.tipUpload}>上传店铺Logo</span>
                    <div className={styles.UploadContant}>
                    <Upload
                        data={{token:'Jh2044695',
                        module:'shoppic',
                        kresize:'{"key": "k", "Thumbonly":0, "Width":100, "Height":100}'
                    }}
                        name="k"
                        listType="picture-card"
                        className="avatar-uploader"
                        showUploadList={false}
                        action="/multiupload"
                        beforeUpload={this.beforeUpload}
                        onChange={this.handleChangeLogoimg}
                    >
                        {logoimgImageUrl ? <img src={logoimgImageUrl} alt="avatar" style={{ width: '100%' }} /> : uploadButton}
                    </Upload>
                    </div>
                    </div>
                    <div>
                    <div className="clearfix">
                    <span className={styles.tipUpload}>上传店铺广告</span>
                    <div className={styles.UploadContant1}>
                        <Upload
                 action="/multiupload"
                     listType="picture-card"
                      fileList={fileList}
                      onPreview={this.handlePreview}
                      onChange={this.handleChange} >
                         {fileList.length >= 6 ? null : uploadButton}
                    </Upload>
                      </div>
                    <Modal visible={previewVisible} footer={null} onCancel={this.handleCancel}>
                     <img alt="example" style={{ width: '100%' }} src={previewImage} />
                    </Modal>
                        </div>
                        </div>
                    
            </Form>
            <Button type="primary" onClick={() => this.handleSubmit()} className={styles.commit}>提交</Button>
            </PageHeaderWrapper>
        );
      }
}
export default NewShop;