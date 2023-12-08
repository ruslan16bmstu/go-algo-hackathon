<template>
  <div ref="chart"></div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
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

onMounted(() => {
  if (chart.value) {
    const priceChart = createChart(chart.value, {
      width: chart.value.clientWidth,
      height: chart.value.clientHeight,
    })

    // Создание свечного графика
    candleSeries = priceChart.addCandlestickSeries({
      upColor: 'green',
      downColor: 'red',
      borderVisible: true,
      wickVisible: true,
    })

    candleSeries.setData(props.candles)

    // Создание линии прогноза
    predictionSeries = priceChart.addLineSeries({
      color: 'blue',
      lineWidth: 1,
    })

    predictionSeries.setData(props.prediction.map(point => ({
      time: point.time,
      value: point.value,
    })))
  }
})
</script>