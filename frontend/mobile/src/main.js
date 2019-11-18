// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Mock from './mock'
import axios from 'axios'
import "swiper/css/swiper.min.css";
Vue.config.productionTip = false

// axios.defaults.baseURL='http://mockjs.com/api'//设置默认请求的url

// Vue.prototype.$axios=axios

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>'
})
