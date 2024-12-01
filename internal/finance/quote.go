package finance

import (
	"fmt"
	"strings"
	"wwmbd/helpers"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/equity"
	"github.com/piquette/finance-go/quote"
)

type StockFinancials struct {
	QuoteData     *finance.Quote
	FinancialData *finance.Equity
}

func (s *StockFinancials) SetQuotes(ticker string) {
	ticker = strings.TrimSpace(strings.ToUpper(ticker))
	s.QuoteData, _ = quote.Get(ticker)
}

func (s *StockFinancials) SetEquities(ticker string) {
	ticker = strings.TrimSpace(strings.ToUpper(ticker))
	s.FinancialData, _ = equity.Get(ticker)
}

// Search for best way to display information from the Quote
func PrettyTickerData(tickerQuote finance.Quote) *widget.Table {
	data := [][]string{
		{"Symbol", tickerQuote.Symbol},
		{"Shortname", tickerQuote.ShortName},
		{"RegularMarketPrice", fmt.Sprintf("%.2f", tickerQuote.RegularMarketPrice)},
		{"FiftyTwoWeekLow", fmt.Sprintf("%.2f", tickerQuote.FiftyTwoWeekLow)},
		{"FiftyTwoWeekHigh", fmt.Sprintf("%.2f", tickerQuote.FiftyTwoWeekHigh)},
		{"AverageDailyVolume3Month", fmt.Sprintf("%d", tickerQuote.AverageDailyVolume3Month)},
	}

	columnOneLength := helpers.LongestStringInColumn(data, 0)

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {

			return widget.NewLabel(columnOneLength)
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return list
}
