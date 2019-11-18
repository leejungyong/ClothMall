import React, {
  PureComponent
} from 'react'
import {
  connect
} from 'dva'
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
import {
  router
} from 'umi';

const formItemLayout = {
  labelCol: {
    xs: {
      span: 0
    },
    sm: {
      span: 3
    },
  },
  wrapperCol: {
    xs: {
      span: 10
    },
    sm: {
      span: 10
    },
  },
};
@Form.create()
class EditImg extends PureComponent {
  state = {
    loading: false,
    wechaturlImageUrl: '', //图像链接微信
    wechaturlPicParams: "", //传给后端的链接微信
    shopshowImageUrl: '', //图像链接店铺展示
    shopshowPicParams: "", //传给后端的链接店铺展示
    logoimgImageUrl: '', //图像链接店铺图标
    logoimgPicParams: "", //传给后端的链接店铺图标
    bannerimgPicParams:[],//传给后端的链接广告图
    fileList: '', //动态图
    previewVisible: false,
    fileListParams: new Array, //上传图片数组
    previewImage: '', //大屏展示
    editValue: "" //编辑信息
  }
 
  async componentDidMount() {
    const editV = await axios.post('/api2', qs.stringify({
      token: 'Jh2044695',
      cmd: 'getShopInfo',
      id: this.props.location.query.id
    }))

    this.setState({
      editValue: editV.data.data,
      bannerimgPicParams:editV.data.data.bannerimg?editV.data.data.bannerimg.split(','):[],
      fileList: editV.data.data.bannerimg?editV.data.data.bannerimg.split(',').map((item, index) => {
        var newItem = {
          uid: index - 1,
          name: 'image.png',
          status: 'done',
          url: item,
        }
        return newItem
      }):[]
    })
  }
  handleChange = ({
    fileList,
    file
  }) => {
    this.setState({
      fileList
    })
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
  handleCancel = () => this.setState({
    previewVisible: false
  });
  //确认新增店铺
  handleSubmit = () => {
    this.props.form.validateFields((err, values) => {
      console.log(err)
      if (!err) {
        console.log(values);
        axios.post('/api2', qs.stringify({
          cmd: 'updateShopBannerImg',
          token: 'Jh2044695',
          shopid: this.props.location.query.id,
          bannerimg:this.state.bannerimgPicParams
        })).then(res => {
          if (res.data.success) {
            message.success(res.data.errmsg)
            router.push({
              pathname: '/shopmanagement'
            })
          }
          if (!res.data.success) {
            message.error(res.data.errmsg)

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
      this.setState({
        loading: true
      });
      return;
    }
    if (info.file.status === 'done') {
      // Get this url from response in real world.

      this.getBase64(info.file.originFileObj, imageUrl =>
        this.setState({
          wechaturlImageUrl: imageUrl,
          loading: false,
          wechaturlPicParams: info.file.response.data[0].image
        }),
      );
    }
  };
  handleChangeShopshow = info => {
    if (info.file.status === 'uploading') {
      this.setState({
        loading: true
      });
      return;
    }
    if (info.file.status === 'done') {
      // Get this url from response in real world.

      this.getBase64(info.file.originFileObj, imageUrl =>
        this.setState({
          shopshowImageUrl: imageUrl,
          loading: false,
          shopshowPicParams: info.file.response.data[0].image
        }),
      );
    }
  };
  handleChangeLogoimg = info => {
    if (info.file.status === 'uploading') {
      this.setState({
        loading: true
      });
      return;
    }
    if (info.file.status === 'done') {
      // Get this url from response in real world.

      this.getBase64(info.file.originFileObj, imageUrl =>
        this.setState({
          logoimgImageUrl: imageUrl,
          loading: false,
          logoimgPicParams: info.file.response.data[0].image
        }),
      );
    }
  };
  render() {
    const {
      getFieldDecorator
    } = this.props.form;
    const uploadButton = ( <
      div >
      <
      Icon type = {
        this.state.loading ? 'loading' : 'plus'
      }
      /> <
      div className = "ant-upload-text" > Upload < /div> <
      /div>
    );
    const {
      wechaturlImageUrl,
      shopshowImageUrl,
      bannerimgImageUrl,
      logoimgImageUrl,
      previewVisible,
      previewImage,
      fileList
    } = this.state;
    return ( <
      PageHeaderWrapper title = "编辑广告轮播图" >
      <
      Form {
        ...formItemLayout
      }
      onSubmit = {
        this.handleSubmit
      }
      className = "" >
      <
      div >
      <
      span className = {
        styles.tipUpload
      } > 编辑店铺轮播图 < /span> <
      div className = {
        styles.UploadContant1
      } >
      <
      Upload action = "/multiupload"
      listType = "picture-card"
      fileList = {
        fileList
      }
      onPreview = {
        this.handlePreview
      }
      onChange = {
        this.handleChange
      } > {
        fileList.length >= 6 ? null : uploadButton
      } <
      /Upload> <
      /div> <
      Modal visible = {
        previewVisible
      }
      footer = {
        null
      }
      onCancel = {
        this.handleCancel
      } >
      <
      img alt = "example"
      style = {
        {
          width: '100%'
        }
      }
      src = {
        previewImage
      }
      /> <
      /Modal> <
      /div>

      <
      /Form> <
      Button type = "primary"
      onClick = {
        () => this.handleSubmit()
      }
      className = {
        styles.commit
      } > 提交 < /Button> <
      /PageHeaderWrapper>
    );
  }
}
export default EditImg;
