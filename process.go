package main

import "fmt"

func processAllClientData() [][]string {
	tableData := [][]string{
		{"Id", "Name", "Mobile", "Email", "Address"},
	}
	rows, err := GetClient()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(rows); i++ {
		var tempRow []string
		tempRow = append(tempRow, dataConveter(rows[i]["id"]))
		tempRow = append(tempRow, dataConveter(rows[i]["name"]))
		tempRow = append(tempRow, dataConveter(rows[i]["mobile"]))
		tempRow = append(tempRow, dataConveter(rows[i]["email"]))
		tempRow = append(tempRow, dataConveter(rows[i]["address"]))

		tableData = append(tableData, tempRow)
	}

	return tableData
}

func processClientSearchData(searchData string) [][]string {
	tableData := [][]string{
		{"Id", "Name", "Mobile", "Email", "Address"},
	}
	rows := Search(searchData, "client", "name")
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(rows); i++ {
		var tempRow []string
		tempRow = append(tempRow, dataConveter(rows[i]["id"]))
		tempRow = append(tempRow, dataConveter(rows[i]["name"]))
		tempRow = append(tempRow, dataConveter(rows[i]["mobile"]))
		tempRow = append(tempRow, dataConveter(rows[i]["email"]))
		tempRow = append(tempRow, dataConveter(rows[i]["address"]))

		tableData = append(tableData, tempRow)
	}

	return tableData
}

func processProductSearchData(searchData string) [][]string {
	tableData := [][]string{
		{"Id", "Product Name", "Group", "Price", "warranty", "warranty Period"},
	}
	rows := Search(searchData, "item", "item_name")
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

func processAllGroupData() []string {
	tableData := []string{}
	rows, err := GetProductGroup()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(rows); i++ {
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["group_name"]))
	}
	return tableData
}
