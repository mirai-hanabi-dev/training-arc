/**
 * @param {number[]} prices
 * @return {number}
 */
const maxProfit = prices => {
  prices.unshift(Number.MAX_SAFE_INTEGER)
  prices.push(Number.MIN_SAFE_INTEGER)

  let currentBuy = null
  let result = 0
  for (let i = 1; i < prices.length - 1; i++) {
    // If lowest
    if (prices[i] < prices[i - 1] && prices[i] <= prices[i + 1]) {
      currentBuy = prices[i]
      continue
    }
    // If highest
    if (prices[i] >= prices[i - 1] && prices[i] > prices[i + 1]) {
      result += prices[i] - currentBuy
      currentBuy = null
    }
  }

  return result
}

const prices = [7, 6, 4, 3, 1]
console.log(maxProfit(prices))
