<template>
  <div class="sIframe">
    <iframe
      class="content"
      runat="server"
      src="http://www.baidu.com"
      frameborder="no"
      border="0"
      marginwidth="0"
      marginheight="0"
      scrolling="no"
      allowtransparency="yes"
    ></iframe>
    <div class="action">
      <div class="inner" @click="back">返回</div>
      <div class="inner" @click="collect">收藏</div>
    </div>
  </div>
</template>
<script>
import axios from "axios";
import qs from "qs";
export default {
  data() {
    return {
    };
  },
  methods: {
     back() {
      this.$router.push({path:'goods-detail',query:{id:JSON.parse(localStorage.getItem("goodA")).id}})
    },
     async collect(){
     const res = await axios.post(
        "/api",
        qs.stringify({
          cmd: "insertCollection",
          token: "Jh2044695",
          userid: localStorage.getItem("userid"),
          shopid: 1,
          goodsid: "",
          collectionType: "场景",
          scenesURL: this.$route.query.scenesURL
        })
      );
      if(res.data.success){
        alert("收藏成功！")
      }
    }


  }
 
};
</script>
<style lang="less" scoped>
.sIframe {
  width: 100%;
  min-height: 6.09rem;
  .action {
    z-index: 99;
    position: fixed;
    color: #d8a163;
    bottom: 0;
    left: 0;
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
  .content {
    width: 100%;
    min-height: 6.69rem;
  }
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
      .cancel-love {
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
</style>