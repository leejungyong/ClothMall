<template>
  <div class="contact">
    <div class="tel">
      <div class="icon">
        <img src="../assets/icon/235座机.png" alt />
      </div>
      <div class="right">
        <div class="title">座机号</div>
        <div class="content">{{contact.telnum}}</div>
      </div>
    </div>
    <div class="tel">
      <div class="icon">
        <img src="../assets/icon/手机.png" alt />
      </div>
      <div class="right">
        <div class="title">手机号</div>
        <div class="content">{{contact.phonenum}}</div>
      </div>
    </div>
    <div class="tel">
      <div class="icon">
        <img src="../assets/icon/微信.png" alt />
      </div>
      <div class="right">
        <div class="title">微信号</div>
        <div class="content">{{contact.wechat}}</div>
        <div>
          <div class="code">
            <!-- :style="{backgroundImage:'url('+contact.wechaturl+')'}" -->
            <img :src="contact.wechaturl" style="width:100%;height:100%;" alt />
            <!-- <canvas id="QRCode" class="qr"></canvas> -->
          </div>
          <div class="tip">识别二维码 添加微信</div>
        </div>
      </div>
    </div>
    <div class="tel address">
      <div class="icon">
        <img src="../assets/icon/地址.png" alt />
      </div>
      <div class="right">
        <div class="title">详细地址</div>
        <div class="content">{{contact.location}}</div>
      </div>
    </div>
    <div id="container" style="width:100%;height:200px;"></div>
    <!-- <button @click="getloc">获取</button> -->
  </div>
</template>

<script>
import QRCode from "qrcode";
import axios from "axios";
import qs from "qs";
import { getShopContactInfo } from "@/api/index.js";
export default {
  data() {
    return {
      QRCodeMsg: "",
      shopid: "",
      contact: {},
      location: ""
    };
  },
  async created() {
    // this.getQRCode();
    this.shopid = this.$route.query.shopid;
    await this.fetchData();
    // this.map()
  },
  mounted() {},
  watch: {
    // 通过监听获取数据
    QRCodeMsg(val) {
      // 获取页面的canvas
      var msg = document.getElementById("QRCode");
      // 将获取到的数据（val）画到msg（canvas）上
      QRCode.toCanvas(msg, val, function(error) {
        console.log(error);
      });
    }
  },
  methods: {
    fetchData() {
      axios
        .post(
          "/api",
          qs.stringify({
            shopid: 1,
            token: "Jh2044695",
            cmd: "getShopContactInfo"
          })
        )
        .then(res => {
          if (res.data.success) {
            this.contact = res.data.data;
            this.location = res.data.data.location;
            this.getlocation(this.location);
          }
        });
    },
    //根据详细地址获取lnglat
    getlocation(loc) {
      console.log(loc);
      axios
        .get(
          "http://restapi.amap.com/v3/geocode/geo?key=389880a06e3f893ea46036f030c94700&s=rsv3&city=35&address=" +
            loc
        )
        .then(res => {
          console.log(res.data);
          if (res.data.status == 1) {
            this.location = res.data.geocodes[0].location;
            console.log("dasdas", this.location);
            this.map();
          }
        });
    },
    //渲染地图
    map() {
      console.log(this.location);
      let latlng = this.location.split(",");
      var marker = new AMap.Marker({
        position: [latlng[0], latlng[1]], // 经纬度对象，也可以是经纬度构成的一维数组[116.39, 39.9]
      });
      var map = new AMap.Map("container", {
        zoom: 10, //级别
        center: [latlng[0], latlng[1]], //中心点坐标
        viewMode: "3D" //使用3D视图
      });
    map.add(marker);
    },
    //生成二维码
    getQRCode() {
      this.QRCodeMsg = window.location.href; //生成的二维码为URL地址
    }
  }
};
</script>

<style lang="less" scoped>
.contact {
  padding: 0.1rem 0.2rem;
  .tel {
    width: 3.35rem;
    // height: 0.75rem;
    padding-left: 0.1rem;
    background: rgba(255, 255, 255, 1);
    border-bottom: 0.01rem solid rgba(214, 214, 214, 0.5);
    display: table;
  }
  .icon {
    height: 0.75rem;
    width: 0.3rem;
    display: table-cell;
    vertical-align: middle;
    img {
    }
  }

  .right {
    display: inline-block;
    box-sizing: border-box;
    margin: 0.15rem;
    .title {
      color: #ccc;
    }
    .content {
      font-size: 0.18rem;
      color: #333;
      width: 1.62rem;
      overflow: hidden;
      text-overflow: ellipsis;
    }
    .code {
      display: inline-block;
      width: 0.68rem;
      height: 0.68rem;
      // line-height: 0.68rem;
      // text-align: center;
      // opacity:0.65;
      border: 0.01rem solid rgba(0, 0, 0, 1);
      padding: 0.04rem;
    }
    .qr {
      width: 0.59rem !important;
      height: 0.59rem !important;
    }
    .tip {
      display: inline-block;
      font-size: 0.12rem;
      color: #333;
      width: 0.65rem;
    }
  }
}
</style>