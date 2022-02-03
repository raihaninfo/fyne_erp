package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func InvoiceDash(a fyne.App) {
	win := myWindow

	btn1 := widget.NewButton("Back", func() {
		ShowDashbod(myApp)
	})
	btn2 := widget.NewButton("Create Invoice", func() {

	})

	win.SetContent(
		container.NewVBox(btn1, btn2),
	)
	win.Show()
}
