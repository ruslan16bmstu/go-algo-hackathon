import { StockWithPrice } from '../../api'
import { Point } from '../../utils/point'

export type Prediction = {
  delta: number
  points: Point[]
}

export type StockWithPrediction = StockWithPrice & Prediction


export interface PredictionData {
  time: string
  value: number
}