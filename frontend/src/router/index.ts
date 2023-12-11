import { createRouter, createWebHistory } from 'vue-router'
import StockPage from '../pages/StockPage.vue'
import StocksPage from '../pages/StocksPage.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: StocksPage
    },
    {
      path: '/:stockId',
      component: StockPage
    },
    {path: "/:pathMatch(.*)*", redirect: '/'}
  ]
})

export default router