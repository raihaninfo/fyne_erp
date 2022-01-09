package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Mini ERP - Add new Client")
	w.Resize(fyne.NewSize(500, 500))

	nameEntry := widget.NewEntry()
	mobileEntry := widget.NewEntry()
	emailEntry := widget.NewEntry()
	addressEntry := widget.NewMultiLineEntry()

	name := widget.NewFormItem("Name", nameEntry)
	mobile := widget.NewFormItem("Mobile", mobileEntry)
	email := widget.NewFormItem("Email", emailEntry)
	address := widget.NewFormItem("Address", addressEntry)
	clientForm := widget.NewForm(name, mobile, email, address)
	
	clientForm.SubmitText = "Save"

	clientForm.OnSubmit = func() {

	}
	

	w.SetContent(
		container.NewGridWithColumns(1,
			clientForm,
		),
	)

	w.ShowAndRun()
}
