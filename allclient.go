package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func ShowData(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("< Back to Dashboard", func() {
		ShowDashbod(myApp)
	})
	searchEntry := widget.NewEntry()
	searchEntry.PlaceHolder = "Client Name"
	searchFil := widget.NewFormItem("Search", searchEntry)

	searchform := widget.NewForm(searchFil)

	// button resize
	btnHead.Resize(fyne.NewSize(200, 40))
	btnHead.Move(fyne.NewPos(10, 0))

	searchform.Resize(fyne.NewSize(200, 700))
	searchform.Move(fyne.NewPos(550, 0))

	searchEntry.OnChanged = func(s string) {
		mainData := processClientSearchData(s)
		ShowDataOnList(mainData, btnHead, searchform)

	}

	data := processAllClientData()
	ShowDataOnList(data, btnHead, searchform)
	win.Show()
}

func ShowDataOnList(mainData [][]string, btnHead *widget.Button, searchform *widget.Form) {
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

	list.Resize(fyne.NewSize(800, 520))
	list.Move(fyne.NewPos(10, 50))

	list.SetColumnWidth(0, 60.0)
	list.SetColumnWidth(1, 150.0)
	list.SetColumnWidth(3, 210.0)

	list.Refresh()
	myWindow.SetContent(
		container.NewWithoutLayout(btnHead, searchform, list),
	)
}
