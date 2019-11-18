<template>
  <div class="goods">
    <div class="background">
      <div class="current" @click="goFull" :style="{backgroundImage:'url('+curPic+')'}">
        <!-- <img :src="curPic"/> -->
        <div class="collect" @click.stop="goCollect">{{goodsObj.iscollection?"已收藏":"收藏"}}</div>
      </div>
      <div class="mini">
        <ul class="none" style="width: auto;overflow-x: auto;white-space: nowrap;">
          <li
            v-for="(item,index) in goodsObj.picList"
            :key="index"
            :class="[current==index?'active':'', 'img']"
            :style="{backgroundImage:'url('+item.compresspic +')'}"
            @click="changePic(index)"
          ></li>
        </ul>
      </div>
    </div>
    <div class="detail-box">
      <div>名称：{{goodsObj.goodsname}}</div>
      <div class="text">品牌：{{goodsObj.brand}}</div>
      <div class="text">风格：{{goodsObj.style}}</div>
      <div class="text">型号：{{goodsObj.picList[current].model}}</div>
      <div class="text">规格：{{goodsObj.unit}}</div>
      <div class="text">材质:{{goodsObj.material}}</div>
      <div class="text">产地:{{goodsObj.madein}}</div>
      <div class="text">价格:{{goodsObj.price}}/m2</div>
      <div style="margin-top:0.1rem;">实物位于机器{{goodsObj.machinedetail.machinename}}-编号{{goodsObj.machinedetail.slotnum}}</div>
    </div>
    <div class="btn-box">
      <!-- <span> -->
      <div class="border fl clearfix" @click="gomachineModal">实体机展示</div>
      <div style="width:2%;display:inline-block;">
        <span style="background-color:#fff;height:0.3rem;width:2px;display:inline-block;vertical-align: middle;"></span>
      </div>
      <div class="fr clearfix" @click="goSimulation">模拟场景</div>
      <!-- </span> -->
    </div>

    <div class="fullScreenModal" v-show="isShowfull" :style="{backgroundImage:'url('+curPic+')'}">
      <div class="action">
        <div class="inner" @click="back">返回</div>
        <div class="inner" @click.stop="goCollect">{{goodsObj.iscollection?"已收藏":"收藏"}}</div>
        <div class="inner">3D效果</div>
        <div class="inner">机器展示</div>
      </div>
    </div>
    <!-- 实体机展示弹窗 -->
    <div v-show="isShowMachine" class="machineModal">
      <div class="contant">
        <div class="tips">{{machineTips}}</div>
        <div class="confirm" @click="confirm">确定</div>
      </div>
    </div>
    <!-- 收藏弹窗 -->
    <div v-show="isShowCollectModal" class="collectModal">
      <div class="contant">
        <div class="title">
          选择分类
          <div class="colorTop"></div>
        </div>

        <div class="selectContant mTop5" v-for="(item,index) in collectData " :key='index'>
          <span class="sort">{{item.title}}</span>
          <span
            :class="['checkContant fr',{'checked':item.checked}]"
            @click="collectSelect(item,index)"
          ></span>
        </div>

        <!-- 按钮 -->

        <div class="bottomButton">
          <div :class="[{'active':this.cancleClicked},'button']" @click="cancleCollcet">取消</div>
          <div :class="[{'active':this.confirmClicked},'button']" @click="confirmCollcet">确定</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import qs from "qs";

export default {
  data() {
    return {
      isShowfull: false, //查看大图
      isShowMachine: false, //实体机展示弹窗
      isShowCollectModal: false, //收藏弹窗展示
      cancleClicked: false, //收藏取消按钮
      confirmClicked: false, //收藏确认按钮
      collectValue: "", //收藏的值
      machineTips: "", //实体机展示获取值
      collectData: [
        {
          title: "客厅",
          value: 0,
          checked: false
        },
        {
          title: "卧室",
          value: 1,
          checked: false
        },
        {
          title: "餐厅",
          value: 2,
          checked: false
        },
        {
          title: "书房",
          value: 3,
          checked: false
        },
        {
          title: "儿童房",
          value: 4,
          checked: false
        }
      ],
      lastIndex: 0, //上一个被选择的
      selectTitle: "",
      goodsObj: {
        id: "",
        goodsname: "",
        brand: "",
        style: "",
        unit: "",
        material: "",
        madein: "",
        price: "",
        machineid: "",
        machineadder: "",
        picList: [{ model: "'" }]
      },
      goodsid: "", //商品id
      smallpicList: null, //缩略图数组
      current: 0, //当前选择的缩略图的index
      curPic: "", //当前的小图
      recordList: [] //浏览记录数组
    };
  },
  created() {
    this.goodsid = this.$route.params.id;
    this.fetch();
  },
  mounted() {},
  methods: {
    fetch() {
      axios
        .post(
          "/api",
          qs.stringify({
            goodsid: this.$route.query.id,
            token: "Jh2044695",
            cmd: "getGoodsDetail",
            userid: localStorage.getItem("userid")
          })
        )
        .then(res => {
          console.log(res.data.data);
          if (res.data.success) {
            this.goodsObj = res.data.data;
            this.smallpicList = res.data.data.picList.map(i => {
              return i.compresspic;
            });
            this.curPic = res.data.data.picList[0].compresspic;
            this.saveRecord();
          }
        });
    },
    //保存浏览记录
    saveRecord() {
      if (localStorage.getItem("record") == null) {
        this.recordList.push(this.goodsObj);
        localStorage.setItem("record", JSON.stringify(this.recordList));
      } else {
        let record = JSON.parse(localStorage.getItem("record")),
          flex = false,
          oindex = null;

        //判断历史记录中是不是已经存在该条记录
        record.map((item, index) => {
          if (item.id == this.goodsObj.id) {
            flex = true;
            oindex = index;
          }
        });
        if (flex) {
          //如果记录已存在,移除
          this.recordList.splice(oindex, 1);
        }
        this.recordList.push(this.goodsObj);
        localStorage.setItem("record", JSON.stringify(this.recordList));
      }
    },
    login() {
      this.$router.push("/full");
    },
    goFull() {
      this.isShowfull = true;
    },
    //模拟场景展示
    goSimulation() {
      this.$router.push({
        name: "Simulation",
        query: {
          currentStep:0
        }
      });
      localStorage.setItem("goodA",JSON.stringify(this.goodsObj));
    },
    //实体机展示
    async gomachineModal() {
      //获取本地ip
      const ip = await axios.post("/cityjson?ie=utf-8");
      const cip = ip.data
        .split("=")[1]
        .substr(0, ip.data.split("=")[1].length - 1);
      const result = await axios.post(
        "/api",
        qs.stringify({
          userip: JSON.parse(cip).cip,
          token: "Jh2044695",
          cmd: "machineCheck"
        })
      );
      if (!result.data.success) {
        this.machineTips = result.data.errmsg;
        this.isShowMachine = true;
      }
      //增加机器展示任务
      if (result.data.success) {
        const insertMachineTaskRes = await axios.post(
          "/api",
          qs.stringify({
            machineid: this.goodsObj.machineid,
            userid: localStorage.getItem("userid"),
            shopid: 1,
            goodsid: this.goodsObj.id,
            slotid: this.goodsObj.machineadder,
            token: "Jh2044695",
            cmd: "insertMachineTask"
          })
        );
        if (insertMachineTaskRes.data.success) {
          this.machineTips = insertMachineTaskRes.data.errmsg;
        }
      }
    },
    back() {
      this.isShowfull = false;
    },
    confirm() {
      this.isShowMachine = false;
    },
    async goCollect() {
      //收藏
      if(!this.goodsObj.iscollection){
      this.isShowCollectModal = true;
      this.cancleClicked = false;
      this.confirmClicked = false;
      }
      //取消收藏
      if(this.goodsObj.iscollection){
     const res = await axios.post(
        "/api",
        qs.stringify({
          cmd: "delCollection",
          token: "Jh2044695",
          collectionid: this.goodsObj.collectionid,
        })
      );
      if (res.data.success) {
        this.fetch();
      }
      }
    },
    cancleCollcet() {
      this.cancleClicked = true;
      this.isShowCollectModal = false;
    },
    async confirmCollcet() {
      this.confirmClicked = true;
      const res = await axios.post(
        "/api",
        qs.stringify({
          cmd: "insertCollection",
          token: "Jh2044695",
          userid: localStorage.getItem("userid"),
          shopid: 1,
          goodsid: this.goodsObj.id,
          collectionType: this.collectValue,
          scenesURL: ""
        })
      );
      if (res.data.success) {
        this.isShowCollectModal = false;
        this.fetch();
      }
    },
    changePic(index) {
      this.current = index;
      this.curPic = this.smallpicList[index];
    },
    //选择
    collectSelect(item, index) {
      this.collectValue = item.title;
      if (this.lastIndex != index) {
        this.collectData = [
          {
            title: "客厅",
            value: 0,
            checked: false
          },
          {
            title: "卧室",
            value: 1,
            checked: false
          },
          {
            title: "餐厅",
            value: 2,
            checked: false
          },
          {
            title: "书房",
            value: 3,
            checked: false
          },
          {
            title: "儿童房",
            value: 4,
            checked: false
          }
        ];
      }
      this.$set(
        this.collectData[index],
        "checked",
        !this.collectData[index].checked
      );
      this.lastIndex = index;
      console.log(this.collectData);
    }
  }
};
</script>
<style lang="less" scoped>
.goods {
  .background {
    .current {
      background: #d8a163;
      height: 3.07rem;
      padding: 0.1rem 0.1rem 0 3.21rem;
      background-position: center center;
      background-repeat: no-repeat;
      background-size: cover;
      .collect {
        width: 0.44rem;
        height: 0.24rem;
        line-height: 0.24rem;
        color: #d8a163;
        background: rgba(255, 255, 255, 0.8);
        border-radius: 1px;
        text-align: center;
        position: absolute;
        z-index: 98;
      }
    }
    .mini {
      height: 0.68rem;
      background: rgba(255, 255, 255, 0.9);
      padding: 0.1rem;
      .img {
        width: 0.48rem;
        height: 0.48rem;
        margin-left: 0.1rem;
        display: inline-block;
        // background-position：center;
        background-size: cover;
        background-position: center center;
        background-repeat: no-repeat;
      }
      .active {
        border: 1px solid rgba(216, 161, 99, 1);
      }
      .none::-webkit-scrollbar {
        display: none;
      }
      .img:first-child {
        margin-left: 0;
      }
    }
  }
  .detail-box {
    padding: 0.1rem 0.2rem;
    width: 100%;
    .text {
      display: inline-block;
      width: 49%;
      margin-top: 0.1rem;
    }
  }
  .btn-box {
    width: 100%;
    height: 0.48rem;
    line-height: 0.48rem;
    position: fixed;
    bottom: 0;
    left: 0;
    background: #d8a163;
    div {
      width: 49%;
      color: #fff;
      text-align: center;
    }
    // .border::after {
    //   content: "|";
    //   height: 0.25rem;
    //   width: 0.01rem;
    //   background: #fff;
    //   color: #fff;
    // }
  }
  .fullScreenModal {
    background-position: center center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    z-index: 999;
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
    height: 100%;
    .action {
      color: #d8a163;
      margin: 4.96rem 0.06rem 0 2.95rem;
      .inner {
        width: 0.74rem;
        height: 0.24rem;
        text-align: center;
        line-height: 0.24rem;
        background: rgba(255, 255, 255, 1);
        margin-bottom: 0.15rem;
      }
    }
  }
  .collectModal,
  .machineModal {
    height: 100%;
    background: rgba(0, 0, 0, 0.4);
    z-index: 999;
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
  }
  .machineModal {
    .contant {
      width: 2.95rem;
      height: 2.84rem;
      background: rgba(255, 255, 255, 1);
      padding: 0.75rem 0.3rem 0.4rem;
      margin: 0 auto;
      margin-top: 1.59rem;
    }
    .tips {
      color: #333333;
      font-size: 0.2rem;
      width: 2.34rem;
      height: 0.56rem;
      text-align: center;
      font-weight: 400;
      margin: 0 auto;
    }
    .confirm {
      width: 1.37rem;
      height: 0.4rem;
      border-radius: 0.01rem;
      border: 1px solid rgba(216, 161, 99, 1);
      color: #d8a163;
      line-height: 0.28rem;
      font-size: 0.2rem;
      text-align: center;
      margin: 0 auto;
      margin-top: 0.73rem;
    }
  }
  .collectModal {
    .mTop5 {
      margin-top: 0.05rem;
    }
    .contant {
      width: 2.95rem;
      height: 4.26rem;
      background: rgba(255, 255, 255, 1);
      border-radius: 0.04rem;
      margin: 0 auto;
      margin-top: 1.35rem;
      padding: 0.4rem 0.3rem;
      .title {
        font-size: 0.24rem;
        font-weight: 400;
        color: rgba(51, 51, 51, 1);
        line-height: 0.33rem;
        text-align: center;
        margin-bottom: 0.18rem;
      }
      .colorTop {
        width: 1.12rem;
        height: 0.1rem;
        margin: 0 auto;
        margin-top: -0.09rem;
        background: rgba(255, 231, 204, 1);
      }
      .selectContant {
        width: 2.35rem;
        height: 0.4rem;
        box-shadow: 0px 0px 0px 0px rgba(214, 214, 214, 0.5);
        border-bottom: 1px solid rgba(214, 214, 214, 0.5);
        padding: 0.08rem 0.2rem;
        .sort {
          font-size: 0.18rem;
          font-weight: 400;
          color: rgba(51, 51, 51, 1);
          line-height: 0.22rem;
        }

        .checkContant {
          width: 0.2rem;
          height: 0.2rem;
          border-radius: 1px;
          border: 1px solid rgba(151, 151, 151, 1);
          display: flex;
          margin-top: 2px;
          align-items: center; /*垂直居中*/
        }
        .checkContant.checked::before {
          width: 0.06rem;
          height: 0.14rem;
          border-color: #d8a163;
          border-style: solid;
          border-width: 0 0.05rem 0.05rem 0;
          transform: rotate(45deg);
          content: "";
          display: inline-block;
          position: relative;
          left: 0.04rem;
          top: -0.02rem;
        }
        // .checked::active {
        //   width: 0.2rem;
        //   height: 0.2rem;
        //   border-color: #009933;

        //   border-style: solid;

        //   border-width: 0 3px 5px 0;

        //   transform: rotate(45deg);
        //   margin-top: 2px;
        //   align-items: center; /*垂直居中*/
        // }
        // .advice {
        //   height: 12px;
        //   width: 12px;
        //   display: inline-block;
        //   background-image: url("https://caiyunupload.b0.upaiyun.com/newweb/imgs/icon-unchecked.png");
        //   background-repeat: no-repeat;
        //   background-position: center;
        //   vertical-align: middle;
        //   margin-top: -4px;
        // }
      }
      .bottomButton {
        margin-top: 0.3rem;
        display: flex;
        justify-content: space-between;
      }
      .button {
        font-size: 0.18rem;
        color: rgba(216, 161, 99, 1);
        line-height: 0.4rem;
        width: 1.17rem;
        height: 0.4rem;
        text-align: center;
      }
      .active {
        border: 1px solid rgba(216, 161, 99, 1);
      }
    }
  }
}
</style>