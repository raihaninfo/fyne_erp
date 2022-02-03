package main

import "fmt"

func dataConveter(input interface{}) string {
	return fmt.Sprintf("%v", input)
}

// is company table emty, replease with old data
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
