package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowAllClient(a fyne.App) {
	win := myWindow
	btnHead:= widget.NewButton("Dashbord", func() {
		ShowDashbord(myApp)
	})

	label := widget.NewLabel("All Client")

	win.SetContent(
		container.NewVBox(btnHead, label),
	)

	win.Show()
}
