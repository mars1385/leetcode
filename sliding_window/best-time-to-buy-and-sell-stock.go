package slidingwindow

func MaxProfit(prices []int) int {

	low := 0
	high := 0
	lasHighPrice := []int{}

	for index, price := range prices {
		if price < prices[low] && len(prices)-1 > index {
			lasHighPrice = append(lasHighPrice, prices[high]-prices[low])
			low = index
			high = index

		}

		if price > prices[high] {
			high = index
		}
	}
	lasHighPrice = append(lasHighPrice, prices[high]-prices[low])
	price := 0
	for _, val := range lasHighPrice {
		if val > price {
			price = val
		}
	}

	return price
}
