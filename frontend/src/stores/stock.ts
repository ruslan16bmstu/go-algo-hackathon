import { defineStore } from 'pinia'
import { ref } from 'vue'
import { getAllStocks, Stock } from '../api'
import { useLoading } from '../utils/loading'

export const useStockStore = defineStore('stock', () => {
  const data = ref<Stock[]>([])
  const {isLoading, load} = useLoading()
  
  const loadData = async () => {
    const resData = await load(() => getAllStocks())
    if (resData != null) {
      data.value = resData
    }
  }
  
  const getStockBySecId = (secId: string) => {
    return data.value.find(d => d.secId === secId)
  }
  
  return {data, isLoading, loadData, getStockBySecId}
})