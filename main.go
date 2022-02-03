package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var (
	myApp    fyne.App    = app.New()
	myWindow fyne.Window = myApp.NewWindow("Mini ERP v0.0.1")
	db       *sql.DB
	err      error
)

// Database connection.
func init() {
	dbcon()
}

func main() {
	myWindow.Resize(fyne.NewSize(800, 600))
	

	emailEntry := widget.NewEntry()
	emailEntry.PlaceHolder = "Enter your email"

	passwordEntry := widget.NewPasswordEntry()
	passwordEntry.PlaceHolder = "Enter your password"

	email := widget.NewFormItem("Email", emailEntry)
	password := widget.NewFormItem("Password", passwordEntry)

	loginForm := widget.NewForm(email, password)

	loginForm.SubmitText = "Login"

	loginForm.OnSubmit = func() {
		// appEmail := "admin@abc.com"
		// appPass := "12345678"
		appEmail := ""
		appPass := ""
		if emailEntry.Text == appEmail && passwordEntry.Text == appPass {
			ShowDashbod(myApp)
		} else {
			dialog.NewInformation("Invalid Email or password", "Email or Passwor invalid", myWindow).Show()
		}

	}

	loginForm.OnCancel = func() {
		myWindow.Close()
	}

	myWindow.SetContent(
		loginForm,
	)
	myWindow.ShowAndRun()
}
