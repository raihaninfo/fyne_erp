package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func SearchClient(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("< Back to Dashbord", func() {
		ShowDashbod(myApp)
	})
	searchEntry := widget.NewEntry()
	searchEntry.PlaceHolder = "Client Name"
	SearchFil := widget.NewFormItem("", searchEntry)

	Searchform := widget.NewForm(SearchFil)
	// var mainData [][]string
	// var list *widget.Table

	searchEntry.OnChanged = func(s string) {
		searchData := searchEntry.Text
		mainData := processClientSearchData(searchData)
		fmt.Println(mainData)
	}

	data := processAllClientData()
	// fmt.Println(mainData)
	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)
	

	list.Resize(fyne.NewSize(800, 520))
	list.Move(fyne.NewPos(10, 50))

	list.SetColumnWidth(0, 60.0)
	list.SetColumnWidth(1, 150.0)
	list.SetColumnWidth(3, 210.0)
	// button resize
	btnHead.Resize(fyne.NewSize(200, 40))
	btnHead.Move(fyne.NewPos(10, 0))

	Searchform.Resize(fyne.NewSize(200, 400))
	Searchform.Move(fyne.NewPos(300, 0))
	// list.Refresh()

	myWindow.SetContent(
		container.NewWithoutLayout(btnHead, Searchform, list),
	)

	win.Show()
}
