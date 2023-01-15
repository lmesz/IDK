package main

import (
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Dir"), 40, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right"), 0, 1, false)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
