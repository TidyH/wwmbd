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

type BurryReport struct {
	StockData finance.Quote
	A         fyne.App
}

func NewReport(a fyne.App, stockData *finance.Quote) fyne.Window {
	w := a.NewWindow("The Burry Report")
	return w
}

func (r *BurryReport) CreateReport() {
	w := r.A.NewWindow("The Burry Report")
	l := widget.NewLabel(r.StockData.ShortName)

	graph := r.createGraph(datetime.Datetime{Month: 1, Day: 1, Year: 2017}, datetime.Datetime{Month: 1, Day: 1, Year: 2018}, datetime.OneDay) // TODO: need to allow user query or trailing 52 week period
	graph.SetMinSize(fyne.NewSize(600, 400))
	content := container.NewVBox()
	content.Add(l)
	content.Add(graph)

	w.SetContent(content)
	w.Show()
}

func (r *BurryReport) createGraph(start datetime.Datetime, end datetime.Datetime, interval datetime.Interval) *canvas.Image {
	var prices []decimal.Decimal
	var dates []time.Time
	// var dates []float64

	chartConfig := chart.Params{
		Symbol:   r.StockData.Symbol,
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
	graphConfig.Title = r.StockData.ShortName

	img, err := writer.Image()
	if err != nil {
		fmt.Println(err)
	}

	return canvas.NewImageFromImage(img)
}
