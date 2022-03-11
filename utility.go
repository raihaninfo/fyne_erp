package main

import (
	"fmt"
	"os"
	"time"
)

func CreateFolderUseingDate() (string, string) {
	month := time.Now().Month().String()
	year := time.Now().Year()
	folder := fmt.Sprintf("%v-%v", month, year)
	fl := fmt.Sprintf("pdf/%s", folder)
	os.Mkdir(fl, 0755)

	return folder, fl
}

func dataConveter(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

// is company table empty, replease with old data
func isEmpty(item, value string) (string, error) {
	companyInfo, err := GetCompanyInfo("1")
	if err != nil {
		fmt.Println(err)
	}
	if item == "" {
		item = dataConveter(companyInfo[value])
	}
	return item, nil
}
