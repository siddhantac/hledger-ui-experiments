package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	//showlist()
	showtable()
}
func showtable() {
	var data = [][]string{[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"}}
	myApp := app.New()
	myWindow := myApp.NewWindow("Table Widget")

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		})

	list := showlist()

	lab := widget.NewLabel("hello")

	// layout := layout.NewMaxLayout()
	// layout := layout.NewVBoxLayout()
	lay := layout.NewGridLayout(1)
	vb := container.NewHBox(lab, layout.NewSpacer(), list)
	container := container.New(lay, vb, table)
	myWindow.Resize(fyne.NewSize(300, 300))
	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
