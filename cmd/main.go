package main

import (
	"wwmbd/internal/ui"
)

func main() {
	println("Starting App")
	app := ui.NewApp()
	app.Run()
}
