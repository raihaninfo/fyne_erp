package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowAddGroupItem(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("Dashbord", func() {
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
		id, err := AddProductGroup(product, description)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(id)
	}

	// Set content
	win.SetContent(container.NewVBox(btnHead, btnHead1, productForm))

	// Show
	win.Show()
}
