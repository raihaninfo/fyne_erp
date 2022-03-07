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

var totalLabel *widget.Label
var formDiscount float64

// var clientFullName string
var mainData = [][]string{{"Product Name", "QT", "Price PerUnit", "Total"}}

var input *widget.Select

func CreateInvoice(a fyne.App) {
	clientName := GetClientName()

	input = widget.NewSelect(clientName, func(s string) {

	})
	input.Resize(fyne.NewSize(myWindow.Canvas().Size().Width, 40))
	input.Move(fyne.NewPos(0, 0))

	productName := GetProductName()

	in1 := widget.NewSelect(productName, func(s string) {})

	in1.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in1.Move(fyne.NewPos(0, 50))

	in2 := widget.NewEntry()
	in2.PlaceHolder = "Quality must be number"
	in2.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in2.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2.5+5, 50))
	var btn *widget.Button

	dis := widget.NewEntry()
	dis.PlaceHolder = "Discount Amount"
	dis.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/4, 40))
	dis.Move(fyne.NewPos(0, myWindow.Canvas().Size().Height-120))

	btn = widget.NewButton("Add", func() {
		productName := in1.Selected
		qt := in2.Text

		dis.OnChanged = func(s string) {
			am, _ := strconv.ParseFloat(s, 32)
			formDiscount = am
			showInvoiceDataOnList(mainData, input, in1, in2, btn, dis)
		}
		if len(productName) > 0 && len(qt) > 0 {
			priceMap := GetPrice(productName)
			productPrice := priceMap[0]["price"]
			pdPrice := fmt.Sprintf("%v", productPrice)

			priceInt, _ := strconv.ParseFloat(pdPrice, 64)
			qtInt, _ := strconv.ParseFloat(qt, 64)
			total := priceInt * qtInt

			singleData := []string{productName, qt, pdPrice, fmt.Sprintf("%.0f", total)}
			mainData = append(mainData, singleData)

			showInvoiceDataOnList(mainData, input, in1, in2, btn, dis)
		} else {
			dialog.NewInformation("Warning!", "Please Check Product and quantity", myWindow).Show()
		}
		in1.ClearSelected()
		in2.Text = ""
		in2.Refresh()

	})

	btn.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	btn.Move(fyne.NewPos(myWindow.Canvas().Size().Width-myWindow.Canvas().Size().Width/5.5, 50))

	showInvoiceDataOnList(mainData, input, in1, in2, btn, dis)
	btn.Icon = theme.ContentAddIcon()

	myWindow.Show()
}

func showInvoiceDataOnList(mainData [][]string, input *widget.Select, in1 *widget.Select, in2 *widget.Entry, btn *widget.Button, dis *widget.Entry) {
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

	list.SetColumnWidth(0, 230.0)
	list.SetColumnWidth(1, 50.0)
	list.SetColumnWidth(3, 210.0)

	disco := fmt.Sprintf("Discount = %v", formDiscount)
	disLabel := widget.NewLabel(disco)
	disLabel.Refresh()
	disLabel.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 10))
	disLabel.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2.5, myWindow.Canvas().Size().Height-130))

	clearBtn := widget.NewButton("Clear", func() {
		input.ClearSelected()
		mainData = [][]string{{"Product Name", "QT", "Price PerUnit", "Total"}}
		showInvoiceDataOnList(mainData, input, in1, in2, btn, dis)
	})
	clearBtn.SetIcon(theme.DeleteIcon())
	clearBtn.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	clearBtn.Move(fyne.NewPos(myWindow.Canvas().Size().Width-myWindow.Canvas().Size().Width/5.5, myWindow.Canvas().Size().Height-130))

	confirmBtn := widget.NewButton("Confirm", func() {
		if len(input.Selected) > 0 {
			ShowInvoice()
			mainData = [][]string{{"Product Name", "QT", "Price PerUnit", "Total"}}
			input.ClearSelected()

			showInvoiceDataOnList(mainData, input, in1, in2, btn, dis)
		} else {
			dialog.NewInformation("Warning!", "Please Select Client", myWindow).Show()
		}
	})
	confirmBtn.SetIcon(theme.ConfirmIcon())
	confirmBtn.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	confirmBtn.Move(fyne.NewPos(myWindow.Canvas().Size().Width-myWindow.Canvas().Size().Width/5.5, myWindow.Canvas().Size().Height-80))

	cancelBtn := widget.NewButton("Cancel", func() {
		InvoiceDash(myApp)
	})
	cancelBtn.SetIcon(theme.CancelIcon())
	cancelBtn.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	cancelBtn.Move(fyne.NewPos(myWindow.Canvas().Size().Width-myWindow.Canvas().Size().Width/2.7, myWindow.Canvas().Size().Height-80))

	var discountAmount float64
	var totalAmount float64

	for i := 1; i < len(mainData); i++ {
		price, _ := strconv.ParseFloat(mainData[i][3], 32)
		totalAmount += price
	}
	subtotalForm := fmt.Sprintf("Total = %.0f", totalAmount-formDiscount)

	fmt.Println(discountAmount)

	totalLabel = widget.NewLabel(subtotalForm)
	totalLabel.Refresh()

	totalLabel.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	totalLabel.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2.5, myWindow.Canvas().Size().Height-100))

	myWindow.SetContent(
		container.NewWithoutLayout(input, in1, in2, btn, list, dis, disLabel, totalLabel, confirmBtn, clearBtn, cancelBtn),
	)
}
