package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func CompanyInfoUpdate(a fyne.App) {
	win := a.NewWindow("Company Info")
	win.Resize(fyne.NewSize(500, 400))
	win.SetFixedSize(true)

	companyNameEntry := widget.NewEntry()
	companyAddressEntry := widget.NewMultiLineEntry()
	companyWebsiteEntry := widget.NewEntry()
	companyEmailEntry := widget.NewEntry()
	companyMobileEntry := widget.NewEntry()

	companyName := widget.NewFormItem("Name", companyNameEntry)
	companyAddress := widget.NewFormItem("Address", companyAddressEntry)
	companyWebsite := widget.NewFormItem("Website", companyWebsiteEntry)
	companyEmail := widget.NewFormItem("Email", companyEmailEntry)
	companyMobile := widget.NewFormItem("Mobile", companyMobileEntry)

	updateCompanyInfo := widget.NewForm(companyName, companyAddress, companyWebsite, companyEmail, companyMobile)

	updateCompanyInfo.SubmitText = "Update"
	updateCompanyInfo.OnSubmit = func() {
		name := companyNameEntry.Text
		address := companyAddressEntry.Text
		website := companyWebsiteEntry.Text
		email := companyEmailEntry.Text
		mobile := companyMobileEntry.Text
		UpdateCompany(name, address, website, email, mobile)
		// fmt.Println(name, address, website, email, mobile)
		win.Close()

	}

	win.SetContent(updateCompanyInfo)
	win.Show()
}
