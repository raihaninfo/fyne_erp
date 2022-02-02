package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func processAllProductData() [][]string {
	tableData := [][]string{
		{"Id", "Product Name", "Group", "Price", "warranty", "warranty Period"},
	}
	rows, err := GetProduct()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(rows); i++ {
		var tempRow []string

		warranty := dataConveter(rows[i]["warranty"])
		var isWarranty string
		if warranty == "0" {
			isWarranty = "NO"
		} else {
			isWarranty = "YES"
		}

		tempRow = append(tempRow, dataConveter(rows[i]["id"]))
		tempRow = append(tempRow, dataConveter(rows[i]["item_name"]))
		tempRow = append(tempRow, dataConveter(rows[i]["item_group"]))
		tempRow = append(tempRow, dataConveter(rows[i]["price"]))
		tempRow = append(tempRow, isWarranty)
		tempRow = append(tempRow, dataConveter(rows[i]["warrant_period"]))

		tableData = append(tableData, tempRow)
	}

	return tableData
}

func ShowAllProduct(a fyne.App) {
	win := myWindow
	btnHead := widget.NewButton("< Back to Dashbord", func() {
		ShowDashbod(myApp)
	})

	// button resize
	btnHead.Resize(fyne.NewSize(200, 40))
	btnHead.Move(fyne.NewPos(10, 0))

	data := processAllProductData()
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

	list.Resize(fyne.NewSize(800, 520))
	list.Move(fyne.NewPos(10, 50))

	list.SetColumnWidth(0, 60.0)
	list.SetColumnWidth(1, 200.0)
	list.SetColumnWidth(3, 100.0)

	myWindow.SetContent(
		container.NewWithoutLayout(btnHead, list),
	)
	win.Show()
}
