package app

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/theme"
	_"github.com/spatocode/keeper/icons"
)

func (app *Application) content() fyne.CanvasObject {
	return widget.NewVBox(
		widget.NewTabContainer(
			widget.NewTabItem("Current", app.buildCurrentTab()),
			widget.NewTabItem("Encrypted", app.buildEncryptedTab()),
		),
		widget.NewGroup(app.currentFile,
			fyne.NewContainerWithLayout(layout.NewGridLayout(3),
				widget.NewButtonWithIcon("Encrypt", theme.ConfirmIcon(), func() {}),
				widget.NewButtonWithIcon("Property", theme.ContentCutIcon(), func() {}),
				widget.NewButtonWithIcon("Decrypt", theme.CancelIcon(), func() {}),
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

func (app *Application) buildEncryptedTab() fyne.Widget{
	return widget.NewLabelWithStyle(app.currentFile, fyne.TextAlignCenter, fyne.TextStyle{Bold:true})
}
