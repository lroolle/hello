package hello

import (
	"github.com/lroolle/hello/utils"
	"rsc.io/quote"
	quoteV3 "rsc.io/quote/v3"
)

func Hello() string {
	utils.PrintHello()
	return quote.Hello()
}

func Proverb() string {
	return quoteV3.Concurrency()
}
