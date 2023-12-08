import axios, { AxiosResponse } from 'axios'

export interface Stock {
  name: string;
  code: string;
  secId: string;
  // Другие свойства акции, которые вам могут быть интересны
}

export async function getAllStocks(): Promise<Stock[]> {
  try {
    const allStocks: Stock[] = []
    let start = 0
    let hasMoreStocks = true
    
    while (hasMoreStocks) {
      const response: AxiosResponse = await axios.get('https://iss.moex.com/iss/securities.json', {
        params: {
          engine: 'stock',
          market: 'shares',
          'iss.only': 'securities',
          'is_trading': true,
          'group_by': 'type',
          'group_by_filter': 'common_share',
          start,
          limit: 100,
        },
      })
      
      const data = response.data
      const stocksData = data.securities.data as any[]
      
      if (stocksData.length > 0) {
        const stocks: Stock[] = stocksData.map((stockData: any) => ({
          code: stockData[0], // Код акции
          secId: stockData[1], // secid
          name: stockData[2] // Название акции
          // Другие свойства акции, которые вам могут быть интересны
        }))
        
        allStocks.push(...stocks)
        start += 100 // Увеличение параметра start для следующего запроса
      } else {
        hasMoreStocks = false // Если больше акций нет, выход из цикла
      }
    }
    
    return allStocks
  } catch (error) {
    console.error('Произошла ошибка:', error)
    return []
  }
}
