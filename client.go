package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func ShowAnother(a fyne.App) {
	// time.Sleep(time.Second * 5)

	win := myWindow
	nameEntry := widget.NewEntry()
	mobileEntry := widget.NewEntry()
	emailEntry := widget.NewEntry()
	addressEntry := widget.NewMultiLineEntry()

	name1 := widget.NewFormItem("Name", nameEntry)
	mobile1 := widget.NewFormItem("Mobile", mobileEntry)
	email1 := widget.NewFormItem("Email", emailEntry)
	address1 := widget.NewFormItem("Address", addressEntry)
	clientForm := widget.NewForm(name1, mobile1, email1, address1)

	clientForm.SubmitText = "Save"

	clientForm.OnSubmit = func() {
		name := nameEntry.Text
		mobile := mobileEntry.Text
		email := emailEntry.Text
		address := addressEntry.Text

		fmt.Println(name, mobile, email, address)

	}
	win.SetContent(clientForm)

	win.Show()

}
