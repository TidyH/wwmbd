package finance

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/piquette/finance-go"
	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
	"github.com/shopspring/decimal"
	myChart "github.com/wcharczuk/go-chart/v2"
)

// S stands for Stonks ;)
type BurryReport struct {
	S *StockFinancials
	A fyne.App
}

func NewReport(a fyne.App, stockData *finance.Quote) fyne.Window {
	w := a.NewWindow("The Burry Report")
	return w
}

func (r *BurryReport) CreateReport() {
	content := container.NewVBox()
	w := r.A.NewWindow("The Burry Report")
	l := widget.NewLabel(r.S.QuoteData.ShortName)

	// Opening Graph, trailing 52 week performance
	timeNow := datetime.Datetime{Month: int(time.Now().Month()), Day: time.Now().Day(), Year: time.Now().Year()}
	timeLastYear := datetime.Datetime{Month: int(time.Now().Month()), Day: time.Now().Day(), Year: time.Now().AddDate(-1, 0, 0).Year()}
	graph := r.createGraph(timeLastYear, timeNow, datetime.OneDay)
	graph.SetMinSize(fyne.NewSize(600, 400))
	content.Add(l)
	content.Add(graph)

	// EVEBITDA := r.calcEVEBITDA()

	w.SetContent(content)
	w.Show()
}

func (r *BurryReport) createGraph(start datetime.Datetime, end datetime.Datetime, interval datetime.Interval) *canvas.Image {
	var prices []decimal.Decimal
	var dates []time.Time

	chartConfig := chart.Params{
		Symbol:   r.S.QuoteData.Symbol,
		Start:    &start,
		End:      &end,
		Interval: interval,
	}

	iter := chart.Get(&chartConfig)

	for iter.Next() {
		b := iter.Bar()
		unixtimeUTC := time.Unix(int64(b.Timestamp), 0)

		prices = append(prices, b.Close)
		dates = append(dates, unixtimeUTC)
	}

	pricesFloat := make([]float64, len(dates))
	for i, p := range prices {
		pricesFloat[i], _ = decimal.Decimal.Float64(p)
	}

	graphConfig := myChart.Chart{
		Series: []myChart.Series{
			myChart.TimeSeries{
				Style: myChart.Style{
					StrokeColor: myChart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   myChart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: dates,
				YValues: pricesFloat,
			},
		},
	}

	writer := &myChart.ImageWriter{}
	graphConfig.Render(myChart.PNG, writer)
	graphConfig.Title = r.S.QuoteData.ShortName

	img, err := writer.Image()
	if err != nil {
		fmt.Println(err)
	}

	return canvas.NewImageFromImage(img)
}

// EV = market_cap + total_debt - cash
// func (r *BurryReport) calcEVEBITDA() *string {
// 	marketCap := r.S.FinancialData.MarketCap
// }
