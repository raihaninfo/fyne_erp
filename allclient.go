package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func processData() [][]string {
	tableData := [][]string{
		{"Id", "Name", "Mobile", "Email", "Address"},
	}
	rows := GetClient()
	for i := 0; i < len(rows); i++ {
		var tempRow []string
		tempRow = append(tempRow, rows[i]["id"].(string))
		tempRow = append(tempRow, rows[i]["name"].(string))
		tempRow = append(tempRow, rows[i]["mobile"].(string))
		tempRow = append(tempRow, rows[i]["email"].(string))
		tempRow = append(tempRow, rows[i]["address"].(string))

		tableData = append(tableData, tempRow)
	}

	return tableData
}

func ShowData(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("Dashbord", func() {
		ShowDashbord(myApp)
	})
	data := processData()
	list := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	list.SetColumnWidth(0, 60.0)
	list.SetColumnWidth(1, 150.0)
	list.SetColumnWidth(3, 210.0)

	myWindow.SetContent(
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(1,
				btnHead,
			),
			container.NewGridWithColumns(1,
				list),
		),
	)
	win.Show()
}
