<template>
  <div class="flex space-x-0.5 bg-[#f2f2f2] min-h-[500px]">
    <div class="grow basis-1/2 bg-white">
      <template v-for="(stock, i) in profitStocks" :key="i">
        <router-link :to="'/'+stock.secId">
          <StockShortInfo
              class="share-in-list"
              :stock="stock"
              :prediction="stock.points"
          />
        </router-link>
      </template>
    </div>
    <div class="grow basis-1/2 bg-white">
      <template v-for="(stock, i) in nonProfitStocks" :key="i">
        <router-link :to="'/'+stock.secId">
          <StockShortInfo
              class="share-in-list"
              :stock="stock"
              :prediction="stock.points"
          />
        </router-link>
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import StockShortInfo from './StockShortInfo.vue'
import type { StockWithPrediction } from './types'

const props = defineProps<{
  stockPredictions: Array<StockWithPrediction>
}>()

const profitStocks = computed(() => props.stockPredictions.filter(s => s.delta > 0))
const nonProfitStocks = computed(() => props.stockPredictions.filter(s => s.delta <= 0))

</script>

<style scoped>
.share-in-list {
  @apply select-none hover:cursor-pointer hover:shadow-md hover:text-[#5142ab] transition duration-300 border-b;
}
</style>