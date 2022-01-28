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
func AddProduct(item_name, item_group_id, price, warranty, warranty_period string) (int64, error) {
	data := make(url.Values)
	data.Set("table", "item")
	data.Set("dbtype", "sqlite3")
	data.Set("item_name", item_name)
	data.Set("item_group_it", item_group_id)
	data.Set("price", price)
	data.Set("warranty", warranty)
	data.Set("warranty_period", warranty_period)

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
func GetClient() []map[string]interface{} {
	qs := "SELECT * FROM client;"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		panic(err)
	}
	return rows
}

// get product group query
func GetProductGroup() []map[string]interface{} {
	qs := "SELECT * FROM item_group;"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		panic(err)
	}
	return rows
}

func GetCompanyInfo() []map[string]interface{} {
	qs := "SELECT * FROM company;"
	rows, err := msql.GetAllRowsByQuery(qs, db)
	if err != nil {
		panic(err)
	}
	return rows
}
