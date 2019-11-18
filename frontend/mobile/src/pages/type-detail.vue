<template>
  <div class="type-detail">
    <div class="top-line">
      <div class="num">{{total}}种</div>
      <div class="right fr clearfix" @click="changeType">
        <span>{{oneLevelName}}-{{twoLevelName}}</span>
        <img src="../assets/icon/icon_xiala copy.png" alt />
      </div>
    </div>
    <div class="list">
      <div
        class="item"
        v-for="(item,index) in list"
        :key="index"
        :style="{backgroundImage:'url('+item.smallpic+')'}"
        @click="toGoodsDetail(item.id)"
      >
        <!-- <span class="t-left">中式风格</span> -->
        <span class="t-right fr clearfix">
          <img src="../assets/icon/icon_3d.png" alt />
        </span>
        <div class="bottom">
          <img src="../assets/icon/icon_huo.png" alt />
          <span>{{item.clicknum}}</span>
          <span class="fr clearfix">{{item.colornum}}种</span>
        </div>
      </div>
    </div>

    <div class="mask" v-if="showMask">
      <div class="box">
        <div style="padding-left:.2rem;">
          <div class="block" v-for="(item,index) in typeList" :key="index">
            <div class="text">{{item.menuname}}</div>
            <div class="typelist">
              <div
                v-for="(type,minindex) in item.children"
                :key="minindex"
                :class="[{'active':type.checked},'item']"
                @click="select(index,minindex)"
              >{{type.menuname}}</div>
            </div>
          </div>
          <!-- <div class="block">
            <div class="text">墙布类</div>
            <div class="typelist">
              <div class="item" v-for="(type,index) in qblist" :key="index">{{type}}</div>
            </div>
          </div>
          <div class="block">
            <div class="text">其他类</div>
            <div class="typelist">
              <div class="item" v-for="(type,index) in otherlist" :key="index">{{type}}</div>
            </div>
          </div>-->
        </div>
        <div class="btn">
          <div class="cancel fl clearfix" @click="cancel">取消</div>
          <div class="sure fr clearfix" @click="sure">确定</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import qs from "qs";
import { getGoods, getAllClass } from "@/api/index.js";
let originTypelist=[]
export default {
  data() {
    return {
      total: "",
      list: [],
      showMask: false,
      oneLevelName: "",
      twoLevelId: "",
      twoLevelName: "",
      typeList: [],
      current: null,
      menuid: "", //请求接口时的参数
      goodsid: null,
      qzlist: [],
      qblist: [],
      otherlist: [],
      preindex:null,
      preminindex:null
    };
  },
  created() {
    // console.log(this.$route.params)
    let params = this.$route.query;
    this.oneLevelName = params.oneLevelName;
    this.twoLevelId = params.twoLevelId;
    this.twoLevelName = params.twoLevelName;

    this.fetch();
  },
  methods: {
    fetch() {
      axios
        .post(
          "/api",
          qs.stringify({
            menuid: this.twoLevelId,
            token: "Jh2044695",
            cmd: "getGoods"
          })
        )
        .then(res => {
          console.log(res);
          if (res.data.success) {
            (this.total = res.data.data.total),
              (this.list = res.data.data.goodslist);
          }
        });
    },
    //下拉菜单选择
    changeType() {
      this.showMask = true;
      console.log("gggg");
      axios
        .post(
          "/api",
          qs.stringify({ shopid: 1, cmd: "getAllClass", token: "Jh2044695" })
        )
        .then(res => {
          if (res.data.success) {
            let typelist = res.data.data;
            this.typeList = res.data.data;

            for (let i = 0; i < typelist.length; i++) {
              this.typeList[i].children =  this.typeList[i].children&&typelist[i].children.map(
                item => {
                  item.checked = false;
                  return item;
                }
              );
            }
            // console.log(this.typeList);
             originTypelist=this.typeList
          }
        });
    },
    //商品详情
    toGoodsDetail(id) {
      this.$router.push({ name: "GoodsDetail", query: { id } });
    },
    //选中某二级菜单
    select(index, minindex) {
      
      this.$forceUpdate();
      // this.$set()
      if(this.preindex!=null&&this.preminindex!=null){
        this.$set(
        this.typeList[this.preindex].children[this.preminindex],
        `checked`,
        !this.typeList[this.preindex].children[this.preminindex].checked
      );
      }
       
      this.$set(
        this.typeList[index].children[minindex],
        `checked`,
        !this.typeList[index].children[minindex].checked
      );
      this.preindex=index
      this.preminindex=minindex
      this.twoLevelId = this.typeList[index].children[minindex].id;
    },
    cancel() {
      this.showMask = false;
    },
    sure() {
      this.showMask = false;

      this.fetch();
    }
  }
};
</script>

<style lang="less" scoped>
* {
  margin: 0;
  padding: 0;
  outline: 0;
  border: 0;
}
.type-detail {
  .top-line {
    width: 100%;
    height: 0.48rem;
    line-height: 0.48rem;
    color: #000;
    padding: 0 0.2rem;
    box-shadow: 0 -0.01rem 0 0 rgba(210, 210, 210, 0.5);
    border-bottom: 1px solid #D2D2D2;
    .num {
      width: 1rem;
      display: inline-block;
    }
    .right {
    }
  }
  .list {
    padding: 0.2rem;
    display: flex;
    justify-content: space-between;
    flex-wrap: wrap;
    .item {
      display: inline-block;
      width: 1.58rem;
      height: 1.78rem;
      border-radius: 0.04rem;
      box-shadow: 0 0 0.16rem 0 rgba(214, 214, 214, 0.5);
      background: #d8a163;
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
  .mask {
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.6);
    position: fixed;
    top: 0;
    left: 0;
    .box {
      width: 100%;
      height: 5.74rem;
      background: #fff;
      opacity: 1;
      padding-top: 0.1rem;
      position: absolute;
      top: 0;
      left: 0;
      .block {
        margin-bottom: 0.2rem;
      }
      .text {
        font-size: 0.2rem;
      }
      .typelist {
        display: flex;
        flex-wrap: wrap;
        .item {
          margin-top: 0.1rem;
          width: 1.05rem;
          height: 0.32rem;
          line-height: 0.32rem;
          text-align: center;
          background: rgba(240, 240, 240, 1);
          border-radius: 0.02rem;
          color: #333;
          margin-right: .1rem;
        }
        .active {
          background: #d8a163;
        }
      }
      .btn {
        height: 0.4rem;
        width: 100%;
        line-height: 0.4rem;
        background: #fff;
        position: absolute;
        bottom: 0;
        left: 0;
        .cancel {
          text-align: center;
          width: 50%;
          font-size: 0.18rem;
          background: #fff;
          color: #d8a163;
        }
        .sure {
          text-align: center;
          width: 50%;
          font-size: 0.18rem;
          color: #fff;
          background: #d8a163;
        }
      }
    }
  }
}
</style>