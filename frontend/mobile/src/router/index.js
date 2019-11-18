import Vue from 'vue'
import Router from 'vue-router'
import Login from '@/pages/login'
import Home from '@/pages/home'
import AllTypes from '@/pages/all-types'
import MyTracks from '@/pages/track/my-tracks'
import TypeDetail from '@/pages/type-detail'
import Contact from '@/pages/contact'
import GoodsDetail from '@/pages/goods/goods-detail'
import Simulation from '@/pages/goods/simulation'
import simulationSelect from '@/pages/goods/simulation-select'
import simulationLove from '@/pages/goods/simulation-love'
import simulatiionIframe from '@/pages/goods/simulatiion-iframe'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Login',
      component: Login
    },
    {
      path: '/login',
      name: 'Login',
      component: Login
    },
    {
      path: "/home",
      name: 'Home',
      component: Home
    },
    {
      path: '/alltypes',
      name: 'AllTypes',
      component: AllTypes
    },
    {
      path: '/simulation',//模拟场景
      name: 'Simulation',
      component: Simulation
    },
    {
      path: '/simulatiion-iframe',//模拟场景返回展示
      name: 'simulatiionIframe',
      component: simulatiionIframe
    },
    {
      path: '/simulation-select',//模拟场景选择型号
      name: 'simulationSelect',
      component: simulationSelect
    },
    {
      path: '/simulation-love',//模拟场景选择型号
      name: 'simulationLove',
      component: simulationLove
    },
    {
      path: '/mytracks',
      name: 'MyTracks',
      component: MyTracks
    },
    {
      path: '/typedetail',
      name: 'TypeDetail',
      component: TypeDetail
    },
    {
      path: '/contact',
      name: 'Contact',
      component: Contact
    },
    {
      path: '/goods-detail',
      name: 'GoodsDetail',
      component: GoodsDetail
    }
  ]
})
