import { AxiosResponse } from 'axios'
import { apiClient } from './base'

export type StockPrediction = {
  industry: number
  secId: string
  stockPrice: number
  delta: number
}

export async function getAllStocksPrediction(): Promise<StockPrediction[]> {
  const API_URL = '/rating' // Поменяйте на ваше реальное API URL
  const limit = 100 // Задайте максимальный лимит, который можно запросить за один раз
  let skip = 0
  let allStocks: StockPrediction[] = []
  
  try {
    let response: AxiosResponse<{ stocks: StockPrediction[] }> | null = null
    
    do {
      response = await apiClient.get(API_URL, {
        params: {
          skip: skip,
          limit: limit,
        },
      }) as AxiosResponse<{ stocks: StockPrediction[] }>
      
      if (response.status === 200) {
        const currentStocks = response.data.stocks
        allStocks = allStocks.concat(currentStocks)
        skip += limit
      } else {
        throw new Error('Failed to fetch data')
      }
    } while (response.data.stocks.length > 0)
    
    return allStocks
  } catch (error) {
    console.error('Error fetching data:', error)
    return []
  }
}

export async function getStockPredictionBySecId(secId: string): Promise<StockPrediction | null> {
  const API_URL = `/stock/${secId}` // Поменяйте на ваше реальное API URL
  try {
    
    const response = await apiClient.get(API_URL, {})
    
    if (response.status === 200) {
      return response.data
    } else {
      throw new Error('Failed to fetch data')
    }
  } catch (error) {
    console.error('Error fetching data:', error)
    return null
  }
}
