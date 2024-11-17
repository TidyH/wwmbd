package finance

import (
	"strings"

	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
)

func GetTickerQuote(ticker string) (*finance.Quote, error) {
	ticker = strings.TrimSpace(strings.ToUpper(ticker))
	return quote.Get(ticker)
}
