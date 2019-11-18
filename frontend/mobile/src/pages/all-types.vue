<template>
    <div class="types">
        <div class="tab-box">
            <ul class="tab">
                <li class="type fl clearfix"  v-for='(tab,index) in tabs' :key="index" @click="changeTab(index)" :class="{active:currentIndex==index}">{{tab}}</li>
            </ul>
        </div>
        <div v-for="(item,index) in allClassList" :key="index" v-show="index==currentIndex">
            <div class="type-list">
                        <div class="item " v-for="(twoLevel,minindex) in item.children" :key="minindex" @click="toTypeDetail(item.menuname,twoLevel.id,twoLevel.menuname,twoLevel.num)">
                            <div class="style">{{twoLevel.menuname}}</div>
                            <div class="num">{{twoLevel.num}}种</div>
                        </div>
                    </div>
        </div>
       
    </div>
</template>

<script>
import axios from "axios";
import qs from "qs";
import {getAllClass} from '@/api/index.js'
export default {
    data(){
        return{
            tabs:['墙纸类','墙布类','其他类'],
            currentIndex:0,
            allClassList:[]
        }
    },
    created(){
        this.fetchData()
    },
    methods:{
        fetchData(){
            axios.post('/api',qs.stringify({shopid:1,token:"Jh2044695",cmd:"getAllClass"})).then(res=>{
                        console.log(res.data)
                        if(res.data.success){
                            this.allClassList=res.data.data;
                            console.log(this.allClassList);
                        }
                    })
        },
        //切换tab
        changeTab(index){
            this.currentIndex=index
        },
        toTypeDetail(oneLevelName,twoLevelId,twoLevelName,twoLevelNum){
            this.$router.push({name:'TypeDetail',query:{oneLevelName,twoLevelId,twoLevelName,twoLevelNum}})
        }
    }
}
</script>
<style lang="less" scoped>
.types{
    padding: 0.2rem 0.22rem;
    .tab{
        width:3.3rem;
        height:0.32rem;
        background:rgba(240,240,240,1);
        border-radius:0.04rem;
        margin: 0 auto;
        .type{
            width: 1.1rem;
            height: 0.32rem;
            line-height: 0.32rem;
            text-align: center;
            background:rgba(240,240,240,1);
            color: #D8A163;
            border-radius: 0.04rem;
        }
        .active{
            background: #D8A163;
            color: #fff;
        }
    }
    .type-list{
        margin-top: 0.2rem;
        display: flex;
        flex-wrap: wrap;
        justify-content: space-between;
        .item{
            display: inline-block;
            width: 1.58rem;
            height: 0.77rem;
            margin-bottom: 0.2rem;  
            background:linear-gradient(225deg,rgba(234,242,237,1) 0%,rgba(211,221,214,1) 100%);
            border-radius:0.04rem;
            box-sizing: border-box;
            padding: 0.1rem;
            .style{
                font-size:0.18rem;
                font-weight:500;
                color:rgba(0,0,0,1);
            }
            .num{
                font-size: 0.14rem;
            }
        }
        
    }
}
</style>