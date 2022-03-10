package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func ShowClient(a fyne.App) {

	win := myWindow
	btnHead := widget.NewButton("Dashboard", func() {
		ShowDashbod(myApp)
	})

	nameEntry := widget.NewEntry()
	nameEntry.PlaceHolder = "Enter Client Name"
	mobileEntry := widget.NewEntry()
	mobileEntry.PlaceHolder = "Enter Client Mobile"
	emailEntry := widget.NewEntry()
	emailEntry.PlaceHolder = "Enter Client Email Address"
	addressEntry := widget.NewMultiLineEntry()
	addressEntry.PlaceHolder = "Enter Client Address"

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
		if len(nameEntry.Text) == 0 || len(mobileEntry.Text) == 0 || len(emailEntry.Text) == 0 || len(addressEntry.Text) == 0 {
			dialog.NewInformation("Warning!", "Please Fill All Information", myWindow).Show()
		} else {
			id, err := AddClient(name, mobile, email, address)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(id)
			nameEntry.Text = ""
			mobileEntry.Text = ""
			emailEntry.Text = ""
			addressEntry.Text = ""
			clientForm.Refresh()
		}

	}

	win.SetContent(container.NewVBox(btnHead, clientForm))
	win.Show()

}
