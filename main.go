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

	c1 := NewBoderedTable(top(), left(), joint(), data())
	c2 := NewBoderedTable(top(), nil, nil, data())
	c3 := NewBoderedTable(nil, left(), joint(), data())
	c4 := NewBoderedTable(nil, nil, nil, data())
	c2.Hide()
	c3.Hide()
	c4.Hide()

	b1 := widget.NewButton("Header and First column", func() {
		c1.Show()
		c2.Hide()
		c3.Hide()
		c4.Hide()
	})
	b2 := widget.NewButton("Header", func() {
		c1.Hide()
		c2.Show()
		c3.Hide()
		c4.Hide()
	})
	b3 := widget.NewButton("First column", func() {
		c1.Hide()
		c2.Hide()
		c3.Show()
		c4.Hide()
	})
	b4 := widget.NewButton("Plain", func() {
		c1.Hide()
		c2.Hide()
		c3.Hide()
		c4.Show()
	})

	content := container.NewBorder(
		container.NewVBox(widget.NewLabel("To see the fyne magic just chhose layout and pull scrollbar :)"), container.NewHBox(b1, b2, b3, b4)),
		nil,
		nil,
		nil,
		container.NewStack(c1, c2, c3, c4),
	)
	w.SetContent(content)

	w.Resize(fyne.NewSize(1000, 500))
	w.CenterOnScreen()
	w.ShowAndRun()
}

var template string = "Cell 000, 000"
var width, height int = 150, 500

func top() *widget.Table {
	t := widget.NewTable(
		func() (int, int) { return 1, width - 1 },
		func() fyne.CanvasObject {
			return widget.NewLabel(template)
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.TextStyle.Bold = true
			label.Alignment = fyne.TextAlignCenter
			label.SetText(fmt.Sprintf("Column %d", id.Col+2))
		})
	return t
}

func left() *widget.Table {
	t := widget.NewTable(
		func() (int, int) { return height, 1 },
		func() fyne.CanvasObject {
			return widget.NewLabel(template)
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.TextStyle.Bold = true
			label.Alignment = fyne.TextAlignCenter
			label.SetText(fmt.Sprintf("Row %d", id.Row+1))
		})
	return t
}

func joint() *widget.Table {
	t := widget.NewTable(
		func() (int, int) { return 1, 1 },
		func() fyne.CanvasObject {
			return widget.NewLabel(template)
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.TextStyle.Bold = true
			label.Alignment = fyne.TextAlignCenter
			label.SetText("Column 1")
		})
	return t
}

func data() *widget.Table {
	t := widget.NewTable(
		func() (int, int) { return height, width - 1 },
		func() fyne.CanvasObject {
			return widget.NewLabel(template)
		},
		func(id widget.TableCellID, cell fyne.CanvasObject) {
			label := cell.(*widget.Label)
			label.SetText(fmt.Sprintf("Cell %d, %d", id.Row+1, id.Col+2))
		})
	return t
}
