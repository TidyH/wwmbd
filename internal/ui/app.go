package ui

import (
	"wwmbd/internal/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

type App struct {
	a fyne.App
	w fyne.Window
}

func NewApp() *App {
	a := app.New()
	w := a.NewWindow(config.Title)
	w.Resize(fyne.Size{Width: 1024, Height: 800})

	// Set up the UI components and layout here
	middleGrid := container.New(layout.NewGridWrapLayout(fyne.NewSize(450, 250))) // size is found by feel, can i find this automatically?
	middleGridScroll := container.NewVScroll(middleGrid)

	top := NewTop()
	topContent := container.NewVBox(top.labelTop.Layout(), top.entryTop.Layout(middleGrid))

	left := newList()

	mainContent := container.NewBorder(topContent, nil, left, nil, middleGridScroll)

	w.SetContent(mainContent)

	return &App{a: a, w: w}
}

func (a *App) Run() {
	a.w.ShowAndRun()
}
