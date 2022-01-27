package main

import (
	"fmt"
	"image/color"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowDashbord(a fyne.App) {
	win := myWindow
	comName := "Mini ERP"
	welcome := fmt.Sprintf("Welcome to %s", comName)
	headLable := widget.NewCard(
		welcome,
		"2no gate sheikh farid market, ground floor, Nasirabad Chittagong",
		widget.NewLabel("email@gmail.com, 01852-52525"),
	)

	// headLable.TextSize = 40

	btn1 := widget.NewButton("Add New Client", func() {
		ShowClient(myApp)
	})
	btn2 := widget.NewButton("All Client", func() {
		ShowData(myApp)
	})
	btn3 := widget.NewButton("Add Product", func() {
		ShowProductAdd(myApp)
	})
	btn4 := widget.NewButton("Add Product Group", func() {
		ShowProductAdd(myApp)
	})

	totalClient := processAllClientData()
	totalClientCount := strconv.Itoa((len(totalClient) - 1))

	card1 := widget.NewCard(
		"Total Client",
		totalClientCount,
		canvas.NewRectangle(color.Black),
	)

	card2 := widget.NewCard(
		"Total Product",
		"12",
		canvas.NewRectangle(color.White),
	)

	win.SetContent(
		container.NewVBox(
			container.NewGridWithColumns(1,
				container.NewGridWithColumns(1,
					container.NewCenter(headLable),
				),
				container.NewGridWithColumns(3,
					container.NewVBox(btn1, btn2, btn3, btn4),
					container.NewVBox(card1),
					container.NewVBox(card2),
				), container.NewGridWithColumns(3), // container.NewVBox(),
				// container.NewVBox(),
				// container.NewVBox(),

			),
		),
	)

	win.Show()
}
