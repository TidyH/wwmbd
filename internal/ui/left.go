package ui

import (
	"wwmbd/helpers"
	"wwmbd/internal/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func newList() *widget.List {
	data := config.StockData
	longestItem := helpers.LongestString(data)

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewButton(longestItem, nil) // i don't know how else to fit the text...
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Button).SetText(data[int(i)])
		},
	)
	return list
}

// TODO: Not really sure what to use the left bar for
//	currently just leaving some placeholder data
