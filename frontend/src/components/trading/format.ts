export const formatStockName = (name: string) => {
  name = name.replace(/"/g,'').replace(/\(/g,'').replace(/\)/g, '')
    .replace('ПАО', '').replace('ао', '').replace(' -', '')
  
  const maxLen = 13
  if (name.length > maxLen) {
    name = name.slice(0, maxLen - 2) + '...'
  }
  
  return name
}

export const formatPrice = (price: number, precision?: number) => price.toPrecision(precision).replace('.', ',')
export const formatDelta = (delta: number) => (delta > 0 ? '+' : '') + formatPrice(delta, 3) + '%'
