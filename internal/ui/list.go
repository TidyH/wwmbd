package ui

import (
	"mburry_stonks/internal/config"
	"mburry_stonks/internal/finance"

	"fyne.io/fyne/v2"
)

func newList() *widget.List {
	data := config.StockData
	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return container.NewPadded(
				widget.NewLabel("          "),
				widget.NewButton("          ", nil),
			)
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*fyne.Container).Objects[0].(*widget.Label).SetText(data[int(i)])
			o.(*fyne.Container).Objects[1].(*widget.Button).SetText(data[int(i)])

			o.(*fyne.Container).Objects[1].(*widget.Button).OnTapped = func() {
				ticker := data[int(i)]
				tickerQuote, err := finance.GetTickerQuote(ticker)
				if err != nil {
					// Handle error
				}

				labelPrice := widget.NewLabel(fmt.Sprintf("Current price is: $%.2f", tickerQuote.RegularMarketPrice))
				w4 := config.NewWindow(ticker)
				w4.Resize(fyne.Size{Width: 400, Height: 200})
				w4.SetContent(labelPrice)
				w4.Show()
			}
		},
	)
	return list
}
