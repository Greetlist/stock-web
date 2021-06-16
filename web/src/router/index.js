import Vue from 'vue'
import VueRouter from 'vue-router'

Vue.use(VueRouter)

const routes = [
  {
    path: '/dailyOverView',
    name: 'DailyOverView',
    component: () => import('@/views/DailyOverView.vue')
  },
  {
    path: '/queryStockView',
    name: 'queryStockView',
    component: () => import('@/views/QueryStockView.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
