package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func buildRegisterList(data []string) *widget.List {
	return widget.NewList(
		func() int {
			return len(data)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i])
		})
}

func buildRegisterTable(data [][]string) *widget.Table {
	return widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Register")
		},
		func(cellID widget.TableCellID, obj fyne.CanvasObject) {
			obj.(*widget.Label).SetText(data[cellID.Row][cellID.Col])
		},
	)
}
