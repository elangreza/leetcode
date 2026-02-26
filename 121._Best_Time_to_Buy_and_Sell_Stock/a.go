package a

func maxProfit(prices []int) int {
	// track buy with min value
	// iterate sell over the prices
	buy := 0
	// calculate res with max
	res := 0

	for idx, price := range prices {

		if price < prices[buy] {
			buy = idx
			// sell = idx
		}

		// if price > prices[sell] {
		// 	sell = idx
		// }

		res = max(res, price-prices[buy])
	}

	return res
}
