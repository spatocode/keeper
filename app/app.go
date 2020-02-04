package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

type Application struct {
	window			fyne.Window
	menu			*fyne.MainMenu
	encryptedFiles	[]string
	currentFile		string
}

func Load() *Application {
	a := app.NewWithID("Keeper-OSS")
	win := a.NewWindow("Keeper")

	appl := &App{window: win}
	appl.menu = app.buildMenu()
	appl.currentFile = "No file selected"

	win.SetMaster(appl.content())
	win.SetMainMenu(mainMenu)

	win.SetContent()
}