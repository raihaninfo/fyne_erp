package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var mainData = [][]string{{"Product Name", "QT", "Price PerUnit", "Total"}}

func CreateInvoice(a fyne.App) {
	clientName := GetClientName()

	input := widget.NewSelect(clientName, func(s string) {

	})
	input.Resize(fyne.NewSize(myWindow.Canvas().Size().Width, 40))
	input.Move(fyne.NewPos(0, 0))

	productName := GetProductName()

	in1 := widget.NewSelect(productName, func(s string) {})

	in1.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in1.Move(fyne.NewPos(0, 50))

	in2 := widget.NewEntry()
	in2.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in2.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2.5+5, 50))
	var btn *widget.Button

	btn = widget.NewButton("Add", func() {
		productName := in1.Selected
		qt := in2.Text

		if len(productName) > 0 && len(qt) > 0 {
			priceMap := GetPrice(productName)
			productPrice := priceMap[0]["price"]
			pdPrice := fmt.Sprintf("%v", productPrice)

			priceInt, _ := strconv.ParseFloat(pdPrice, 64)
			qtInt, _ := strconv.ParseFloat(qt, 64)
			total := priceInt * qtInt

			singleData := []string{productName, qt, pdPrice, fmt.Sprintf("%.0f", total)}
			mainData = append(mainData, singleData)

			showInvoiceDataOnList(mainData, input, in1, in2, btn)
		} else {
			dialog.NewInformation("Warning!", "Please Check Product and quantity", myWindow).Show()
		}
		in1.ClearSelected()
		in2.Text = ""
		in2.Refresh()

	})

	btn.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	btn.Move(fyne.NewPos(myWindow.Canvas().Size().Width-myWindow.Canvas().Size().Width/5.5, 50))

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

	disLabel := widget.NewLabel("Discount = 500")
	disLabel.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	disLabel.Move(fyne.NewPos(0, myWindow.Canvas().Size().Height-130))

	totalLabel := widget.NewLabel("Total = 500")
	totalLabel.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	totalLabel.Move(fyne.NewPos(0, myWindow.Canvas().Size().Height-100))
	ShowInvoice()
	myWindow.SetContent(
		container.NewWithoutLayout(input, in1, in2, btn, list, disLabel, totalLabel),
	)
}
