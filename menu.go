package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
)

func mainMenu() {
	menuItem1 := fyne.NewMenuItem("Light", func() {
		Light := theme.LightTheme()
		myApp.Settings().SetTheme(Light)
	})
	menuItem2 := fyne.NewMenuItem("Dark", func() {
		Dark := theme.DarkTheme()
		myApp.Settings().SetTheme(Dark)
	})
	menuItem3 := fyne.NewMenuItem("About", func() {
		dialog.NewInformation("About", "Lorem Ipsum is simply dummy text of the printing\n and typesetting industry. Lorem Ipsum has been the industry's standard\n dummy text ever since the 1500s, ", myWindow).Show()
	})
	menuItem4 := fyne.NewMenuItem("Company Info", func() {})

	newMenu1 := fyne.NewMenu("File", menuItem3, menuItem4)

	newMenu := fyne.NewMenu("Theme", menuItem1, menuItem2)

	menu := fyne.NewMainMenu(newMenu1, newMenu)
	myWindow.SetMainMenu(menu)

}
