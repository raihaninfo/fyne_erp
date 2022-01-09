package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	a := app.New()

	w := a.NewWindow("Mini ERP - Add new Client")
	w.Resize(fyne.NewSize(500, 500))

	nameEntry := widget.NewEntry()
	mobileEntry := widget.NewEntry()
	emailEntry := widget.NewEntry()
	addressEntry := widget.NewMultiLineEntry()

	name1 := widget.NewFormItem("Name", nameEntry)
	mobile1 := widget.NewFormItem("Mobile", mobileEntry)
	email1 := widget.NewFormItem("Email", emailEntry)
	address1 := widget.NewFormItem("Address", addressEntry)
	clientForm := widget.NewForm(name1, mobile1, email1, address1)

	clientForm.SubmitText = "Login"

	clientForm.OnSubmit = func() {
		name := nameEntry.Text
		mobile := mobileEntry.Text
		email := emailEntry.Text
		address := addressEntry.Text

		fmt.Println(name, mobile, email, address)

	}

	w.SetContent(
		container.NewGridWithColumns(1,
			clientForm,
		),
	)

	w.ShowAndRun()
}
