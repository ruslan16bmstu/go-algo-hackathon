<template>
  <div class="flex items-center pad text-sm justify-between">
    <div class="font-bold">{{ share.symbol }}</div>
    <div class="flex space-x-4 items-center">
      <div class="text-black">{{ formatPrice(share.price) }}</div>
      <div :class="{'green-bg': share.delta > 0, 'red-bg': share.delta <= 0}"
           class="flex justify-center w-20 py-0.5"
      >
        {{ formatDelta(share.delta) }}
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

const props = defineProps<{
  share: {
    symbol: string
    price: number
    delta: number
  }
  prediction: { x: number, y: number }[]
}>()

const formatPrice = (price: number) => price.toPrecision().replace('.', ',')
const formatDelta = (delta: number) => (delta > 0 ? '+' : '') + formatPrice(delta) + '%'

const linesData = computed(() => {
  const lastPoint = props.prediction[props.prediction.length - 1]
  const nextPoint = {
    x: lastPoint.x + props.prediction.length / 3,
    y: props.share.delta > 0 ? lastPoint.y * 1.05 : lastPoint.y * 0.95
  }
  return [
    {
      points: props.prediction,
      width: 1,
      color: props.share.delta > 0 ? 'green' : 'red',
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