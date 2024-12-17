package main

import "fmt"

func main() {

	stocks := map[string]float64{
		"AMZN": 1699.8,
		"GOOG": 1129.19,
		"MSFT": 98.61,
	}

	fmt.Println(len(stocks))

	// Get a value
	fmt.Println(stocks["MSFT"])

	// if the key doesn't exist, it prints 0
	fmt.Println(stocks["TSLA"]) // 0

	// what if i want to check if it is found?
	teslaStock, ok := stocks["TSLA"]
	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(teslaStock)
	}

	// SET
	stocks["TSLA"] = 322.12
	fmt.Println(stocks)

	// REMOVE
	delete(stocks, "AMZN")
	fmt.Println(stocks)

}
