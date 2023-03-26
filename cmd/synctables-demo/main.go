// Package main loads a demo of boarded table.
package main

import (
	"fmt"
	"synctables/pkg/borderedtable"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	w := a.NewWindow("Table scroll synchronisation demo")

	selectItems := []string{"Header and First column", "Header", "First column", "Plain"}
	table := []*borderedtable.BorderedTable{}
	for i := 0; i < len(selectItems); i++ {
		table = append(table, borderedtable.NewBoderedTable(top(i), left(i), joint(i), data(i)))
		table[i].Hide()
	}

	selectInput := widget.NewSelect(selectItems, func(s string) {
		for i := 0; i < len(selectItems); i++ {
			if selectItems[i] == s {
				table[i].Show()
			} else {
				table[i].Hide()
			}
		}
	})
	selectInput.Alignment = fyne.TextAlignCenter
	selectInput.PlaceHolder = "To see the Fyne magic just choose layout and scroll"
	content := container.NewBorder(
		selectInput,
		nil,
		nil,
		nil,
		container.NewStack(table[0], table[1], table[2], table[3]),
	)
	w.SetContent(content)

	w.Resize(fyne.NewSize(500, 250))
	w.CenterOnScreen()
	w.ShowAndRun()
}

var template string = "Cell 000, 000"
var width, height int = 150, 500

func top(mode int) *widget.Table {
	switch mode {
	case 0, 1:
		length := func() (int, int) { return 1, width }
		create := func() fyne.CanvasObject {
			return widget.NewLabel(template)
		}
		update := func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.TextStyle.Bold = true
			label.Alignment = fyne.TextAlignCenter
			label.SetText(fmt.Sprintf("Header %d", id.Col+1))
		}
		return widget.NewTable(length, create, update)
	}
	return nil
}

func left(mode int) *widget.Table {
	switch mode {
	case 0, 2:
		length := func() (int, int) { return height, 1 }
		create := func() fyne.CanvasObject {
			return widget.NewLabel(template)
		}
		update := func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.TextStyle.Bold = true
			label.Alignment = fyne.TextAlignCenter
			label.SetText(fmt.Sprintf("Row %d", id.Row+1))
		}
		return widget.NewTable(length, create, update)
	}
	return nil
}

func joint(mode int) *widget.Table {
	switch mode {
	case 0:
		length := func() (int, int) { return 1, 1 }
		create := func() fyne.CanvasObject {
			return widget.NewLabel(template)
		}
		update := func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.TextStyle.Bold = true
			label.Alignment = fyne.TextAlignCenter
			label.SetText("Row/Header")
		}
		return widget.NewTable(length, create, update)
	}
	return nil
}

func data(mode int) *widget.Table {
	length := func() (int, int) { return height, width }
	create := func() fyne.CanvasObject {
		return widget.NewLabel(template)
	}
	update := func(id widget.TableCellID, cell fyne.CanvasObject) {
		label := cell.(*widget.Label)
		label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+1))
	}
	return widget.NewTable(length, create, update)
}
