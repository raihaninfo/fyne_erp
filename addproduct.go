package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func AddProduct(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("Dashbord", func() {
		ShowDashbord(myApp)
	})

	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Enter Product name"

	priceEntry := widget.NewEntry()
	priceEntry.PlaceHolder = "Enter Price"

	warentyEntry := widget.NewCheck("yes or not", func(b bool) {})

	warrantyPeriod := widget.NewSelect([]string{"10 Days", "30 Days", "6 Month", "1 Year"}, func(s string) {
		fmt.Println(s)
	})

	name1 := widget.NewFormItem("Product Name", nameEntry)
	mobile1 := widget.NewFormItem("Product Price", priceEntry)
	email1 := widget.NewFormItem("Product Warranty", warentyEntry)
	address1 := widget.NewFormItem("Warranty Period", warrantyPeriod)
	clientForm := widget.NewForm(name1, mobile1, email1, address1)

	clientForm.SubmitText = "Add"

	clientForm.OnSubmit = func() {
		name := nameEntry.Text
		mobile := priceEntry.Text
		email := warentyEntry.Text

		fmt.Println(name, mobile, email)

	}
	win.SetContent(container.NewVBox(btnHead, clientForm))

	win.Show()
}
