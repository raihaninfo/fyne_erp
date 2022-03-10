package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowProductAdd(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("Dashboard", func() {
		ShowDashbod(myApp)
	})

	// product name entry
	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Enter Product name"

	// product price entry
	priceEntry := widget.NewEntry()
	priceEntry.PlaceHolder = "Enter Price"

	// Product group data
	allGroupData := processAllGroupData()

	// group dropdown
	groupEntry := widget.NewSelect(allGroupData, func(s string) {

	})

	// is warranty check
	warrantyEntry := widget.NewCheck("Check if there is a warranty", func(b bool) {})

	// warranty period
	warrantyPeriodEntry := widget.NewSelect([]string{"10 Days", "30 Days", "3 Month", "6 Month", "1 Year", "2 Years", "Life Time"}, func(s string) {

	})

	productName := widget.NewFormItem("Product Name", nameEntry)
	productPrice := widget.NewFormItem("Product Price", priceEntry)
	productWarranty := widget.NewFormItem("Product Warranty", warrantyEntry)
	productWarrantyPeriod := widget.NewFormItem("Warranty Period", warrantyPeriodEntry)
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
		warrantyPeriod := warrantyPeriodEntry.Selected
		groupValue := groupEntry.Selected

		isWarranty := "0"
		if warranty {
			isWarranty = "1"
		}

		if len(name) == 0 || len(price) == 0 || groupValue == "" {
			dialog.NewInformation("Warning!", "Please Fill All Information", myWindow).Show()
		} else {
			AddProduct(name, groupValue, price, isWarranty, warrantyPeriod)
			nameEntry.Text = ""
			priceEntry.Text = ""
			groupEntry.ClearSelected()
			warrantyPeriodEntry.ClearSelected()

			warrantyEntry.Checked = false
			productForm.Refresh()
		}

	}

	productForm.OnCancel = func() {
		ShowDashbod(myApp)
	}
	productForm.CancelText = "Back"

	win.SetContent(container.NewVBox(btnHead, productForm))
	win.Show()
}
