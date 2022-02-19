package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func CreateInvoice(a fyne.App) {
	win := myWindow
	input := widget.NewSelectEntry([]string{"Raihan", "Rezaul"})
	input.Resize(fyne.NewSize(myWindow.Canvas().Size().Width, 40))
	input.Move(fyne.NewPos(0, 0))

	in1 := widget.NewSelect([]string{"Mango", "Apple"}, func(s string) {

	})

	in1.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in1.Move(fyne.NewPos(0, 50))

	in2 := widget.NewEntry()
	in2.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/2.5, 40))
	in2.Move(fyne.NewPos(myWindow.Canvas().Size().Width/2.5+5, 50))

	btn := widget.NewButton("Add", func() {
		a := in1.Selected
		b := in2.Text
		fmt.Println(a, b)

	})
	btn.Resize(fyne.NewSize(myWindow.Canvas().Size().Width/5.5, 40))
	btn.Move(fyne.NewPos(myWindow.Canvas().Size().Width-myWindow.Canvas().Size().Width/5.5, 50))

	btn.Icon = theme.ContentAddIcon()

	var mainData = [][]string{[]string{"Product", "QT", "Price"},
		[]string{"Mi Note 9 Pro", "1", "12000"}}

	list := widget.NewTable(
		func() (int, int) {
			return len(mainData), len(mainData[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(mainData[i.Row][i.Col])
		},
	)

	list.Resize(fyne.NewSize(myWindow.Canvas().Size().Width, myWindow.Canvas().Size().Height-250))
	list.Move(fyne.NewPos(0, 100))

	// list.SetColumnWidth(0, 60.0)
	// list.SetColumnWidth(1, 150.0)
	// list.SetColumnWidth(3, 210.0)

	// widget.NewEntry().Validate()

	win.SetContent(
		container.NewWithoutLayout(input, in1, in2, btn, list),
	)
	myWindow.Show()
}
