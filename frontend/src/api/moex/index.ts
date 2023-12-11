import axios, { AxiosResponse } from 'axios'

export type Stock = {
  name: string;
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
  const price = (await getStockInfo(symbol))?.price
  return price ? Number(price) : null
}

async function getStockInfo(symbol: string): Promise<StockWithPrice | null> {
  const API_URL = `https://iss.moex.com/iss/engines/stock/markets/shares/boards/TQBR/securities/${symbol}.json`
  
  try {
    const response = await axios.get(API_URL)
    
    // Проверяем успешность запроса и наличие данных
    if (response.status === 200 && response.data.marketdata.data.length) {
      const stock = response.data['securities'].data[0]
      const secId = stock[0]
      const price = stock[3] // Индекс цены в полученных данных
      const name = stock[9]
      
      return {price, name, secId}
    } else {
      return null
    }
  } catch (error) {
    console.error('Error fetching data:', error)
    return null
  }
}

export interface CandleData {
  open: number; // Цена открытия
  close: number; // Цена закрытия
  high: number; // Максимальная цена
  low: number; // Минимальная цена
  time: string; // Временная метка (время свечи)
  volume: number; // Объем свечи
}

export async function getStockCandles(symbol: string, from: string, to: string): Promise<CandleData[]> {
  const API_URL = `https://iss.moex.com/iss/history/engines/stock/markets/shares/boards/TQBR/securities/${symbol}.json`
  
  try {
    const response = await axios.get(API_URL, {
      params: {
        from: from,
        to: to,
        interval: '1',
      },
    })
    
    if (response.status === 200) {
      return response.data.history.data.map((candle: any[]) => ({
        open: candle[6],
        close: candle[11],
        high: candle[8],
        low: candle[7],
        time: candle[1],
        volume: candle[12],
      })).filter((c: CandleData) => c.low != null && c.open != null && c.high != null && c.close != null)
    } else {
      throw new Error('Failed to fetch data')
    }
  } catch (error) {
    console.error('Error fetching data:', error)
    return []
  }
}


export type StockWithCandles = StockWithPrice & {
  candles: CandleData[]
}

export async function getStockInfoAndCandles(symbol: string, candlesFrom: string, candlesTo: string) {
  const [stockInfo, candles] = await Promise.allSettled([getStockInfo(symbol), getStockCandles(symbol, candlesFrom, candlesTo)])
  
  return {...(stockInfo as any).value, candles: (candles as any).value}
}