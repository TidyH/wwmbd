package ui

import (
	"wwmbd/internal/finance"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type LabelTop struct {
	label *widget.Label
	top   fyne.Canvas
}

func NewLabelTop() *LabelTop {
	l := widget.NewLabel("Search for stoock ticker")
	return &LabelTop{label: l}
}

func (lt *LabelTop) Layout() *fyne.Container {
	c := container.NewVBox()
	c.Add(lt.label)
	return c
}

type EntryTop struct {
	entry *widget.Entry
	top   fyne.Canvas
}

func NewEntryTop() *EntryTop {
	e := widget.NewEntry()
	return &EntryTop{entry: e}
}

func (et *EntryTop) Layout(mg *MiddleGrid, a *App) *fyne.Container {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Stock Ticker")

	entry.OnSubmitted = func(text string) {
		newQuote, err := finance.GetTickerQuote(text)

		if err != nil {
			panic((err))
		}

		// report := canvas.NewText(newQuote.FullExchangeName, color.Black)
		// mg.Add(report)

		tickerShortName := widget.NewLabel(newQuote.ShortName)
		content := container.New(layout.NewVBoxLayout(), tickerShortName)
		newWindow := a.a.NewWindow("hi")
		newWindow.SetContent(content)
		newWindow.Show()
	}

	c := container.NewVBox()
	c.Add(entry)
	return c
}

type Top struct {
	labelTop *LabelTop
	entryTop *EntryTop
	top      fyne.Canvas
}

func NewTop() *Top {
	top := &Top{
		labelTop: NewLabelTop(),
		entryTop: NewEntryTop(),
	}
	return top
}
