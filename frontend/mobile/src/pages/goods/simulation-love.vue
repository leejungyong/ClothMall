<template>
  <div class="tracks">
    <div
      style="font-size:.12rem;font-weight:400;color:rgba(204,204,204,1);"
    >对比墙纸来自我的收藏，如需对比更多，请提前收藏哦</div>
    <div class="goods-love-list" v-if="goodsLove">
      <div class="title">卧室</div>
      <div class="item" v-for="(item,index) in goodsLove" :key="index" @click="selectB(item)">
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
      sceneLove: [] //收藏的场景
    };
  },
  components: {},
  created() {
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
          if (res.data.success) {
            // this.record=res.data.data.Record
            this.goodsLove = res.data.data.shopCollection;
            this.sceneLove = res.data.data.scenesCollection;
          }
        });
    },
    selectB(item) {
     localStorage.setItem("goodB",JSON.stringify(item));
     this.$router.push({name:"Simulation",query: {
          currentStep:3,
          isFromB:true
        }})
    },
    //退出登录
    logOut() {},
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
      delCollection({ token: "Jh2044695", userid: 1, collectionid: id }).then(
        res => {
          if (res.data.success) {
            console.log(res.data);
            console.log("取消收藏成功");
            this.fetch();
          }
        }
      );
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
    margin-top: 0.2rem;
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