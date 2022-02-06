package main

import (
	"fmt"

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

	btn3 := widget.NewButton("Show Invoice", func() {
		ShowInvoice()
		fmt.Println("Invoice Created .....")
	})

	win.SetContent(
		container.NewVBox(btn1, btn2, btn3),
	)
	win.Show()
}
