package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowDashbord(a fyne.App) {
	win := myWindow

	headLable := canvas.NewText("Welcome to MiniERP", color.RGBA{R: 13, G: 13, B: 235, A: 100})
	headLable.TextSize = 40

	btn1 := widget.NewButton("Add New Client", func() {
		ShowClient(myApp)
	})
	btn2 := widget.NewButton("All Client", func() {
		ShowData(myApp)
	})

	card1 := widget.NewCard(
		"Total Cleint 33",
		"",
		canvas.NewLine(color.Opaque),
	)

	card2 := widget.NewCard(
		"Total Product 54",
		"",
		canvas.NewLine(color.Opaque),
	)

	card3 := widget.NewCard(
		"Total Cleint 35",
		"",
		canvas.NewLine(color.Opaque),
	)

	win.SetContent(
		container.NewVBox(
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(1,
					container.NewCenter(headLable),
				),
				container.NewGridWithColumns(3,
					container.NewVBox(btn1, btn2),
					container.NewVBox(card1),

					container.NewVBox(card3),
				), container.NewGridWithColumns(3,
					container.NewVBox(),
					container.NewVBox(card1),
					container.NewVBox(card2),
				),
			),
		),
	)

	win.Show()
}
