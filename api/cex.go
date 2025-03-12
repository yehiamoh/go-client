package api

import (
	"fmt"
	"io"
	"net/http"
	"strings"

	data "example.com/crypto-masters/datatypes"
	"example.com/crypto-masters/util"
)

const apiUrl = "https://cex.io/api/ticker/%v/USD"

func GetRate(currency string) (*data.Rate, error) {
	//	strings package have all strings functions needed
	upperCaseCurrency := strings.ToUpper(currency)

	//	http package is builtin package in go that have all http functionallity needed
	//	(S)printf return a string not just printing
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCaseCurrency))
	if err != nil {
		// when changing data.Rate to *data.Rate (adding pointer) that means you are returning pointer to rate which can be null
		return nil, err
	}
	var response CEXResponse
	// we have list of status code constants (eg StatuOk 200)
	if res.StatusCode == http.StatusOK {
		// resbody type isnot a stream or byte its ReadCloser so we have to use io package to read body sync
		bodybytes, err := io.ReadAll(res.Body) // some sort of await
		if err != nil {
			return nil, err
		}
		// Parse the JSON response body into the response struct
		err = util.ParseJson(bodybytes, &response)
		//fmt.Println(response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code recevied : %v", res.StatusCode)
	}

	var rate data.Rate
	rate = *rate.NewRate(currency, response.Bid)
	// we have to use & to get pointer of variable
	//fmt.Println(rate)
	return &rate, nil
}
