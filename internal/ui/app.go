package ui

import (
	"wwmbd/internal/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	top := NewTop()
	topContent := container.NewVBox(top.labelTop.Layout(), top.entryTop.Layout())

	mainContent := container.NewBorder(topContent, nil, nil, nil, nil)

	w.SetContent(mainContent)

	return &App{a: a, w: w}
}

func (a *App) Run() {
	a.w.ShowAndRun()
}
