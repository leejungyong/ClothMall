<template>
  <div class="home">
    <div class="top">
      <div class="img-box">
        <img :src="shopObj.logoimg" style="width:100%;height:100%;" alt />
      </div>
      <div class="name">{{shopObj.shopname}}</div>
      <div class="btn-box fr clearfix">
        <span class="collect" @click="love">收藏本店</span>
        <span class="collect" @click="toContact">联系方式</span>
      </div>
    </div>
    <div class="introduce">
      <img :src="shopObj.shopshow" style="width:100%;height:100%;" alt />
    </div>
    <div class="swiper">
      <div class="swiper-container">
        <div class="swiper-wrapper">
          <div class="swiper-slide" v-for="(item,index) in swiperList" :key="index">
            <img :src="item" style="width:3.35rem;height:1.88rem" alt />
          </div>
        </div>
        <div class="swiper-pagination"></div>
      </div>
    </div>
    <div class="nav">
      <div class="common all-type" @click="toAllType">
        <div>
          <span>全部分类</span>
          <img style="margin-top:0.04rem;" src="../assets/icon/home_icon_more.png" alt />
        </div>
        <div class="fr clearfix icon">
          <img src="../assets/icon/home_icon_fenlei.png" alt />
        </div>
      </div>
      <div class="common my-tracks fr clearfix" @click="toMyTracks">
        <div>
          <span>我的足迹</span>
          <img style="margin-top:0.04rem;" src="../assets/icon/home_icon_more.png" alt />
        </div>
        <div class="fr clearfix icon">
          <img src="../assets/icon/home_icon_love.png" alt />
        </div>
      </div>
    </div>
    <div class="hot-goods">
      <div class="title">热门商品</div>
      <div class="list">
        <div
          class="item"
          v-for="(item,index) in hotgoods"
          :key="index"
          @click="toDetail(item.id)"
          :style="{backgroundImage:'url('+item.smallpic+')'}"
        >
          <span class="t-left">{{item.menuname}}</span>
          <span class="t-right fr clearfix">
            <img src="../assets/icon/icon_3d.png" alt />
          </span>
          <div class="bottom">
            <img src="../assets/icon/icon_huo.png" alt />
            <span>{{item.clicknum}}</span>
            <span class="fr clearfix">{{item.colornum}}种颜色</span>
          </div>
        </div>
      </div>
    </div>
    <div class="tips">
      <div class="box">
        <span>没有了，去查看</span>
        <span style="color:#d8a163;" @click="toAllType">全部分类</span>
      </div>
    </div>
    <div class="mask" v-if="showMask" @click="hiddenMask">
      <img src="../assets/icon/弹层.png" alt  />
    </div>
  </div>
</template>

<script>
import Swiper from "swiper";
import axios from "axios";
import qs from "qs";
import { getIndexInfo } from "@/api/index.js";
export default {
  data() {
    return {
      shopObj: {
        id: null,
        shopname: "",
        logoimg: "",
        shopshow: "",
        bannerImg: "",
        hotgoods: []
      },
      hotgoods: [1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1], //热门商品列表
      showMask: false,
      swiperList: []
    };
  },
  mounted() {
    this.fetchData();
  },
  created() {},
  methods: {
    fetchData() {
      axios
        .post(
          "/api",
          qs.stringify({ shopid: 1, token: "Jh2044695", cmd: "getIndexInfo" })
        )
        .then(res => {
          console.log(res.data);
          if (res.data.success) {
            this.shopObj = res.data.data;
            this.swiperList = res.data.data.bannerImg.split(",");
            this.hotgoods = res.data.data.hotgoods;
            setTimeout(function() {
              var mySwiper = new Swiper(".swiper-container", {
                loop: true,
                pagination: {
                  el: ".swiper-pagination"
                },
                autoplay: {
                  delay: 2500,
                  disableOnInteraction: false
                }
              });
            }, 100);
          }
        });
    },
    hiddenMask() {
      this.showMask = false;
    },
    toAllType() {
      this.$router.push("/alltypes");
    },
    toMyTracks() {
      this.$router.push("/mytracks");
    },
    toContact() {
      this.$router.push({ path: "/contact", query: { shopid: 1 } });
    },
    love() {
      this.showMask = true;
    },
    toDetail(id) {
      this.$router.push({ name: "GoodsDetail", query: { id } });
    }
  }
};
</script>

<style lang="less" scoped>
@import "../css/common.css";
.home {
  padding: 0.2rem;
  .top {
    // width: 100%;
    height: 0.3rem;
    line-height: 0.3rem;
    // background: red;
    .img-box {
      display: inline-block;
      width: 0.3rem;
      height: 0.3rem;
      border-radius: 0.04rem;
    }
    .img {
      display: inline-block;
      width: 0.3rem;
      height: 0.3rem;
      border-radius: 0.04rem;
    }
    .name {
      display: inline-block;
      margin-left: 0.1rem;
      width: 1.5rem;
      font-size: 0.18rem;
      font-weight: 500;
      color: #333;
      overflow: hidden;
      text-overflow: ellipsis;
      white-space: nowrap;
      font-family:PingFangSC-Medium,PingFang SC;

    }
    .btn-box {
      display: inline-block;
      height: 0.3rem;
      line-height: 0.3rem;
      .collect {
        display: inline-block;
        background: #f0f0f0;
        // margin-left: 0.1rem;
        width: 0.6rem;
        height: 0.24rem;
        line-height: 0.24rem;
        border-radius: 0.12rem;
        color: #d8a163;
        font-size: 0.12rem;
        text-align: center;
      }
    }
  }
  .introduce {
    margin-top: 0.25rem;
    // width: 3.35rem;
    width: auto;
    height: 1.3rem;
    background: #d8a163;
  }
  .swiper {
    margin-top: 0.2rem;
    height: 1.88rem;
    width: auto;  
  }
  .nav {
    margin: 0.25rem 0;
    .common {
      display: inline-block;
      width: 1.6rem;
      height: 0.7rem;
      color: #d8a163;
      background: #ffe7cc;
      border-radius: 0.04rem;
      font-size: 0.18rem;
      padding: 0.05rem 0.1rem;
      box-sizing: border-box;
    }
    .icon {
      margin-top: 0.05rem;
    }
  }
  .hot-goods {
    .title {
      height: 0.28rem;
      font-size: 0.2rem;
      font-weight: 500;
      font-family:PingFangSC-Medium,PingFang SC;
      color: rgba(51, 51, 51, 1);
      line-height: 0.28rem;
    }
    .list {
      padding: 0.2rem 0;
      display: flex;
      justify-content: space-between;
      flex-wrap: wrap;
      .item {
        display: inline-block;
        width: 1.58rem;
        height: 1.78rem;
        border-radius: 0.04rem;
        box-shadow: 0 0 0.16rem 0 rgba(214, 214, 214, 0.5);
        // background: #d8a163;
        margin-bottom: 0.2rem;
        position: relative;
        .t-left {
          display: inline-block;
          width: 0.8rem;
          height: 0.27rem;
          line-height: 0.27rem;
          font-size: 0.12rem;
          color: #333;
          background: linear-gradient(
            90deg,
            rgba(255, 255, 255, 0) 0%,
            rgba(255, 255, 255, 1) 100%
          );
        }
        .t-right {
          display: inline-block;
          margin: 0.05rem 0.05rem 0 0;
        }
        .bottom {
          width: 100%;
          height: 0.2rem;
          line-height: 0.2rem;
          background: rgba(255, 255, 255, 1);
          border-radius: 0 0 0.04rem 0.04rem;
          opacity: 0.8;
          color: #333;
          padding: 0 0.05rem;
          position: absolute;
          bottom: 0;
          left: 0;
        }
      }
    }
  }
  .tips {
    text-align: center;
    .box {
      display: inline-block;
      width: 1.56rem;
      height: 0.26rem;
      line-height: 0.26rem;
      background: rgba(240, 240, 240, 1);
      border-radius: 0.13rem;
      color: #999;
      font-size: 0.12rem;
    }
  }
  .mask {
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.4);
    position: fixed;
    top: 0;
    left: 0;
    text-align: center;
    z-index: 99;
  }
}
</style>