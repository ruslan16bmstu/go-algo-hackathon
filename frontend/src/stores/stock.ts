import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getAllStocksPrediction, getStockPredictionBySecId, StockPrediction } from '../api/backend'
import { getStockInfoAndCandles, StockWithCandles } from '../api/moex'
import { formatDate } from '../utils/date'
import { useLoading } from '../utils/loading'

export type StockWithIndustryWithPredictionWithCandles = StockWithCandles & StockPrediction

export const candleIntervalDays = 40

export const useStockStore = defineStore('stock', () => {
  const data = ref<StockWithIndustryWithPredictionWithCandles[]>([])
  const {isLoading, load} = useLoading()
  
  const loadData = async () => await load(async () => {
    const stocksPrediction = await getAllStocksPrediction()
    if (stocksPrediction != null) {
      const stocksWithCandles = await Promise.allSettled(stocksPrediction.map(({secId}) => {
        const d = new Date()
        const to = formatDate(d)
        d.setDate(d.getDate() - candleIntervalDays)
        const from = formatDate(d)
        return getStockInfoAndCandles(secId, from, to)
      }))
      
      data.value = stocksPrediction.map((s, i) => {
        const stockWithCandles = ((stocksWithCandles[i] as any).value as StockWithCandles)
        
        return {
          secId: s.secId,
          industry: s.industry,
          stockPrice: s.stockPrice, // прогноз
          candles: stockWithCandles.candles,
          name: stockWithCandles.name,
          price: stockWithCandles.price, // текущая
          delta: (s.stockPrice - stockWithCandles.price) / s.stockPrice * 100,
        }
      })
    }
  })
  
  const loadOne = async (secId: string) => await load(async () => {
    if (data.value.find(s => s.secId == secId)) {
      return
    }
    
    const s = await getStockPredictionBySecId(secId)
    if (s != null) {
      const d = new Date()
      const to = formatDate(d)
      d.setDate(d.getDate() - candleIntervalDays)
      const from = formatDate(d)
      const stockWithCandles = await getStockInfoAndCandles(secId, from, to)
      
      data.value = data.value.concat({
        secId: s.secId,
        industry: s.industry,
        stockPrice: s.stockPrice, // прогноз
        candles: stockWithCandles.candles,
        name: stockWithCandles.name,
        price: stockWithCandles.price, // текущая
        delta: (s.stockPrice - stockWithCandles.price) / s.stockPrice * 100,
      })
    }
  })
  
  const getStockBySecId = (secId: string) => {
    return data.value.find(d => d.secId === secId)
  }
  
  return {data, isLoading, loadData, getStockBySecId, loadOne}
})