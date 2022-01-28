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

func processCompanyInfo() []string {
	tableData := []string{}
	rows := GetCompanyInfo()
	for i := 0; i < len(rows); i++ {
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["company_name"]))
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["address"]))
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["website"]))
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["email"]))
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["mobile"]))
	}
	return tableData
}

func ShowDashbord(a fyne.App) {
	win := myWindow
	comInfo := processCompanyInfo()
	comName := comInfo[0]
	comAddress := comInfo[1]
	comWeb := comInfo[2]
	comEmail := comInfo[3]
	comMobile := comInfo[4]

	welcome := fmt.Sprintf("Welcome to %s", comName)
	onlineInfo := fmt.Sprintf("%s, %s, %s", comWeb, comEmail, comMobile)
	headLable := widget.NewCard(
		welcome,
		comAddress,
		widget.NewLabel(onlineInfo),
	)

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
