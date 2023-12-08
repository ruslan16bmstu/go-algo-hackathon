<template>
  <PageLayout>
    <div class="plate pad">
      <div>
        <div v-if="stock" class="flex items-center">
          <div class="flex items-center mr-2">
            <div class="flex items-center mr-1">
              <!--            <ArrowIcon/>-->
              <div class="font-bold text-4xl">{{ formatPrice(stock.price) }}</div>
            </div>
            <div class="text-xs flex flex-col">
              <div>&nbsp;</div>
              <!--            <div>+0,77%</div>-->
              <div>RUB</div>
            </div>
          </div>

          <div class="flex flex-col ml-2">
            <div class="text-xl leading-5 font-bold">{{ stock.name }}</div>
            <div class="text-sm opacity-50">{{ stock.secId }}</div>
          </div>
        </div>
      </div>

      <div class="mt-10 flex justify-center" v-if="candles.length">
        <CandlesWithPredictionChart class="w-[800px] h-[500px]" :candles="candles" :prediction="[]"/>
      </div>
    </div>
  </PageLayout>
</template>

<script lang="ts" setup>
import NProgress from 'nprogress'
import { computed, onBeforeMount, onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { CandleData, getStockCandles } from '../api'
import CandlesWithPredictionChart from '../components/trading/CandlesWithPredictionChart.vue'
import { formatPrice, formatStockName } from '../components/trading/format'
import PageLayout from '../layouts/PageLayout.vue'
import ArrowIcon from '../assets/arrow-drop-down.svg'
import { useStockStore } from '../stores/stock'
import { getDateRange } from '../utils/date'

const route = useRoute()
const router = useRouter()

const stockStore = useStockStore()
const stock = computed(() => stockStore.getStockBySecId(route.params.stockId as string))
const candles = ref<CandleData[]>([])

onBeforeMount(() => {
  if (!stockStore.getStockBySecId(route.params.stockId as string)) {
    router.replace('/')
  }
})

onMounted(() => {
  if (stock.value) {
      NProgress.start()
    const [to, from] = getDateRange(90)
    getStockCandles(stock.value.secId, from, to).then((res) => {
      candles.value = res
      NProgress.done()
    })
  }
})

</script>

<style scoped>

</style>