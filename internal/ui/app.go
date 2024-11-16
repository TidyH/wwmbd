package ui

import (
	"mburry_stonks/internal/config"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
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

	return &App{a: a, w: w}
}

func (a *App) Run() {
	a.w.ShowAndRun()
}
