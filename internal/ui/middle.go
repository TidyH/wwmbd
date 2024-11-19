package ui

import (
	"fyne.io/fyne/v2"
)

type MiddleGrid struct {
	*fyne.Container
}

func (mg *MiddleGrid) Add(canvas fyne.CanvasObject) *MiddleGrid {
	res := mg.Add(canvas)
	if res == nil {
		return nil
	}
	return res
}
