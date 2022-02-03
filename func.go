package main

import "fmt"

func dataConveter(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

func isEmty(item, value string) (string, error) {
	companyInfo, err := GetCompanyInfo("1")
	if err != nil {
		fmt.Println(err)
	}
	if item == "" {
		item = dataConveter(companyInfo[value])
		// item = fmt.Sprintf(item, value)
	}
	return item, nil
}
