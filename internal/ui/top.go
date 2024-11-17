package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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

func (et *EntryTop) Layout() *fyne.Container {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("Enter Stock Ticker")
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
