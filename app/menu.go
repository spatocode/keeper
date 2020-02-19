package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	fdialog "github.com/sqweek/dialog"
)

var aboutInfo = `
Keeper is a program that keeps your files protected 
from a third party access by encrypting its content 
with strong encryption scheme.

Copyright (c) 2020 Ekene Izukanne
`

func (app *Application) buildMenu() *fyne.MainMenu {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", app.handleOpen),
			fyne.NewMenuItem("Remove current", func() {}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Copy", func() {}),
			fyne.NewMenuItem("Paste", func() {}),
			fyne.NewMenuItem("Preference", func() {}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("Check for updates", func() {}),
			fyne.NewMenuItem("About", app.handleAbout),
		),
	)
}

func (app *Application) handleOpen() {
	file, err := fdialog.File().Load()
	if err != nil {
		return
	}
	app.currentFile = file
}

func (app *Application) handleAbout() {
	dialog.ShowInformation("About", aboutInfo, app.window)
}
