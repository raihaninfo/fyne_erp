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
	myWindow fyne.Window = myApp.NewWindow("Mini ERP v1.0")
	db       *sql.DB
	err      error
)

// Database connection.
func init() {
	dbConnection()
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

			app.New().SendNotification(fyne.NewNotification("congratulations", "You have logind"))
			ShowDashbod(myApp)
		} else {
			app.New().SendNotification(fyne.NewNotification("Invalid Email Or Password", "Email or Password invalid"))
			dialog.NewInformation("Invalid Email or password", "Email or Password invalid", myWindow).Show()
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
