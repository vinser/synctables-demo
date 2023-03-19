package main // import "fyne.io/fyne/v2/container"

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type TableHeader struct{}

// Declare conformity with Widget interface.
var _ fyne.Widget = (*BorderedTable)(nil)

type BorderedTable struct {
	widget.BaseWidget

	top   *widget.Table
	left  *widget.Table
	joint *widget.Table
	data  *widget.Table
}

func NewBoderedTable(top, left, joint, data *widget.Table) *BorderedTable {
	if data == nil {
		fyne.LogError("Missing data", nil)
		return nil
	}
	if top != nil && left != nil && joint == nil {
		fyne.LogError("Missing joint cell", nil)
		return nil
	}
	if top == nil || left == nil {
		joint = nil
	}

	if top != nil {
		top.OnScrolled = func(p fyne.Position) {
			data.SyncHPos(p)
			top.BaseWidget.Refresh()
		}
	}
	if left != nil {
		left.OnScrolled = func(p fyne.Position) {
			data.SyncVPos(p)
			left.BaseWidget.Refresh()
		}
	}
	if top != nil || left != nil {
		data.OnScrolled = func(p fyne.Position) {
			if top != nil {
				top.SyncHPos(p)
			}
			if left != nil {
				left.SyncVPos(p)
			}
			data.BaseWidget.Refresh()
		}
	}
	t := &BorderedTable{top: top, left: left, joint: joint, data: data}
	t.ExtendBaseWidget(t)

	return t
}

func (t *BorderedTable) CreateRenderer() fyne.WidgetRenderer {
	nv := func(t *widget.Table, o []fyne.CanvasObject) (fyne.CanvasObject, []fyne.CanvasObject) {
		if t != nil {
			o = append(o, t)
			return t, o
		}
		return nil, o
	}
	o := []fyne.CanvasObject{}
	top, o := nv(t.top, o)
	left, o := nv(t.left, o)
	joint, o := nv(t.joint, o)
	data, o := nv(t.data, o)

	return &boderedTableRenderer{
		borderedTable: t,
		container:     container.New(NewBorderedTableLayout(top, left, joint, data), o...),
	}
}

var _ fyne.WidgetRenderer = (*boderedTableRenderer)(nil)

type boderedTableRenderer struct {
	borderedTable *BorderedTable
	container     *fyne.Container
}

func (*boderedTableRenderer) Destroy() {
}

func (r *boderedTableRenderer) Layout(s fyne.Size) {
	r.container.Resize(s)
}

func (r *boderedTableRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

func (r *boderedTableRenderer) Objects() []fyne.CanvasObject {
	return []fyne.CanvasObject{r.container}
}

func (r *boderedTableRenderer) Refresh() {
	r.container.Refresh()
}
