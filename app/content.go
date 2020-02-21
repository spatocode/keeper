package app

import (
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
	for _, file := range app.encryptedFiles {
		if app.currentFile == file {
			handleAction(decryptor.Decrypt, app)
		}
	}
}

func (app *Application) handleEncryption() {
	if app.currentFile != "No file selected" {
		handleAction(encryptor.Encrypt, app)
	}
}

func (app *Application) handleFileProperty() {
}

func handleAction(action func(string, string) error, app *Application) {
	password := widget.NewPasswordEntry()
	dialog.ShowCustomConfirm("Enter password", "Done", "Cancel", password, func(done bool) {
		if done && password.Text != "" {
			err := action(app.currentFile, password.Text)
			if err != nil {
				err = errors.New("An error occured while encrypting file.")
				dialog.ShowError(err, app.window)
			}
		}
	}, app.window)
}
