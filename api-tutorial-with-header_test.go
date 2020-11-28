package main

import "testing"

func TestGetPriceMultiFullWithHeader(t *testing.T) {
	if err := getPriceMultiFullWithHeader("ETH", "USD", ""); err != nil {
		t.Fatal(err)
	}
}
