// test
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Test Window")
	w.Resize(fyne.NewSize(600, 600))

	btn := widget.NewButton("test", func() {

	})

	w.SetContent(btn)

	w.ShowAndRun()
}
