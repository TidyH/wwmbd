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

func (et *EntryTop) Layout(mg *fyne.Container) *fyne.Container {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Stock Ticker")

	entry.OnSubmitted = func(text string) {
		newQuote, err := finance.GetTickerQuote(text)

		if err != nil {
			panic((err))
		}

		// report := canvas.NewText(newQuote.FullExchangeName, color.Black)
		// obj := container.NewWithoutLayout(report)
		// obj := container.New(layout.NewVBoxLayout(), finance.PrettyTickerData(*newQuote))
		obj := finance.PrettyTickerData(*newQuote)

		mg.Add(obj)
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
