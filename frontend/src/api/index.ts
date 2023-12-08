import axios, { AxiosResponse } from 'axios'

export type Stock = {
  name: string;
  code: string;
  secId: string;
}

export type StockWithPrice = Stock & {
  price: number
}

export async function getAllStocks(): Promise<StockWithPrice[]> {
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
          name: stockData[4] // Название акции
          // Другие свойства акции, которые вам могут быть интересны
        }))
        
        allStocks.push(...stocks)
        start += 100 // Увеличение параметра start для следующего запроса
      } else {
        hasMoreStocks = false // Если больше акций нет, выход из цикла
      }
    }
    
    const stockWithPriceArr: StockWithPrice[] = []
    
    const reqs: Promise<any>[] = []
    for (const s of allStocks) {
      reqs.push(getStockPrice(s.secId))
    }
    const res = await Promise.allSettled(reqs)
    
    for (let i = 0; i < allStocks.length; i++) {
      const price = (res[i] as any).value
      if (price != null) {
        stockWithPriceArr.push({...allStocks[i], price})
      }
    }
    
    return stockWithPriceArr
  } catch (error) {
    console.error('Произошла ошибка:', error)
    return []
  }
}

interface MarketData {
  data: any[]; // Здесь можно использовать тип данных, соответствующий структуре ответа API
}

async function getStockPrice(symbol: string): Promise<number | null> {
  const API_URL = `https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities/${symbol}.json`
  
  try {
    const response = await axios.get(API_URL)
    
    // Проверяем успешность запроса и наличие данных
    if (response.status === 200 && response.data.marketdata.data.length) {
      const price = response.data['marketdata'].data[0][12] // Индекс цены в полученных данных
      return price ? Number(price) : null
    } else {
      return null
    }
  } catch (error) {
    console.error('Error fetching data:', error)
    return null
  }
}
