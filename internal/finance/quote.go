package finance

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/quote"
)

func GetTickerQuote(ticker string) (*finance.Quote, error) {
	ticker = strings.TrimSpace(strings.ToUpper(ticker))
	return quote.Get(ticker)
}

// Best way to display information from the Quote
func PrettyTickerData(tickerQuote finance.Quote) *widget.Table {
	// thinking we create a list
	// This is surface level data, what do we care about?
	data := [][]string{
		{"Symbol", tickerQuote.Symbol},
		{"Shortname", tickerQuote.ShortName},
		{"RegularMarketPrice", fmt.Sprintf("%.2f", tickerQuote.RegularMarketPrice)},
		{"FiftyTwoWeekLow", fmt.Sprintf("%.2f", tickerQuote.FiftyTwoWeekLow)},
		{"FiftyTwoWeekHigh", fmt.Sprintf("%.2f", tickerQuote.FiftyTwoWeekHigh)},
		{"AverageDailyVolume3Month", fmt.Sprintf("%d", tickerQuote.AverageDailyVolume3Month)},
	}

	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Stock Snapshot")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	return list
}
