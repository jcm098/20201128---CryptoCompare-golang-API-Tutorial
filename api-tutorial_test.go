package main

import "testing"

func TestGetPriceMultiFull(t *testing.T) {
	if err := getPriceMultiFull("ETH", "USD", ""); err != nil {
		t.Fatal(err)
	}
}
