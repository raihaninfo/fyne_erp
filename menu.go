package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

func mainMenu() {
	menuItem1 := fyne.NewMenuItem("Light", func() {
		Light := theme.LightTheme()
		// myApp.Settings().SetTheme(Light)
		myApp.Settings().SetTheme(Light)
	})
	menuItem2 := fyne.NewMenuItem("Dark", func() {
		Dark := theme.DarkTheme()
		// myApp.Settings().SetTheme(Dark)
		myApp.Settings().SetTheme(Dark)
	})
	menuItem3:= fyne.NewMenuItem("About", func() {})

	newMenu1 := fyne.NewMenu("File", menuItem3)

	newMenu := fyne.NewMenu("Theme", menuItem1, menuItem2)

	menu := fyne.NewMainMenu(newMenu1, newMenu)
	myWindow.SetMainMenu(menu)

}
