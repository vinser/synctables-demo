// Package main loads a demo of boarded table.
package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Table scroll synchronisation demo")

	buttonName := []string{"Header and First column", "Header", "First column", "Plain"}
	table := []*BorderedTable{}
	for i := 0; i < len(buttonName); i++ {
		table = append(table, NewBoderedTable(top(i), left(i), joint(i), data(i)))
		if i != 0 {
			table[i].Hide()
		}
	}
	funcs := []func(){}
	funcs = append(funcs, func() {
		table[0].Show()
		table[1].Hide()
		table[2].Hide()
		table[3].Hide()
	})
	funcs = append(funcs, func() {
		table[0].Hide()
		table[1].Show()
		table[2].Hide()
		table[3].Hide()
	})
	funcs = append(funcs, func() {
		table[0].Hide()
		table[1].Hide()
		table[2].Show()
		table[3].Hide()
	})
	funcs = append(funcs, func() {
		table[0].Hide()
		table[1].Hide()
		table[2].Hide()
		table[3].Show()
	})
	button := []*widget.Button{}
	for i := 0; i < len(buttonName); i++ {
		button = append(button,
			widget.NewButton(buttonName[i], funcs[i]))
	}

	content := container.NewBorder(
		container.NewVBox(
			widget.NewLabel("To see the Fyne magic just choose layout and pull scrollbar :)"),
			container.NewHBox(button[0], button[1], button[2], button[3]),
		),
		nil,
		nil,
		nil,
		container.NewStack(table[0], table[1], table[2], table[3]),
	)
	w.SetContent(content)

	w.Resize(fyne.NewSize(1000, 500))
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
