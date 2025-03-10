package main

import (
	"fmt"

	"example.com/crypto-masters/api"
)

func main() {
	rate, err := api.GetRate("BTC")

	if err == nil {
		fmt.Printf("The rate for %v is %f", rate.Currency, rate.Price)
	}

}
