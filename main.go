package main

import (
	"fmt"
	"sync"

	"example.com/crypto-masters/api"
)

func main() {

	var wg sync.WaitGroup
	currencies := []string{"BTC", "ETH", "BCH"}

	for _, currency := range currencies {
		wg.Add(1)
		go getCurrencyData(currency, &wg)
	}
	wg.Wait()
}

func getCurrencyData(currency string, wg *sync.WaitGroup) {
	rate, err := api.GetRate(currency)

	if err != nil {
		fmt.Println("Error in get currency function : ", err)
	}
	fmt.Printf("The rate for %v is %.2f \n", rate.Currency, rate.Price)

	wg.Done()
}
