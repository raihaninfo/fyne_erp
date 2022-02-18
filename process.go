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
	rows := Search(searchData)
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
