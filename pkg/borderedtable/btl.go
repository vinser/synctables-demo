package borderedtable

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

// Declare conformity with Layout interface
var _ fyne.Layout = (*boderedTableLayout)(nil)

// Layout for bordered table widget
type boderedTableLayout struct {
	top, left, joint, data fyne.CanvasObject
}

// NewBorderedTableLayout creates a new borderedTableLayout instance with top, left, joint
// and data objects set.
func NewBorderedTableLayout(top, left, joint, data fyne.CanvasObject) fyne.Layout {
	return &boderedTableLayout{top, left, joint, data}
}

// Layout is called to pack all child objects into a specified size.
// For borderTableLayout this arranges the top header, left column, joint cell table widgets at
// the sides and table widget with data is maximised in the rest of the space.
func (l *boderedTableLayout) Layout(o []fyne.CanvasObject, size fyne.Size) {
	jointSize := fyne.Size{Width: 0, Height: 0}
	switch {
	case l.joint != nil:
		jointSize = fyne.NewSize(l.joint.MinSize().Width+theme.Padding(), l.joint.MinSize().Height+theme.Padding())
	case l.top != nil && l.left == nil:
		jointSize = fyne.NewSize(0, l.top.MinSize().Height+theme.Padding())
	case l.top == nil && l.left != nil:
		jointSize = fyne.NewSize(l.left.MinSize().Width+theme.Padding(), 0)
	}
	if l.joint != nil && l.joint.Visible() {
		l.joint.Resize(fyne.NewSize(l.joint.MinSize().Width, l.joint.MinSize().Height))
		l.joint.Move(fyne.NewPos(0, 0))
	}
	if l.top != nil && l.top.Visible() {
		l.top.Resize(fyne.NewSize(size.Width-jointSize.Width, l.top.MinSize().Height))
		l.top.Move(fyne.NewPos(jointSize.Width, 0))
	}
	if l.left != nil && l.left.Visible() {
		l.left.Resize(fyne.NewSize(l.left.MinSize().Width, size.Height-jointSize.Height))
		l.left.Move(fyne.NewPos(0, jointSize.Height))
	}
	if l.data != nil && l.data.Visible() {
		middleSize := fyne.NewSize(size.Width-jointSize.Width, size.Height-jointSize.Height)
		middlePos := fyne.NewPos(jointSize.Width, jointSize.Height)
		l.data.Resize(middleSize)
		l.data.Move(middlePos)
	}
}

// MinSize finds the smallest size that satisfies all the child objects.
// For borderTableLayout this is determined by the MinSize height of the hader and
// plus the MinSize width of the left column, plus MinSize width of the joint cell
// and table data, plus any padding needed.
// This is then added to the union of the MinSize for any remaining content.
func (l *boderedTableLayout) MinSize(o []fyne.CanvasObject) fyne.Size {
	minSize := fyne.Size{Width: 0, Height: 0}
	if l.data != nil && l.data.Visible() {
		minSize = l.data.MinSize()
	}
	if l.joint != nil && l.joint.Visible() {
		minHeight := fyne.Max(minSize.Height, l.joint.MinSize().Height)
		minWidth := fyne.Max(minSize.Width, l.joint.MinSize().Width)
		minSize = fyne.NewSize(minWidth+theme.Padding(), minHeight+theme.Padding())
	}
	if l.top != nil && l.top.Visible() {
		minWidth := fyne.Max(minSize.Width, l.top.MinSize().Width)
		minSize = fyne.NewSize(minWidth, minSize.Height+l.top.MinSize().Height+theme.Padding())
	}
	if l.left != nil && l.left.Visible() {
		minHeight := fyne.Max(minSize.Height, l.left.MinSize().Height)
		minSize = fyne.NewSize(minSize.Width+l.left.MinSize().Width+theme.Padding(), minHeight)
	}

	return minSize
}
