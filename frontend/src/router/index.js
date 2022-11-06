import Vue from 'vue'
import VueRouter from 'vue-router'
import HomeView from '../views/HomeView.vue'
import DepositView from '../views/DepositView.vue'
import GetVcView from '../views/GetVcView.vue'
import VpWithdrawView from "@/views/VpWithdrawView";

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/deposit',
    name: 'deposit',
    component: DepositView
  },
  {
    path: '/get-vc',
    name: 'getVc',
    component: GetVcView
  },
  {
    path: '/vp-withdraw',
    name: 'vpWithdraw',
    component: VpWithdrawView
  }
]

const router = new VueRouter({
  routes
})

export default router
