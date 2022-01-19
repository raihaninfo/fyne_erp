package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowDashbord(a fyne.App) {
	win := myWindow

	headLable := widget.NewLabel("Welcome to MiniERP")

	btn1 := widget.NewButton("Add New Client", func() {
		ShowClient(myApp)
	})
	btn2 := widget.NewButton("All Client", func() {
		ShowData(myApp)
	})

	win.SetContent(
		container.NewVBox(
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(1,
					headLable,
				),
				container.NewGridWithColumns(2,
					btn1,
					btn2,
				),
			),
		),
	)

	win.Show()
}
