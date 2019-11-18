<template>
    <div class="login">
         <div class="wel-text" style="font-size:0.24rem;">您好，</div>
         <div class="wel-text" style="font-size:0.2rem;">欢迎登陆×××</div>
         <div class="input-holder">
             <input class="tel" type="text" placeholder="手机号" v-model="phone" >
             <!-- <div v-if="errorMsg!=''">{{errorMsg}} </div> -->
             <div class="yzm-box">
                 <div class="yzm"><input type="text" placeholder="验证码" v-model="verifyCode"></div>
                 <button class="btn" @click="getVerifyCode" :disabled='disabled'>{{btnTitle}}</button>
             </div>
             <div class="tip">未注册手机验证后自动登录</div>
         </div>
         <button class="login-btn" @click="login">登录</button>
    </div>
</template>
<script>
import {userLogin} from '@/api/index.js'
import axios from 'axios';
import qs from 'qs';
export default {
    data(){
        return{
            phone:'',
            errorMsg:'',
            time:'',
            btnTitle:'获取验证码', //获取验证码按钮的字
            disabled:false,
            verifyCode:'',       //用户输入的验证码
            code:''             //接口请求返回的验证码
        }
    },
    created(){
        // userLogin({phoneNum:11,password:23}).then(res=>{
        //     console.log(res.data)
        // })
    },
    methods:{
        //点击获取验证码
        getVerifyCode(){
            if(this.validatePhone()){
                this.validateBtn()
                this.$axios.post('/code',{tpl_id:1,key:'dwwf',phone:this.phone}).then(res=>{
                    console.log(res)
                    if(res.data){
                      this.code=res.data.vecode
                    }
                })
            }
            
        },
         // 检查手机号码
        validatePhone(){
            if(this.phone==''){
                this.errorMsg='手机号码不能为空'
            }else if(!/^1[345678]\d{9}$/.test(this.phone)){
                this.errorMsg='请输入正确的手机号！'
            }else{
                this.errorMsg=''
                return true
            }
        },
        //倒计时
        validateBtn(){
            let time = 60;
            let timer = setInterval(() => {
            if(time == 0) {
                clearInterval(timer);
                this.disabled = false;
                this.btnTitle = "获取验证码";
            } else {
                this.btnTitle =time + '秒后重试';
                this.disabled = true;
                time--
            }
            },1000)
        },

        login(){
            // login().then((res)=>{
            //     console.log("login",res);
            // })
            if(this.verifyCode==''||this.phone==''){
                this.errorMsg='手机号或验证码不能为空'
            }else{
                  axios.post('/api',qs.stringify({phoneNum:this.phone,password:this.verifyCode,token:'Jh2044695',cmd:'userLogin'})).then(res=>{
                        console.log(res)
                        if(res.data.success){
                            localStorage.setItem("userid",res.data.data.userid)
                            localStorage.setItem('phonenum',this.phone)
                            this.$router.push('/home') 
                        }
                    })
            }
          
            
        }
    }
}
</script>

<style lang="less" scoped>
@import '../css/common.css';
.login{
    // text-align: center;
    padding: 0.6rem 0.3rem;
    .wel-text{
        color:rgba(51,51,51,1);
    }
    .input-holder{
        margin:0.4rem 0 0.35rem 0;
        .tel{
            width: 3.15rem;
            height: 0.4rem;
            border-bottom: 0.01rem solid #999;
            padding-left: 0.05rem;
        }
        .yzm-box{
            margin-top: 0.3rem;
              width: 3.15rem;
                height: 0.4rem;
                line-height: 0.4rem;
                border-bottom: 0.01rem solid #999;
            .yzm{
              
                // border-bottom: 0.01rem solid #999;
                padding-left: 0.05rem;
                display: inline-block;
                input{
                    width: 2rem;
                    // height: 0.39rem;
                    
                }
            }
            .btn{
            width:0.9rem;
            height:0.3rem;
            text-align: center;
            line-height: 0.3rem;
            background:#F0F0F0;
            border-radius:0.15rem;
            color: #D8A163;
            display: inline-block;
        }
        }
        .tip{
            font-size:0.12rem;
            color:rgba(204,204,204,1);
            margin-top:0.08rem;
        }
        
    }
    .login-btn{
        width:3.15rem;
        height:0.4rem;
        background:rgba(216,161,99,1);
        border-radius:0.04rem;
        color: #fff;
        margin-top:0.35rem;
        font-size: 0.16rem;
    }
}
</style>