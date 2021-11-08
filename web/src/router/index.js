import { createRouter, createWebHistory } from 'vue-router'

const routerHistory = createWebHistory()

const routes = [
  {
    path: '/dailyRecommand',
    name: 'DailyRecommand',
    component: () => import('@/views/DailyRecommandView.vue')
  },
  {
    path: '/queryStockView',
    name: 'queryStockView',
    component: () => import('@/views/QueryStockView.vue')
  },
  {
    path: '/totalMarket',
    name: 'totalMarket',
    component: () => import('@/views/TotalMarket.vue')
  },
  {
    path: '/getStockHold',
    name: 'getStockHold',
    component: () => import('@/views/StockHoldView.vue')
  }
]

const router = createRouter({
  base: process.env.BASE_URL,
  history: routerHistory,
  routes
})

export default router
