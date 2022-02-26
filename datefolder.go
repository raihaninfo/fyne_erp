package main

import (
	"fmt"
	"os"
	"time"
)

func CreateFolderUseingDate() (string, string) {

	day := time.Now().Day()
	month := time.Now().Month().String()
	year := time.Now().Year()

	folder := fmt.Sprintf("%v-%v-%v", day, month, year)

	fl := fmt.Sprintf("pdf/%s", folder)
	os.Mkdir(fl, 0755)

	return folder, fl
}
