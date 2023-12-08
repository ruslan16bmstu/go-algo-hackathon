<template>
  <div ref="chart"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { createChart, ISeriesApi } from 'lightweight-charts'
import { CandleData } from '../../api'

interface PredictionData {
  time: string
  value: number
}

const chart = ref<HTMLElement | null>(null)
let candleSeries: ISeriesApi<'Candlestick'> | null = null
let predictionSeries: ISeriesApi<'Line'> | null = null

const props = defineProps<{
  candles: CandleData[]
  prediction: PredictionData[]
}>()

// Функция для обновления графика при изменении данных
const updateChart = () => {
  if (candleSeries && predictionSeries) {
    candleSeries.setData(props.candles)
    predictionSeries.setData(props.prediction.map(point => ({
      time: point.time,
      value: point.value,
    })))
  }
}

const createStockChart = () => {
  if (chart.value) {
    const priceChart = createChart(chart.value, {
      width: chart.value.clientWidth,
      height: chart.value.clientHeight,
    })

    candleSeries = priceChart.addCandlestickSeries()
    predictionSeries = priceChart.addLineSeries()

    updateChart()
  }
}

onMounted(() => {
  createStockChart()
})

watch(() => [props.candles, props.prediction], () => {
  updateChart()
})
</script>