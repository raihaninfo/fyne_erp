package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/mateors/msql"
)

// client add query
func AddClient(name, mobile, email, address string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "client")
	data.Set("dbtype", "sqlite3")
	data.Set("name", name)
	data.Set("mobile", mobile)
	data.Set("email", email)
	data.Set("address", address)

	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	fmt.Println("Successfully inserted", pid)
	return pid, nil
}

// product add query
func AddProduct(item_name, item_group, price, warranty, warranty_period string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "item")
	data.Set("dbtype", "sqlite3")
	data.Set("item_name", item_name)
	data.Set("item_group", item_group)
	data.Set("price", price)
	data.Set("warranty", warranty)
	data.Set("warrant_period", warranty_period)

	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	fmt.Println("Successfully inserted", pid)
	return pid, nil
}

// product group add query
func AddProductGroup(group_name, description string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "item_group")
	data.Set("dbtype", "sqlite3")
	data.Set("group_name", group_name)
	data.Set("description", description)

	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Println("Successfully inserted", pid)
	return pid, nil
}

// get client query
func GetClient() ([]map[string]interface{}, error) {
	qs := "SELECT * FROM client;"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Get all product
func GetProduct() ([]map[string]interface{}, error) {
	qs := "SELECT * FROM item;"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// get product group query
func GetProductGroup() ([]map[string]interface{}, error) {
	qs := "SELECT * FROM item_group;"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

// Get company info query
func GetCompanyInfo(id string) (map[string]interface{}, error) {
	var row = make(map[string]interface{})
	qs := `SELECT * FROM company WHERE id="%s";`
	qs = fmt.Sprintf(qs, id)
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		return nil, err
	}
	for _, sRow := range rows {
		row = sRow
	}
	return row, nil
}

// update company info
func UpdateCompany(name, address, website, email, mobile string) (bool, error) {
	name, err = isEmpty(name, "company_name")
	address, err = isEmpty(address, "address")
	website, err = isEmpty(website, "website")
	email, err = isEmpty(email, "email")
	mobile, err = isEmpty(mobile, "mobile")
	if err != nil {
		fmt.Println(err)
	}

	qs := `UPDATE company SET company_name = "%s", address = "%s", website= "%s", email="%s", mobile="%s";`
	qs = fmt.Sprintf(qs, name, address, website, email, mobile)
	row := msql.RawSQL(qs, db)
	return row, nil
}

func Search(search, table, searchBy string) []map[string]interface{} {
	qs := fmt.Sprintf("SELECT * FROM %s WHERE %s LIKE '%s%%';", table, searchBy, search)
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func GetPrice(item_name string) []map[string]interface{} {
	// SELECT price FROM item WHERE item_name="Mi Note 9 Pro";
	qs := fmt.Sprintf("SELECT price FROM item WHERE item_name='%s';", item_name)
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func GetClientAddress(name string) []map[string]interface{} {
	qs := fmt.Sprintf("SELECT address FROM client WHERE name='%s';", name)
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		fmt.Println(err)
	}
	return rows
}

func GetProductName() []string {
	tableData := []string{}
	rows, err := GetProduct()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(rows); i++ {
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["item_name"]))
	}
	return tableData
}

func GetClientName() []string {
	tableData := []string{}
	rows, err := GetClient()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < len(rows); i++ {
		tableData = append(tableData, fmt.Sprintf("%v", rows[i]["name"]))
	}
	return tableData
}

func AddInvoiceInfo(client_name, invoice_number, invoice_total, invoice_date, total_price, due, due_amount string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "invoice")
	data.Set("dbtype", "sqlite3")
	data.Set("client_name", client_name)
	data.Set("invoice_number", invoice_number)
	data.Set("invoice_total", invoice_total)
	data.Set("invoice_date", invoice_date)
	data.Set("total_price", total_price)
	data.Set("due", due)
	data.Set("due_amount", due_amount)

	pid, err := msql.InsertIntoAnyTable(data, db)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	fmt.Println("Successfully inserted", pid)
	return pid, nil
}
