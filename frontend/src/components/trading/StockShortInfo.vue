<template>
  <div class="flex items-center pad text-sm justify-between">
    <div class="flex items-center">
      <div class="font-bold mr-2 w-12">{{ stock.secId }}</div>
      <div>{{formatStockName(stock.name)}}</div>
    </div>
    <div class="flex space-x-4 items-center">
      <div class="text-black">{{ formatPrice(stock.price) }}</div>
      <div :class="{'green-bg': stock.delta > 0, 'red-bg': stock.delta <= 0}"
           class="flex justify-center w-[140px] py-0.5"
      >
        <span class="mr-0.5">{{ formatDelta(stock.delta) }}</span>
        <span class="font-bold">({{calcPrice(stock)}})</span>
      </div>
      <div>
        <CandlesWithPredictionChart class="w-[46px] h-[26px]" :prediction="prediction" :candles="stock.candles"
                                    :hide-interface="true" :no-interactions="true"
        />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import { StockWithIndustryWithPredictionWithCandles } from '../../stores/stock'
import { formatDate, getDateRange } from '../../utils/date'
import CandlesWithPredictionChart from './CandlesWithPredictionChart.vue'
import { formatDelta, formatPrice, formatStockName } from './format'

const props = defineProps<{
  stock: StockWithIndustryWithPredictionWithCandles
}>()

const calcPrice = (stock: StockWithIndustryWithPredictionWithCandles) => {
  const formattedPrice = formatPrice(stock.price)
  let len = formattedPrice.replace(',','').length
  len += +(!formattedPrice.includes(','))
  len -= +formattedPrice.startsWith('0,')

  let res = (stock.price * (1 + stock.delta / 100)).toPrecision(len).replace('.', ',')
  if (res.endsWith(',0')) {
    res = res.slice(0, res.length - 2)
  }

  return res
}

const prediction = computed(() => {
  if (!props.stock) {
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
          value: props.stock.stockPrice
        }
    )
  }

  return res
})
</script>

<style scoped>
</style>