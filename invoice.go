package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateInvoice(a fyne.App) {
	clientName := GetClientName()

	input := widget.NewSelect(clientName, func(s string) {

	})
	input.Resize(fyne.NewSize(myWindow.Canvas().Size().Width, 40))
	input.Move(fyne.NewPos(0, 0))

	productName := GetProductName()

	in1 := widget.NewSelect(productName, func(s string) {

	})

	in1.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in1.Move(fyne.NewPos(0, 50))

	in2 := widget.NewEntry()
	in2.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in2.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2.5+5, 50))
	var btn *widget.Button
	btn = widget.NewButton("Add", func() {
		productName := in1.Selected
		qt := in2.Text
		fmt.Println(productName, qt)

		priceMap := GetPrice(productName)
		productPrice := priceMap[0]["price"]
		fmt.Println(productPrice)

		var mainData = [][]string{[]string{"Product", "QT", "Price"},
			[]string{"Mi Note 92 Pro", "12", "12000"}, []string{"Iphone 12 pro max", "1", "12000"}}

		showInvoiceDataOnList(mainData, input, in1, in2, btn)
	})
	btn.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	btn.Move(fyne.NewPos(myWindow.Canvas().Size().Width-myWindow.Canvas().Size().Width/5.5, 50))
	var mainData = [][]string{[]string{},
		[]string{}}
	showInvoiceDataOnList(mainData, input, in1, in2, btn)
	btn.Icon = theme.ContentAddIcon()

	myWindow.Show()
}

func showInvoiceDataOnList(mainData [][]string, input *widget.Select, in1 *widget.Select, in2 *widget.Entry, btn *widget.Button) {

	list := widget.NewTable(
		func() (int, int) {
			return len(mainData), len(mainData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(mainData[i.Row][i.Col])
		},
	)

	list.Resize(fyne.NewSize(myWindow.Canvas().Size().Width, myWindow.Canvas().Size().Height-250))
	list.Move(fyne.NewPos(0, 100))

	list.SetColumnWidth(0, 150.0)
	list.SetColumnWidth(1, 50.0)
	list.SetColumnWidth(3, 210.0)

	// widget.NewEntry().Validate()

	myWindow.SetContent(
		container.NewWithoutLayout(input, in1, in2, btn, list),
	)
}
