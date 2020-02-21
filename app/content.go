package app

import (
	"os"
	"errors"

	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/dialog"
	"github.com/spatocode/keeper/encryptor"
	"github.com/spatocode/keeper/decryptor"
)

func (app *Application) content() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewTabContainer(
			widget.NewTabItem("Current", app.buildCurrentTab()),
			widget.NewTabItem("Encrypted", app.buildEncryptedTab()),
		),
		widget.NewGroup(app.currentFile,
			fyne.NewContainerWithLayout(layout.NewGridLayout(3),
				widget.NewButtonWithIcon("Encrypt", theme.ConfirmIcon(), app.handleEncryption),
				widget.NewButtonWithIcon("Property", theme.InfoIcon(), app.handleFileProperty),
				widget.NewButtonWithIcon("Decrypt", theme.CancelIcon(), app.handleDecryption),
			),
		),
	)
}

func (app *Application) buildCurrentTab() fyne.Widget {
	var text string
	if app.currentFile == "No file selected" {
		text = "Please select a file for encryption"
	} else {
		text = app.currentFile
	}
	return widget.NewVBox(
		layout.NewSpacer(),
		widget.NewLabelWithStyle(text, fyne.TextAlignCenter, fyne.TextStyle{Bold:true}),
		layout.NewSpacer(),
	)
}

func (app *Application) buildEncryptedTab() fyne.Widget {
	if app.encryptedFiles == nil {
		return widget.NewLabelWithStyle("No files encrypted yet", fyne.TextAlignCenter, fyne.TextStyle{Bold:true})
	}
	return nil // TODO: Return list of encrypted files with details
}

func (app *Application) handleDecryption() {
	if app.currentFile == "No file selected" {
		return
	}
	for _, file := range app.encryptedFiles {
		if app.currentFile == file {
			handleAction(decryptor.Decrypt, app)
		}
	}
}

func (app *Application) handleEncryption() {
	if app.currentFile == "No file selected" {
		return
	}
	handleAction(encryptor.Encrypt, app)
}

func (app *Application) handleFileProperty() {
	if app.currentFile == "No file selected" {
		//return
	}
	info, err := os.Stat(app.currentFile)
	if err != nil {
		handleError("Cannot check file property.", app)
	}

	content := widget.NewVBox(
		widget.NewHBox(
			widget.NewLabel("Name: "),
			widget.NewLabel(info.Name()),
		),
		widget.NewHBox(
			widget.NewLabel("Size: "),
			widget.NewLabel(string(info.Size())),
		),
		widget.NewHBox(
			widget.NewLabel("Last modified: "),
			widget.NewLabel(info.ModTime().String()),
		),
	)

	dialog.ShowCustom("Property", " Ok ", content, app.window)
}

func handleAction(action func(string, string) error, app *Application) {
	password := widget.NewPasswordEntry()
	dialog.ShowCustomConfirm("Enter password", "Done", "Cancel", password, func(done bool) {
		if done && password.Text != "" {
			err := action(app.currentFile, password.Text)
			if err != nil {
				handleError("An error occured while encrypting file.", app)
			}
		}
	}, app.window)
}

func handleError(msg string, app *Application) {
	err := errors.New(msg)
	dialog.ShowError(err, app.window)
}