<template>
  <div ref="chart"></div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { createChart, ISeriesApi } from 'lightweight-charts'

// Типы данных для свечей и прогноза
interface CandleData {
  time: number;
  open: number;
  high: number;
  low: number;
  close: number;
}

interface PredictionData {
  time: number;
  value: number;
}

const chart = ref<HTMLElement | null>(null)
let candleSeries: ISeriesApi<'Candlestick'> | null = null
let predictionSeries: ISeriesApi<'Line'> | null = null

onMounted(() => {
  if (chart.value) {
    const candleData: CandleData[] = [
      {close: 4976.16, high: 4977.99, low: 4970.12, open: 4972.89, time: 1587660000000},
      {close: 4977.33, high: 4979.94, low: 4971.34, open: 4973.20, time: 1587660060000},
      {close: 4977.93, high: 4977.93, low: 4974.20, open: 4976.53, time: 1587660120000},
      {close: 4966.77, high: 4968.53, low: 4962.20, open: 4963.88, time: 1587660180000},
      {close: 4961.56, high: 4972.61, low: 4961.28, open: 4961.28, time: 1587660240000},
      {close: 4964.19, high: 4964.74, low: 4961.42, open: 4961.64, time: 1587660300000},
      {close: 4968.93, high: 4972.70, low: 4964.55, open: 4966.96, time: 1587660360000},
      {close: 4979.31, high: 4979.61, low: 4973.99, open: 4977.06, time: 1587660420000},
      {close: 4977.02, high: 4981.66, low: 4975.14, open: 4981.66, time: 1587660480000},
      {close: 4985.09, high: 4988.62, low: 4980.30, open: 4986.72, time: 1587660540000},
      {close: 4985.09, high: 4988.62, low: 4980.30, open: 4986.72, time: 1587660540000},
    ] // Замените этот массив своими данными для свечей
    const predictionData: PredictionData[] = [
      {time: 1587660940000, value: 5000},
      {time: 1587661340000, value: 5040},
    ] // Замените этот массив своими данными для прогноза

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

    candleSeries.setData(candleData)

    // Создание линии прогноза
    predictionSeries = priceChart.addLineSeries({
      color: 'blue',
      lineWidth: 1,
    })

    predictionSeries.setData(predictionData.map(point => ({
      time: point.time,
      value: point.value,
    })))
  }
})
</script>