package ui

import (
	"wwmbd/internal/finance"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type LabelTop struct {
	label *widget.Label
}

func NewLabelTop() *LabelTop {
	l := widget.NewLabel("Search for stock ticker")
	return &LabelTop{label: l}
}

func (lt *LabelTop) Layout() *fyne.Container {
	c := container.NewVBox()
	c.Add(lt.label)
	return c
}

type EntryTop struct {
	entry *widget.Entry
}

func NewEntryTop() *EntryTop {
	e := widget.NewEntry()
	return &EntryTop{entry: e}
}

func (et *EntryTop) Layout(mg *fyne.Container, a *App) *fyne.Container {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Stock Ticker")

	entry.OnSubmitted = func(text string) {
		newQuote, err := finance.GetTickerQuote(text)

		if err != nil {
			panic((err))
		}

		obj := finance.PrettyTickerData(*newQuote)
		mg.Add(obj)

		br := finance.BurryReport{StockData: *newQuote, A: a.a}
		br.CreateReport() // TODO: This needs to be hidden, need to find expand report with more details
	}

	c := container.NewVBox()
	c.Add(entry)
	return c
}

type Top struct {
	labelTop *LabelTop
	entryTop *EntryTop
}

func NewTop() *Top {
	top := &Top{
		labelTop: NewLabelTop(),
		entryTop: NewEntryTop(),
	}
	return top
}
