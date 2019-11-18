<template>
  <div class="simulation">
    <div class="tab-box">
      <span :class='[currentStep>=0? "tipActive" : "tip"]' >选择户型</span>
      <em :class='[currentStep>=1? "arrowActive" : "arrow"]'></em>
      <span :class='[currentStep>=1? "tipActive" : "tip"]' >选择房间</span>
      <em :class='[currentStep>=2? "arrowActive" : "arrow"]'></em>
      <span :class='[currentStep>=2? "tipActive" : "tip"]'>选择形态</span>
      <em :class='[currentStep>=3? "arrowActive" : "arrow"]'></em>
      <span :class='[currentStep>=3? "tipActive" : "tip"]'>选择对比</span>
    </div>
    <div class="box-contant">
      <div class="type-list" v-if="this.currentStep==0||this.currentStep==1">
        <div
          :style="{backgroundImage:'url('+item.pic+')'}"
          :class="['item',{'active':item.checked}]"
          v-for="(item,index) in allMenuList"
          :key="index"
          @click="selectRoom(item,index)"
        >
          <div v-show="item.checked" class="top"></div>
          <div class="bottom">{{item.name}}</div>
        </div>
      </div>
      <div v-if="this.currentStep==2" >
        <div class="threeTips">请选择面墙纸生成形态</div>
        <select v-model="goodsAspliceType " class="com-opt" @change="getSelected">
         <option value ="0">平行拼接</option>
         <option value ="1">落差拼接</option>
          </select>
      </div>
       <div class="type-list" v-if="this.currentStep==3">
        <div
        class="item"
        style="width:1.45rem;height:1.45rem;background-size: 100% 100%;"
          :style="{backgroundImage:'url('+this.goodAurl+')'}"
        >
        <div class="bottom" style="width:1rem">背景墙壁纸</div>
        </div>
         <div
        class="item add"
        style="width:1.45rem;height:1.45rem;background-size: 100% 100%;"
        @click="goAddCollect"
        v-if="!isFromB"
        >
        <div class="bottom" style="width:1.2rem">非背景墙壁纸</div>
        </div>
           <div
        class="item"
        style="width:1.45rem;height:1.45rem;background-size: 100% 100%;"
        @click="goAddCollect"
        :style="{backgroundImage:'url('+this.goodBurl+')'}"
        v-if="isFromB"
        >
        <div class="bottom" style="width:1.2rem">非背景墙壁纸</div>
        </div>
      </div>
    </div>
    <div :class="['next',{'active':this.canClicked}]" @click="goNext">
      <div style=" text-align: center;width:1rem;height:.22rem;margin:0 auto" >{{this.currentStep==3?"开始模拟场景":"下一步"}}</div>
    </div>
    <!-- 轮询遮罩 -->
    <div class="pollModal" v-if="isShowpollModal">
      <div class="loading">
        <span></span>
        <span></span>
        <span></span>
        <span></span>
        <span></span>
</div>
</div>  
  </div>
</template>

<script>
import axios from "axios";
import qs from "qs";
import { getAllClass } from "@/api/index.js";
export default {
  data() {
    return {
      isShowpollModal:false,
      tabs: ["墙纸类", "墙布类", "其他类"],
      currentIndex: 0, //0 1 2 3
      allMenuList: [],
      canClicked: false,
      lastIndex: 0, //ji'lu
      scenesmenuid1: 0, //户型id
      scenesmenuurl1: "", //户型url
      scenesmenuid2: "", //房间id
      scenesmenuurl2: "", //房间url
      goodsBid: "", //商品Bid
      goodBurl: localStorage.getItem('goodB')&&JSON.parse(localStorage.getItem('goodB')).compressPic, //商品Burl
      currentStep: this.$route.query.currentStep, //当前步骤
      goodsAspliceType :0,//A强制生成形态
      goodAurl:localStorage.getItem('goodA')&&JSON.parse(localStorage.getItem('goodA')).picList[0].compresspic,
      isFromB:this.$route.query.isFromB,//是否已经选择好了
    };
  },
  created() {
    this.fetchData();
    if(this.$route.query.isFromB){
      this.canClicked=true;
    }
  },
    beforeRouteUpdate(to,from,next){
        console.log(to,from,next)
        if(to.fullPath!=from.fullPath){
            next()
            // this.changeUser()
        }
    },
  methods: {
    fetchData() {
      axios
        .post(
          "/api",
          qs.stringify({
            token: "Jh2044695",
            cmd: "getMallScenesMenu",
            superid: 0
          })
        )
        .then(res => {
          if (res.data.success) {
            this.allMenuList = res.data.data.map(item => {
              var newItem = {
                superid: item.superid,
                children: item.children,
                id: item.id,
                name: item.name,
                pic: item.pic,
                checked: false
              };
              return newItem;
            });
            console.log(this.allMenuList);
          }
        });
    },
    //选择形态
    getSelected(){
      // localStorage.setItem('goodsAspliceType',this.goodsAspliceType);
      console.log(this.goodsAspliceType );
    },
    //点击下一步
    goNext() {
     //选择户型
      if(this.currentStep==0){
          if (!this.canClicked) {
        return;
      }
        axios
        .post(
          "/api",
          qs.stringify({
            token: "Jh2044695",
            cmd: "getMallScenesMenu",
            superid: this.scenesmenuid1
          })
        )
        .then(res => {
          if (res.data.success) {
            this.currentStep=parseInt(this.currentStep)+1;
           this.changeUrlStep();
            this.canClicked=false;
            this.lastIndex=0;
            this.allMenuList = res.data.data.map(item => {
              var newItem = {
                superid: item.superid,
                children: item.children,
                id: item.id,
                name: item.name,
                pic: item.pic,
                checked: false
              };
              return newItem;
            });
            return;
          }
        });
      }
      //选择房间
      if(this.currentStep==1){
           if (!this.canClicked) {
              return;
            }
             this.currentStep=parseInt(this.currentStep)+1;
           this.changeUrlStep();
             return;
      }
      //选择形态
      if(this.currentStep==2){
        this.currentStep=parseInt(this.currentStep)+1;
          this.canClicked=false;
          localStorage.setItem("goodsAspliceType",this.goodsAspliceType);
      }
      //开始模拟场景
       if(this.currentStep==3){
         
          if (!this.canClicked) {
              return;
            }
            this.isShowpollModal=true;
           //调用模拟场景的接口
            axios
        .post(
          "/api",
          qs.stringify({
            cmd:"insertScenesTask",
            token:"Jh2044695",
            shopid:"1",
             userid:localStorage.getItem("userid"),
            scenesmenuid1:JSON.parse(localStorage.getItem('simulationAParams')).id,
            scenesmenuurl1:JSON.parse(localStorage.getItem('simulationAParams')).pic,  
            scenesmenuid2 :JSON.parse(localStorage.getItem('simulationBParams')).id,
            scenesmenuurl2 :JSON.parse(localStorage.getItem('simulationBParams')).pic,
            goodsAid  :JSON.parse(localStorage.getItem('goodA')).id,
            goodsAurl  :JSON.parse(localStorage.getItem('goodA')).picList[0].compresspic,
            goodsAwidth:JSON.parse(localStorage.getItem('goodA')).width,  
            goodsAheight  :JSON.parse(localStorage.getItem('goodA')).height,
            goodsAspliceType  :JSON.parse(localStorage.getItem('goodsAspliceType')),
            goodsBid  :JSON.parse(localStorage.getItem('goodB')).id,
            goodsBurl  :JSON.parse(localStorage.getItem('goodB')).compressPic,
            goodsBwidth :JSON.parse(localStorage.getItem('goodB')).width,
            goodsBheight  :JSON.parse(localStorage.getItem('goodB')).height,
            goodsBspliceType  :JSON.parse(localStorage.getItem('goodsBspliceType'))
          })
        )
        .then(res => {
          if (res.data.success) {
            //轮询结果，获取结果
            // this.getScenesTaskURL(res.data.data);
            this.getScenesTaskURL(1);
          }
        });


      }
    },
    //获取模拟场景链接，接口
    getScenesTaskURL(id){
      var timer;
       axios.post(
          "/api",
          qs.stringify({
            token: "Jh2044695",
            cmd: "getScenesTaskURL",
            scenesTaskid : id
          })
        )
        .then(res=>{
             if(res.data.success&&res.data.data){
               this.isShowpollModal=false;
                clearTimeout(timer);
                //跳转到iframe页面
                this.$router.push({name:'simulatiionIframe'})
             }
             if(res.data.success&&!res.data.data){
                timer = setTimeout(() => {
                    this.getScenesTaskURL(id)
                }, 1000)   
             }
        })
    },
    //选择B型号
    goAddCollect(){
  this.$router.push({ path: "/simulation-select" });
    },
    changeUrlStep(){
  //改变链接带的步骤
            let query = this.$router.history.current.query;
            let path = this.$router.history.current.path;
             //对象的拷贝
            let newQuery = JSON.parse(JSON.stringify(query));
            newQuery.currentStep = this.currentStep;
            this.$router.push({ path, query: newQuery });
    },
    selectRoom(item, index) {
      console.log(item,index);
      //判断，只可打一个钩
      if (this.lastIndex != index) {
        this.$set(this.allMenuList[this.lastIndex], "checked", false);
      }
      //能不能点击
      this.canClicked = !item.checked;
      //改变样式
      this.$set(item, "checked", !item.checked);
      //记录点击
      this.lastIndex = index;
      //记录当前选择的
      if(this.currentStep==0){
       this.scenesmenuid1=item.id;
       this.scenesmenuurl1=item.pic;
       //存储当前的A图纸的信息
       localStorage.setItem("simulationAParams",JSON.stringify(item));
      }
      //记录房间号
       if(this.currentStep==1){
       this.scenesmenuid2=item.id;
       this.scenesmenuurl2=item.pic;
        localStorage.setItem("simulationBParams",JSON.stringify(item));
      }
    }
  }
};
</script>
<style lang="less" scoped>
.simulation {
   .pollModal{
    height: 100%;
    background: rgba(0, 0, 0, 0.4);
    z-index: 99;
    position: fixed;
    bottom: 0;
    left: 0;
    width: 100%;
     .loading{
            width: 1.5rem;
            height: .15rem;
            margin: 0 auto;
            margin-top:2.5rem;
        }
        .loading span{
            display: inline-block;
            width: .15rem;
            height: 100%;
            margin-right: .05rem;
            border-radius: 50%;
            background: #fff;
            -webkit-animation: load 1.04s ease infinite;
        }
        .loading span:last-child{
            margin-right: 0px; 
        }
        @-webkit-keyframes load{
            0%{
                opacity: 1;
            }
            100%{
                opacity: 0;
            }
        }
        .loading span:nth-child(1){
            -webkit-animation-delay:0.13s;
        }
        .loading span:nth-child(2){
            -webkit-animation-delay:0.26s;
        }
        .loading span:nth-child(3){
            -webkit-animation-delay:0.39s;
        }
        .loading span:nth-child(4){
            -webkit-animation-delay:0.52s;
        }
        .loading span:nth-child(5){
            -webkit-animation-delay:0.65s;
        }
   }
  .threeTips{
    color: #333;
    font-size: .20rem;
    margin-top: .2rem;
  }
  .com-opt {
    border: #d8a163 1px solid;
    color: #999;
    font-size: .18rem;
   margin-top:.2rem;
   padding: .1rem
}

  .next {
    position: fixed;
    bottom: 0;
    z-index: 99;
    width: 100%;
    height: 0.48rem;
    background: #b8b3b3;
    font-size: 0.16rem;
    color: #fff;
    line-height: 0.48rem;
    font-weight: 400;
  }
  .next.active {
    background: rgba(216, 161, 99, 1);
  }

  padding-top: 0.2rem;
  .tab-box {
    margin: 0 auto;
    padding: 0 0.1rem;
    width: 3.3rem;
    height: 0.32rem;
    background: rgba(240, 240, 240, 1);
    border-radius: 0.04rem;
    display: flex;
    justify-content: space-between;
    flex-wrap: wrap;

    .tipActive {
      color: #d8a163;
      font-size: 0.16rem;
      line-height: 0.32rem;
    }
    .tip {
      color: #999999;
      font-size: 0.14rem;
      line-height: 0.32rem;
    }
    .arrowActive {
      background: url("../../assets/icon/arrowActive.png") no-repeat;
      background-size: cover;
      width: 0.15rem;
      height: 0.08rem;
      position: relative;
      margin: -0.04rem;
      top: 50%;
    }
    .arrow {
      background: url("../../assets/icon/arrow.png") no-repeat;
      background-size: cover;
      width: 0.15rem;
      height: 0.08rem;
      position: relative;
      margin: -0.04rem;
      top: 50%;
    }
  }
  .box-contant {
    padding: 0 0.2rem;
  }
  .type-list {
    margin-top: 0.2rem;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
    .item.add{
      background-image:url('../../assets/icon/addpic.png');
      background-size: cover;
    }
    .item {
      display: inline-block;
      width: 1.58rem;
      height: 1.58rem;
      margin-bottom: 0.2rem;
      background: linear-gradient(
        225deg,
        rgba(234, 242, 237, 1) 0%,
        rgba(211, 221, 214, 1) 100%
      );

      border-radius: 0.04rem;
      box-sizing: border-box;
      padding: 0.1rem;
      position: relative;
      .top {
        background: url("../../assets/icon/select@2x.png") no-repeat;
        background-size: cover;
        width: 0.2rem;
        height: 0.2rem;
        position: absolute;
        top: 0.1rem;
        right: 0.1rem;
      }
      .bottom {
        font-size: 0.16rem;
        width: 0.79rem;
        height: 0.26rem;
        background: linear-gradient(
          90deg,
          rgba(255, 255, 255, 0) 0%,
          rgba(255, 255, 255, 1) 100%
        );
        font-weight: 400;
        right: 0;
        line-height: 0.26rem;
        bottom: 0;
        position: absolute;
        padding-right: 0.05rem;
        text-align: right;
      }
    }
    .item.active {
      border: 0.04rem solid rgba(216, 161, 99, 1);
    }
  }
}
</style>