package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
)

// Application contains information about the Keeper app
type Application struct {
	window			fyne.Window
	menu			*fyne.MainMenu
	encryptedFiles	[]string
	currentFile		string
}

// Load initialize the Application
func Load() *Application {
	app := app.NewWithID("Keeper-OSS")
	win := app.NewWindow("Keeper")
	win.Resize(fyne.NewSize(800, 500))
	win.SetFixedSize(true)

	a := &Application{window: win}
	a.menu = a.buildMenu()
	a.currentFile = "No file selected"

	win.SetMaster()
	win.SetMainMenu(a.menu)

	win.SetContent(a.content())
	return a
}
