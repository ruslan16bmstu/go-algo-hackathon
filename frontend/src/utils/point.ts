// Определение типов для точек
export interface Point {
  x: number
  y: number
}

export const convertToCanvasCoords = (point: Point, height: number): Point => {
  return {
    x: point.x, // x остаётся без изменений
    y: height - point.y, // Вычитаем y из общей высоты, чтобы инвертировать его положение
  }
}

export interface MinMax {
  minX: number,
  minY: number,
  maxX: number,
  maxY: number
}

export function findMinMax(points: Point[]) {
  const xArr: number[] = []
  const yArr: number[] = []
  
  points.forEach(({x, y}) => {
    xArr.push(x)
    yArr.push(y)
  })
  
  const maxX = Math.max(...xArr)
  const maxY = Math.max(...yArr)
  const minX = Math.min(...xArr)
  const minY = Math.min(...yArr)
  
  return {maxX, maxY, minX, minY}
}

export const scalePointsToCanvas = (points: Point[], width: number, height: number, minMax: MinMax): Point[] => {
  const {minX, minY, maxX, maxY} = minMax
  
  const scaleX = width / (maxX - minX || 1) // Учесть случай деления на 0
  const scaleY = height / (maxY - minY || 1)
  
  return points.map(point => {
    const p = {
      x: (point.x - minX) * scaleX,
      y: (point.y - minY) * scaleY,
    }
    return convertToCanvasCoords(p, height)
  })
}
