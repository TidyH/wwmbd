package ui

import (
	"wwmbd/internal/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func newList() *widget.List {
	data := config.StockData

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			longestItem := data[0]
			for i := 1; i < len(data); i++ {
				if len(data[i]) > len(longestItem) {
					longestItem = data[i]
				}
			}

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
