package config

import (
	"github.com/piquette/finance-go/quote"
)

const Title = "wwmbd"

var StockData = []string{"red", "green", "blue", "yellow", "stonks"}

func NewWindow(title string) fyne.Window {
	a := app.New()
	w := a.NewWindow(title)
	w.Resize(fyne.Size{Width: 1024, Height: 800})
	return w
}
