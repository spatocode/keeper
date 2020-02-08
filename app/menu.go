package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/dialog"
)

func (app *Application) buildMenu() {
	return fyne.NewMainMenu(
		fyne.NewMenu("File",
			fyne.NewMenuItem("Open", func() {}),
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

func handleAbout() {
	dialog.ShowInformation("About", `
		Keeper is a program that keeps your files protected from a third party 
		by encrypting its content with strong encryption scheme.

		Copyright 2020
	`)
}