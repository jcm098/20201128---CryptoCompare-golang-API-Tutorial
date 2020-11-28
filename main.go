package main

import "log"

var apiEndpoint = "https://min-api.cryptocompare.com/"

func main() {

	fsyms := "ETH"
	tsyms := "USD"
	var apiKey string

	if err := getPriceMultiFull(fsyms, tsyms, apiKey); err != nil {
		log.Fatal(err)
	}

	//if err := getPriceMultiFullWithHeader(fsyms, tsyms, apiKey); err != nil {
	//	log.Fatal(err)
	//}
}