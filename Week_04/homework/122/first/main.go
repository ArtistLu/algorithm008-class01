package main

func main() {

}

func maxProfit(prices []int) int {
	re := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			re += prices[i] - prices[i-1]
		}
	}

	return re
}
