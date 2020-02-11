package main

import (
	"github.com/spatocode/keeper/app"
)

func main() {
	app := app.Load()
	app.Window().ShowAndRun()
}
