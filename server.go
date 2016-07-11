package main

import (
	"log"

	"github.com/xiaonanln/yahoo-finance-prefetch/symbolsets"
	"github.com/xiaonanln/yahoo-finance-prefetch/yahoo_finance"
)

func main() {
	log.Printf("Symbol count: %d", len(symbolsets.AllSymbols))
	yahoo_finance.Fetch("NTES", "2016-07-01", "2016-07-12")
}
