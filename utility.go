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
