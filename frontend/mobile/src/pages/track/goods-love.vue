<template>
        <div class="list">
            <div class="title">卧室</div>
            <div class="item" v-for="(item,index) in list" :key="index">
                <div class="img-box">
                    <img class="cancel-love fr clearfix" src="../../assets/icon/icon_love_1.png" alt="">
                    <img src="" alt="">
                </div>
                <div class="info-box fr clearfix ">
                    <div>名称：金伯利金伯利金伯利</div>
                    <div>型号：JBL-1302A</div>
                    <div>规格：450*550</div>
                    <div>价格：88/m2</div>
                </div>
            </div>
        </div>
</template>
<script>
export default {
    data(){
        return{
              list:[1,1,1,1,1,1,1,1,1]
        }
    },
    watch:{
        items(newinfo){
            console.log(newinfo)
        }
    }
}
</script>
<style lang="less" scoped>
    .list{
        margin-top: 0.2rem;
        .title{
            font-size: 0.2rem;
            margin-bottom: 0.1rem;
        }
        .item{
            width:3.35rem;
            height:1.15rem;
            background:rgba(255,255,255,1);
            box-shadow:0px 0px 0.16rem 0px rgba(214,214,214,0.5);
            border-radius:0.04rem;
            box-sizing: border-box;
            margin-bottom: 0.2rem;
            padding: 0.1rem; 
            .img-box{
                width:0.95rem;
                height:0.95rem;
                border-radius:0.01rem;
                background: #fff;
                display: inline-block;
                text-align: center;
                vertical-align: middle;
                .cancel-love{

                }
                img{
                    vertical-align: middle;
                }
            }
            .info-box{
                display: inline-block;
                color: #333;
                width: 2rem;
                div{
                    margin-bottom: 0.05rem;
                }

            }
        }
    }
</style>