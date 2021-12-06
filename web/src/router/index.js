import { createRouter, createWebHistory } from 'vue-router'

const routerHistory = createWebHistory()

const routes = [
  {
    path: '/dailyRecommend',
    name: 'DailyRecommend',
    component: () => import('@/views/DailyRecommendView.vue')
  },
  {
    path: '/queryStockView',
    name: 'queryStockView',
    component: () => import('@/views/QueryStockView.vue')
  },
  {
    path: '/indexData',
    name: 'indexData',
    component: () => import('@/views/DailyIndexDataView.vue')
  },
  {
    path: '/getStockHold',
    name: 'getStockHold',
    component: () => import('@/views/StockHoldView.vue')
  },
  {
    path: '/showSingleStockMarket',
    name: 'showSingleStockMarket',
    component: () => import('@/views/SingleStockMarket.vue')
  }
]

const router = createRouter({
  base: process.env.BASE_URL,
  history: routerHistory,
  routes
})

export default router
