package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var data = []string{"a", "string", "list"}

func main() {
	accounts, err := listAccounts()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// reg, err := register()
	reg, err := registerLines()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	myApp := app.New()
	myWindow := myApp.NewWindow("hledger ui")
	// myWindow.Resize(fyne.NewSize(500, 550))

	// upper := widget.NewLabel("Upper box")
	// lower := buildRegisterTable(reg)
	// table.Resize(fyne.NewSize(500, 200))
	list := buildRegisterList(accounts)
	dateFilter := widget.NewLabel("Date Filters")
	list2 := buildRegisterList(reg)

	// container := container.NewVBox(lower, upper)
	// container := container.New(layout.NewMaxLayout(), lower)

	lay := layout.NewGridLayout(1)
	topRow := layout.NewGridLayout(2)
	vb := container.New(topRow, list, dateFilter)
	container := container.New(lay, vb, layout.NewSpacer(), list2)

	myWindow.Resize(fyne.NewSize(1000, 700))
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}

/*
func main() {

	// data = append(data, accounts...)

	myApp := app.New()
	myWindow := myApp.NewWindow("hledger ui")
	myWindow.Resize(fyne.NewSize(100, 100))

	list := widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})

	myWindow.SetContent(list)
	myWindow.ShowAndRun()
}
*/
