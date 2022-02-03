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

func ShowDashbod(a fyne.App) {
	mainMenu()
	win := myWindow
	comInfo, err := GetCompanyInfo("1")
	if err != nil {
		fmt.Println(err)
	}
	comName := comInfo["company_name"]
	comAddress := comInfo["address"]
	comWeb := comInfo["website"]
	comEmail := comInfo["email"]
	comMobile := comInfo["mobile"]

	welcome := fmt.Sprintf("Welcome to %s", comName)
	onlineInfo := fmt.Sprintf("%s, %s, %s", comWeb, comEmail, comMobile)
	headLable := widget.NewCard(
		welcome,
		dataConveter(comAddress),
		widget.NewLabel(onlineInfo),
	)
	headLable.Refresh()

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
		ShowAddGroupItem(myApp)
	})
	btn5 := widget.NewButton("Show All Product", func() {
		ShowAllProduct(myApp)
	})
	btn6 := widget.NewButton("Invoice", func() {
		InvoiceDash(myApp)
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
				container.NewCenter(headLable),
			),
			container.NewGridWithColumns(3,
				container.NewVBox(btn1, btn2, btn3, btn4, btn5, btn6),
				container.NewVBox(card1),
				container.NewVBox(card2),
			),
		),
	)

	win.Show()
}
