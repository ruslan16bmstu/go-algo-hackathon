<template>
  <div class="flex items-center pad text-sm justify-between">
    <div class="flex items-center">
      <div class="font-bold mr-2 w-10">{{ stock.secId }}</div>
      <div>{{formatStockName(stock.name)}}</div>
    </div>
    <div class="flex space-x-4 items-center">
      <div class="text-black">{{ formatPrice(stock.price) }}</div>
      <div :class="{'green-bg': stock.delta > 0, 'red-bg': stock.delta <= 0}"
           class="flex justify-center w-20 py-0.5"
      >
        {{ formatDelta(stock.delta) }}
      </div>
      <div>
        <LineChartSVG :lines-data="linesData" :width="46" :height="26"/>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed } from 'vue'
import LineChartSVG from '../shared/LinesChartSVG.vue'
import { formatDelta, formatPrice, formatStockName } from './format'
import type { StockWithPrediction } from './types'

const props = defineProps<{
  stock: StockWithPrediction
  prediction: { x: number, y: number }[]
}>()

const linesData = computed(() => {
  const lastPoint = props.prediction[props.prediction.length - 1]
  const nextPoint = {
    x: lastPoint.x + props.prediction.length / 3,
    y: props.stock.delta > 0 ? lastPoint.y * 1.05 : lastPoint.y * 0.95
  }
  return [
    {
      points: props.prediction,
      width: 1,
      color: props.stock.delta > 0 ? 'green' : 'red',
    },
    {
      points: [lastPoint, nextPoint],
      width: 1,
      color: 'blue',
    }
  ]
})

</script>

<style scoped>
.green-bg {
  color: rgba(14, 181, 30);
  background-color: rgba(14, 181, 30, .1);
}

.red-bg {
  color: #e22;
  background-color: rgba(238, 34, 34, .1);
}
</style>