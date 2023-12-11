<template>
  <div ref="chart"></div>
</template>

<script setup lang="ts">
import { ref, onMounted, watch } from 'vue'
import { createChart, IChartApi, ISeriesApi } from 'lightweight-charts'
import { CandleData } from '../../api/moex'
import { PredictionData } from './types'


const chart = ref<HTMLElement | null>(null)
let candleSeries: ISeriesApi<'Candlestick'> | null = null
let predictionSeries: ISeriesApi<'Line'> | null = null

const props = defineProps<{
  candles: CandleData[]
  prediction: PredictionData[]
  hideInterface?: boolean
}>()

const priceChart = ref<IChartApi | null>(null)

// Функция для обновления графика при изменении данных
const updateChart = () => {
  if (candleSeries && predictionSeries) {
    if (props.hideInterface) {
      candleSeries.applyOptions({
        lastValueVisible: false,
        priceLineVisible: false
      })
      predictionSeries.applyOptions({
        lastValueVisible: false,
        priceLineVisible: false,
        crosshairMarkerVisible: false,
      })
    }

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
      handleScale: !props.hideInterface,
      handleScroll: !props.hideInterface,
      leftPriceScale: {
        visible: !props.hideInterface
      },
      rightPriceScale: {
        visible: !props.hideInterface
      },
      autoSize: true,
      timeScale: {
        visible: !props.hideInterface
      },
      grid: {
        horzLines: {
          visible: !props.hideInterface
        },
        vertLines: {
          visible: !props.hideInterface
        }
      },
      overlayPriceScales: {
        borderVisible: !props.hideInterface,
        ticksVisible: !props.hideInterface,
      },
      crosshair: {
        horzLine: {
          visible: !props.hideInterface
        },
        vertLine: {
          visible: !props.hideInterface
        }
      },
    })

    candleSeries = priceChart.value.addCandlestickSeries()
    predictionSeries = priceChart.value.addLineSeries({lineWidth: props.hideInterface ? 2 : 4})

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