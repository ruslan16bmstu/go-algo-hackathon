<template>
  <div class="flex space-x-0.5 bg-[#f2f2f2] h-[500px]">
    <div class="grow basis-1/2 bg-white overflow-y-auto">
      <template v-for="(stock, i) in profitStocks" :key="i">
        <router-link :to="'/'+stock.secId">
          <StockShortInfo
              class="share-in-list"
              :stock="stock"
          />
        </router-link>
      </template>
      <div v-if="profitStocks.length === 0" class="pad">Нет данных</div>
    </div>
    <div class="grow basis-1/2 bg-white overflow-y-auto">
      <template v-for="(stock, i) in nonProfitStocks" :key="i">
        <router-link :to="'/'+stock.secId">
          <StockShortInfo
              class="share-in-list"
              :stock="stock"
          />
        </router-link>
      </template>
      <div v-if="profitStocks.length === 0" class="pad">Нет данных</div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, onMounted, ref } from 'vue'
import { StockWithIndustryWithPredictionWithCandles } from '../../stores/stock'
import StockShortInfo from './StockShortInfo.vue'

const props = defineProps<{
  stockPredictions: Array<StockWithIndustryWithPredictionWithCandles>
}>()

const profitStocks = computed(() => props.stockPredictions.filter(s => s.delta > 0).sort((a, b) => b.delta - a.delta))
const nonProfitStocks = computed(() => props.stockPredictions.filter(s => s.delta <= 0).sort((a, b) => a.delta - b.delta))

</script>

<style scoped>
.share-in-list {
  @apply select-none hover:cursor-pointer hover:shadow-md hover:text-[#5142ab] transition duration-300 border-b;
}
</style>