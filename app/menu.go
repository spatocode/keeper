package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
	fdialog "github.com/sqweek/dialog"
)

func (app *Application) buildMenu() {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", handleOpen),
			fyne.NewMenuItem("Remove current", func() {}),
		),
		fyne.NewMenu("Edit",
			fyne.NewMenuItem("Copy", func() {}),
			fyne.NewMenuItem("Paste", func() {}),
			fyne.NewMenuItem("Preference", func() {}),
		),
		fyne.NewMenu("Help",
			fyne.NewMenuItem("About", handleAbout),
			fyne.NewMenuItem("Check for updates", func() {}),
		)
	)
}

func (app *Application) handleOpen() {
	file, err := fdialog.File().Load()
	if err != nil {
		return
	}
	app.currentFile = file
}

func handleAbout() {
	dialog.ShowInformation("About", `
		Keeper is a program that keeps your files protected from a third party by 
		encrypting its content with strong encryption scheme.

		Copyright (c) 2020 Ekene Izukanne
	`)
}