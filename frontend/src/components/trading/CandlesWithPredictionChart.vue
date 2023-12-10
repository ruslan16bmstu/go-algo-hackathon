<template>
  <div ref="chart"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { createChart, IChartApi, ISeriesApi } from 'lightweight-charts'
import { CandleData } from '../../api'
import { PredictionData } from './types'


const chart = ref<HTMLElement | null>(null)
let candleSeries: ISeriesApi<'Candlestick'> | null = null
let predictionSeries: ISeriesApi<'Line'> | null = null

const props = defineProps<{
  candles: CandleData[]
  prediction: PredictionData[]
}>()

const priceChart = ref<IChartApi | null>(null)

// Функция для обновления графика при изменении данных
const updateChart = () => {
  if (candleSeries && predictionSeries) {
    candleSeries.setData(props.candles)
    predictionSeries.setData(props.prediction.map(point => ({
      time: point.time,
      value: point.value,
    })))

    priceChart.value!.timeScale().fitContent()
  }
}

const createStockChart = () => {
  if (chart.value) {
    priceChart.value = createChart(chart.value, {
      width: chart.value.clientWidth,
      height: chart.value.clientHeight,
    })

    candleSeries = priceChart.value.addCandlestickSeries()
    predictionSeries = priceChart.value.addLineSeries({lineWidth: 4})

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