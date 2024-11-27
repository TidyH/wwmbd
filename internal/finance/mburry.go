package finance

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/piquette/finance-go"
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
	p := widget.NewLabel(r.StockData.ShortName)

	w.SetContent(p)
	w.Show()
}
