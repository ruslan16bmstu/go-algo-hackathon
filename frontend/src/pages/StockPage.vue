<template>
  <PageLayout>
    <div class="plate pad" v-if="stock">
      <div>
        <div class="flex items-center">
          <div class="flex items-center mr-2">
            <div class="flex items-center mr-1">
              <ArrowDownIcon class="scale-150" :fill="stock.delta > 0 ? '#0eb51eff' : '#e22'" :class="{'rotate-180': stock.delta > 0}"/>
              <div class="font-bold text-4xl">{{ formatPrice(stock.price) }}</div>
            </div>
            <div class="text-xs flex flex-col">
              <div class="font-bold"
                  :class="stock.delta > 0 ? 'green-bg' : 'red-bg'">{{ formatDelta(stock.delta) }}</div>
              <div>RUB</div>
            </div>
          </div>

          <div class="flex flex-col ml-2">
            <div class="text-xl leading-5 font-bold">{{ stock.name }}</div>
            <div class="text-sm opacity-50">{{ stock.secId }}</div>
          </div>
        </div>
      </div>

      <div class="mt-10 flex justify-center" v-if="stock.candles.length || prediction.length">
        <CandlesWithPredictionChart class="w-[800px] h-[500px]" :candles="stock.candles" :prediction="prediction"/>
      </div>
    </div>
  </PageLayout>
</template>

<script lang="ts" setup>
import { computed, onBeforeMount } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import CandlesWithPredictionChart from '../components/trading/CandlesWithPredictionChart.vue'
import { formatDelta, formatPrice } from '../components/trading/format'
import PageLayout from '../layouts/PageLayout.vue'
import ArrowDownIcon from '../assets/arrow-drop-down.svg'
import { useStockStore } from '../stores/stock'
import { formatDate, getDateRange } from '../utils/date'

const route = useRoute()
const router = useRouter()

const stockStore = useStockStore()
const stock = computed(() => stockStore.getStockBySecId(route.params.stockId as string))

onBeforeMount(() => {
  if (!stockStore.getStockBySecId(route.params.stockId as string)) {
    stockStore.loadOne(route.params.stockId as string).then(() => {
      if (!stock.value) {
        router.replace('/')
      }
    })
  }
})

const prediction = computed(() => {
  if (!stock.value) {
    return []
  }

  const res = []
  for (let i = 0; i <= 5; i++) {
    const [to, from] = getDateRange(30)
    const end = new Date(to)
    end.setDate(end.getDate() + i)
    res.push(
        {
          time: formatDate(end),
          value: stock.value.stockPrice
        }
    )
  }

  return res
})

</script>

<style scoped>

</style>