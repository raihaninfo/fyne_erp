package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var (
	myApp    fyne.App    = app.New()
	myWindow fyne.Window = myApp.NewWindow("Mini ERP")
	db       *sql.DB
	err      error
)

func init() {
	dbcon()
}

func main() {
	myWindow.Resize(fyne.NewSize(800, 600))
	mainMenu()

	emailEntry := widget.NewEntry()
	emailEntry.PlaceHolder = "Enter your email"

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Enter your password"

	email := widget.NewFormItem("Email", emailEntry)
	password := widget.NewFormItem("Password", passwordEntry)

	clientForm := widget.NewForm(email, password)

	clientForm.SubmitText = "Login"
	messageLabel := widget.NewLabel("")

	clientForm.OnSubmit = func() {
		// appEmail := "admin@abc.com"
		// appPass := "12345678"
		appEmail := ""
		appPass := ""
		if emailEntry.Text == appEmail && passwordEntry.Text == appPass {
			ShowDashbord(myApp)
		} else {
			dialog.NewInformation("Invalid Email or password", "Email or Passwor invalid", myWindow).Show()
		}

	}

	myWindow.SetContent(
		container.NewVBox(clientForm, messageLabel),
	)
	myWindow.ShowAndRun()
}
