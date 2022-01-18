package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("dashBord")
	w.Resize(fyne.NewSize(500, 500))

	headLable := widget.NewLabel("Welcome to MiniERP")
	dateLabel := widget.NewLabel("Data")

	textL := widget.NewLabel("Welcomte to our aplication")

	btn1 := widget.NewButton("Add", func() {
		textL.Text = "Add"
		textL.Refresh()
	})
	btn2 := widget.NewButton("Edit", func() {
		textL.Text = "Edit"
		textL.Refresh()
	})
	btn3 := widget.NewButton("Cancel", func() {
		textL.Text = "cancel"
		textL.Refresh()
	})

	cl := container.NewGridWithColumns(1, container.NewGridWithColumns(2, headLable, dateLabel))

	title := container.NewVBox(btn1, btn2, btn3)
	work := container.NewVBox(textL)

	lll := container.NewHSplit(title, work)
	lll.Horizontal = true

	w.SetContent(
		container.NewVBox(cl, container.NewVBox(lll)),
	)

	w.ShowAndRun()
}
