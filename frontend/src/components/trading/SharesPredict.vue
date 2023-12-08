<template>
  <div class="flex space-x-0.5 bg-[#f2f2f2] min-h-[500px]">
    <div class="grow basis-1/2 bg-white">
      <template v-for="(share, i) in profitShares" :key="i">
        <ShareShortInfo
            class="share-in-list"
            :share="share"
            :prediction="predictedExampleProfit"
        />
      </template>
    </div>
    <div class="grow basis-1/2 bg-white">
      <template v-for="(share, i) in nonProfitShares" :key="i">
        <ShareShortInfo
            class="share-in-list"
            :share="share"
            :prediction="predictedExampleLoose"
        />
      </template>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import ShareShortInfo from './ShareShortInfo.vue'

const shares = [
  {
    symbol: 'ВТБ',
    price: 0.02242,
    delta: 2.45
  },
  {
    symbol: 'Сбербанк',
    price: 265.32,
    delta: -0.84
  },
  {
    symbol: 'ВТБ',
    price: 0.02242,
    delta: 2.45
  },
  {
    symbol: 'Сбербанк',
    price: 265.32,
    delta: -0.84
  },
  {
    symbol: 'ВТБ',
    price: 0.02242,
    delta: 2.45
  },
  {
    symbol: 'Сбербанк',
    price: 265.32,
    delta: -0.84
  },
  {
    symbol: 'ВТБ',
    price: 0.02242,
    delta: 2.45
  },
  {
    symbol: 'Сбербанк',
    price: 265.32,
    delta: -0.84
  },
  {
    symbol: 'ВТБ',
    price: 0.02242,
    delta: 2.45
  },
  {
    symbol: 'Сбербанк',
    price: 265.32,
    delta: -0.84
  },
]

const profitShares = computed(() => shares.filter(s => s.delta > 0))
const nonProfitShares = computed(() => shares.filter(s => s.delta <= 0))

const predicExampleYProfit = [
  553,
  551,
  558,
  558,
  558.6,
  558.6,
  564,
  563.6,
  563.2,
  561.8,
  561.4,
  562.2,
  562.2,
  564,
  566.4,
  565.4,
  566.4,
  565,
  567,
  563.8,
  560.8,
  559,
  559.4,
  563,
  562,
  565,
  565.8,
  565.4,
  562.2,
  563.6,
  563,
  563,
  561.6,
  561.8,
  562,
  562.6,
  561.8,
  560,
  561.2,
  561.2,
  561.2,
  561.6,
  563.8,
  564.6,
  564.6,
  564.6,
  565.6,
  567,
  567.2,
  568.8,
  567.6,
  567.4,
  567,
  567.2,
  567.4,
  564.4,
  564.8,
  563,
  561.4,
  562.8,
  561.4,
  562.4,
  564,
  565.4,
  565,
  564,
  565,
  565.2,
  565,
  564.2,
  563.6,
  564.6,
  565.4,
  566.4,
  567.4,
  569,
  569.6,
  568.8,
  568.6,
  569.2,
  569,
  568.8,
  568.6,
  567.6,
  567.2,
  567.4,
  567.4,
  567,
  567.4,
  567.2,
  566.6,
  566.6,
  566.8,
  567.2,
  568.2,
  567.2,
  569.2,
  569.8,
  568.6,
  568.4,
  567.6,
  566.4,
  567.4,
  566.2,
  566.6,
  568.6
]

const predictedExampleProfit = predicExampleYProfit.map((y, x) => ({x, y}))
const predictedExampleLoose = predicExampleYProfit.reverse().map((y, x) => ({x, y}))
</script>

<style scoped>
.share-in-list {
  @apply select-none hover:cursor-pointer hover:shadow-md hover:text-[#5142ab] transition duration-300 border-b;
}
</style>