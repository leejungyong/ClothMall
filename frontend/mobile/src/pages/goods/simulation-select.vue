<template>
  <div>
    <div class="threeTips">请选择面墙纸生成形态</div>
    <select v-model="typeB" class="com-opt" @change="getSelected">
      <option value="0">平行拼接</option>
      <option value="1">落差拼接</option>
    </select>
     <div class="next active" @click="goNext">
      <div style=" text-align: center;width:1rem;height:.22rem;margin:0 auto" >选择对比墙纸</div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      typeB: 0
    };
  },
  methods: {
    //选择形态
    getSelected() {
      console.log(this.typeA);
    },
    goNext(){
      localStorage.setItem("goodsBspliceType",this.typeB);
        this.$router.push({path:'/simulation-love',query:{typeB:this.typeB}})
    }
  }
};
</script>
<style lang="less" scoped>
.threeTips {
  color: #333;
  font-size: 0.2rem;
  margin-top: 0.2rem;
}
.com-opt {
  border: #d8a163 1px solid;
  color: #999;
  font-size: 0.18rem;
  margin-top: 0.2rem;
  padding: 0.1rem;
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
</style>