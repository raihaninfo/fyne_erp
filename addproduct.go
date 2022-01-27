package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func processAllGroupData() []string {
	tableData := []string{}
	rows := GetProductGroup()
	for i := 0; i < len(rows); i++ {
		// var tempRow []string
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["group_name"]))
	}
	return tableData
}

func ShowProductAdd(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("Dashbord", func() {
		ShowDashbord(myApp)
	})
	btnHead1 := widget.NewButton("Add Item GRUP", func() {
		ShowAddGroupItem(myApp)
	})

	// product name entry
	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Enter Product name"

	// product price entry
	priceEntry := widget.NewEntry()
	priceEntry.PlaceHolder = "Enter Price"

	allGroupData := processAllGroupData()
	groupEntry := widget.NewSelect(allGroupData, func(s string) {

	})

	// is warranty check
	warrantyEntry := widget.NewCheck("Check if there is a warranty", func(b bool) {})

	// warranty period
	warrantyPeriod := widget.NewSelect([]string{"10 Days", "30 Days", "6 Month", "1 Year"}, func(s string) {

	})

	productName := widget.NewFormItem("Product Name", nameEntry)
	productPrice := widget.NewFormItem("Product Price", priceEntry)
	productWarranty := widget.NewFormItem("Product Warranty", warrantyEntry)
	productWarrantyPeriod := widget.NewFormItem("Warranty Period", warrantyPeriod)
	group := widget.NewFormItem("Group", groupEntry)

	// product Form
	productForm := widget.NewForm(productName, productPrice, group, productWarranty, productWarrantyPeriod)

	// Submit Button Text
	productForm.SubmitText = "Add"

	// Submit logic
	productForm.OnSubmit = func() {
		name := nameEntry.Text
		price := priceEntry.Text
		warranty := warrantyEntry.Checked
		warrantyPeriod := warrantyPeriod.Selected

		isWarranty := "0"
		if warranty {
			isWarranty = "1"
		}

		AddProduct(name, "1", price, isWarranty, warrantyPeriod)

		fmt.Println(name, price, warranty, warrantyPeriod)
	}

	// Set content
	win.SetContent(container.NewVBox(btnHead, btnHead1, productForm))

	// Show
	win.Show()
}
