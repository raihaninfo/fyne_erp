package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowAddGroupItem(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("Dashboard", func() {
		ShowDashbod(myApp)
	})
	btnHead1 := widget.NewButton("Add Product", func() {
		ShowProductAdd(myApp)
	})

	// product name entry
	productGroupNameEntry := widget.NewEntry()
	productGroupNameEntry.PlaceHolder = "Enter Product Group"

	// product price entry
	productDec := widget.NewEntry()
	productDec.PlaceHolder = "Write short description"

	productGroup := widget.NewFormItem("Product Group", productGroupNameEntry)
	groupDec := widget.NewFormItem("Short Description", productDec)

	// product Form
	productForm := widget.NewForm(productGroup, groupDec)

	// Submit Button Text
	productForm.SubmitText = "Add"

	// Submit logic
	productForm.OnSubmit = func() {
		product := productGroupNameEntry.Text
		description := productDec.Text
		if len(product) == 0 || len(description) == 0 {
			dialog.NewInformation("Warning!", "Please Fill All Information", myWindow).Show()
		} else {
			AddProductGroup(product, description)
			productGroupNameEntry.Text = ""
			productDec.Text = ""
			productForm.Refresh()
		}

	}

	win.SetContent(container.NewVBox(btnHead, btnHead1, productForm))
	win.Show()
}
