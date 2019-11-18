<template>
  <div class="tracks">
    <div class="user-info">
      <span class="tel">当前登录账号:{{phonenum}}</span>
      <!-- <span></span> -->
      <span class="logout" @click="logOut">退出登录</span>
    </div>
    <div class="tab-box">
      <ul class="tab">
        <li
          class="type fl clearfix"
          v-for="(tab,index) in tabs"
          :key="index"
          @click="changeTab(index,tab.view)"
          :class="{active:currentIndex==index}"
        >{{tab.type}}</li>
      </ul>
    </div>
    <div class="record-list" v-if="currentIndex==0&&record">
      <div class="item" v-for="(item,index) in record" :key="index" @click="goDetail(item.id)">
        <div class="img-box">
          <img :src="item.picList[1].compresspic" style="width:100%;height:100%;" alt />
        </div>
        <div class="info-box fr clearfix">
          <div>名称：{{item.goodsname}}</div>
          <div>型号：{{item.picList[1].model}}</div>
          <div>规格：{{item.unit}}</div>
          <div>价格：{{item.price}}</div>
        </div>
      </div>
    </div>
    <div class="goods-love-list" v-if="currentIndex==1&&goodsLove">
      <div class="title">卧室</div>
      <div class="item" v-for="(item,index) in goodsLove" :key="index">
        <div class="img-box">
          <img
            class="cancel-love"
            src="../../assets/icon/icon_love_1.png"
            alt
            @click="cancelLove(item.id)"
          />
          <img :src="item.compressPic" style="width:100%;height:100%;" alt />
        </div>
        <div class="info-box fr clearfix">
          <div>名称：{{item.goodsName}}</div>
          <div>风格：{{item.style}}</div>
          <div>规格：{{item.unit}}</div>
          <div>价格：{{item.price}}</div>
        </div>
      </div>
    </div>
    <div class="scene-list" v-if="currentIndex==2&&sceneLove">
      <!-- <div class="title">卧室</div> -->
      <div class="item" v-for="(item,index) in sceneLove" :key="index">
        <div class="img-box">
          <img
            class="cancel-love fr clearfix"
            src="../../assets/icon/icon_love_1.png"
            alt
            @click="cancelLove(item.id)"
          />
          <img :src="item.compressPic" style="width:100%;height:100%;" alt />
        </div>
        <div class="info-box fr clearfix">
          <div>名称：{{item.goodsname}}</div>
          <div>风格：{{item.style}}</div>
          <div>规格：{{item.unit}}</div>
          <div style="width:1rem;display:inline-block;">价格：{{item.price}}/m2</div>
          <div class="to3d fr clearfix" @click="to3D">查看3D效果</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import record from './record.vue'
// import goodslove from './goods-love.vue'
// import scene from './scene-love.vue'
import axios from "axios";
import qs from "qs";
export default {
  data() {
    return {
      currentIndex: 0,
      tabs: [
        {
          type: "浏览记录",
          view: "record"
        },
        {
          type: "商品收藏",
          view: "goodslove"
        },
        {
          type: "场景收藏",
          view: "scene"
        }
      ],

      currentView: "record",
      list: [1, 1, 1],
      record: [], ///浏览记录
      goodsLove: [], //收藏的商品列表
      sceneLove: [], //收藏的场景
      phonenum:''
    };
  },
  components: {},
  created() {
    this.phonenum=localStorage.getItem('phonenum')
    this.fetch();
    this.getRecord();
  },
  methods: {
    fetch() {
      axios
        .post(
          "/api",
          qs.stringify({
            token: "Jh2044695",
            userid: 1,
            shopid: 1,
            cmd: "getCollection"
          })
        )
        .then(res => {
          console.log(res.data);
          if (res.data.success) {
            // this.record=res.data.data.Record
            this.goodsLove = res.data.data.shopCollection;
            this.sceneLove = res.data.data.scenesCollection;
          }
        });
    },

    //退出登录
    logOut() {
      // localStorage.removeItem('phonenum')
      localStorage.clear()
      this.$router.push('/login')
    },
    goDetail(id) {
      this.$router.push({ name: "GoodsDetail", query: { id } });
    },
    //获取localStorage中的浏览记录
    getRecord() {
      let record = JSON.parse(localStorage.getItem("record"));
      this.record = record;
    },

    //切换tab
    changeTab(index, view) {
      this.currentIndex = index;
      this.currentView = view;
    },

    //取消收藏
    cancelLove(id) {
      axios
        .post(
          "/api",
          qs.stringify({
            token: "Jh2044695",
            userid: 1,
            collectionid: id,
            cmd: 'delCollection'
          })
        )
        .then(res => {
          if (res.data.success) {
            console.log(res.data);
            console.log("取消收藏成功");
            this.fetch();
          }
        });
    },

    //查看3D效果
    to3D() {}
  }
};
</script>
<style lang="less" scoped>
.tracks {
  padding: 0.1rem 0.22rem;
  .user-info {
    // width:183px;
    height: 0.2rem;
    line-height: 0.2rem;
    color: rgba(0, 0, 0, 1);
    .logout {
      color: #d8a163;
      display: inline-block;
      margin-left: 0.2rem;
    }
  }
  .tab {
    width: 3.3rem;
    height: 0.32rem;
    background: rgba(240, 240, 240, 1);
    border-radius: 0.04rem;
    margin: 0.2rem auto 0 auto;
    .type {
      width: 1.1rem;
      height: 0.32rem;
      line-height: 0.32rem;
      text-align: center;
      background: rgba(240, 240, 240, 1);
      color: #d8a163;
      border-radius: 0.04rem;
    }
    .active {
      background: #d8a163;
      color: #fff;
    }
  }
  .record-list {
    margin-top: 0.2rem;
    .title {
      font-size: 0.2rem;
      margin-bottom: 0.1rem;
    }
    .item {
      width: 3.35rem;
      height: 1.15rem;
      background: rgba(255, 255, 255, 1);
      box-shadow: 0px 0px 0.16rem 0px rgba(214, 214, 214, 0.5);
      border-radius: 0.04rem;
      box-sizing: border-box;
      margin-bottom: 0.2rem;
      padding: 0.1rem;
      .img-box {
        width: 0.95rem;
        height: 0.95rem;
        border-radius: 0.01rem;
        background: #fff;
        display: inline-block;
        text-align: center;
        vertical-align: middle;
        position: relative;
        .cancel-love {
          position: absolute;
          right: 0;
        }
        img {
          vertical-align: middle;
        }
      }
      .info-box {
        display: inline-block;
        color: #333;
        width: 2rem;
        div {
          margin-bottom: 0.05rem;
        }
      }
    }
  }
  .goods-love-list {
    margin-top: 0.2rem;
    .title {
      font-size: 0.2rem;
      margin-bottom: 0.1rem;
    }
    .item {
      width: 3.35rem;
      height: 1.15rem;
      background: rgba(255, 255, 255, 1);
      box-shadow: 0px 0px 0.16rem 0px rgba(214, 214, 214, 0.5);
      border-radius: 0.04rem;
      box-sizing: border-box;
      margin-bottom: 0.2rem;
      padding: 0.1rem;
      .img-box {
        width: 0.95rem;
        height: 0.95rem;
        border-radius: 0.01rem;
        background: #fff;
        display: inline-block;
        text-align: center;
        vertical-align: middle;
        position: relative;
        .cancel-love {
          position: absolute;
          right: 0;
        }
        img {
          vertical-align: middle;
        }
      }
      .info-box {
        display: inline-block;
        color: #333;
        width: 2rem;
        div {
          margin-bottom: 0.05rem;
        }
      }
    }
  }
  .scene-list {
    margin-top: 0.2rem;
    .title {
      font-size: 0.2rem;
      margin-bottom: 0.1rem;
    }
    .item {
      width: 3.35rem;
      height: 1.15rem;
      background: rgba(255, 255, 255, 1);
      box-shadow: 0px 0px 0.16rem 0px rgba(214, 214, 214, 0.5);
      border-radius: 0.04rem;
      box-sizing: border-box;
      margin-bottom: 0.2rem;
      padding: 0.1rem;
      .img-box {
        width: 0.95rem;
        height: 0.95rem;
        border-radius: 0.01rem;
        background: #fff;
        display: inline-block;
        text-align: center;
        vertical-align: middle;
        position: relative;
        .cancel-love {
          position: absolute;
          right: 0;
        }
        img {
          vertical-align: middle;
        }
      }
      .info-box {
        display: inline-block;
        color: #333;
        width: 2rem;
        div {
          margin-bottom: 0.05rem;
        }
        .to3d {
          // display: inline-block;
          width: 0.77rem;
          height: 0.24rem;
          line-height: 0.24rem;
          background: rgba(240, 240, 240, 1);
          border-radius: 0.12rem;
          color: #d8a163;
          font-size: 0.12rem;
          text-align: center;
          // margin-right: 0.1rem
        }
      }
    }
  }
}
</style>