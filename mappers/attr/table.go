package attr

import (
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns/fo"
	"github.com/kpmy/odf/xmlns/style"
	"github.com/kpmy/odf/xmlns/table"
)

//TableAttributes is a Table Family style fluent builder
type TableAttributes struct {
	named
	easy
}

func (t *TableAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*TableAttributes)
	if ok {
		ok = t.equal(&a.easy)
	}
	return
}

func (t *TableAttributes) Fit() model.LeafName { return table.Table }

func (t *TableAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyTable)
	wr.WritePos(New(style.TableProperties))
	t.apply(wr)
}

//BorderModel of table (table.BorderModelCollapsing, table.BorderModelSeparating
func (t *TableAttributes) BorderModel(m string) *TableAttributes {
	t.put(table.BorderModel, m, nil)
	return t
}

//AlignLeft table on page
func (t *TableAttributes) AlignLeft() *TableAttributes {
	t.put(table.Align, table.AlignLeft, nil)
	return t
}

//AlignRight table on page
func (t *TableAttributes) AlignRight() *TableAttributes {
	t.put(table.Align, table.AlignRight, nil)
	return t
}

//AlignCenter table on page
func (t *TableAttributes) AlignCenter() *TableAttributes {
	t.put(table.Align, table.AlignCenter, nil)
	return t
}

//Width of whole table
func (t *TableAttributes) Width(w float64) *TableAttributes {
	t.put(style.Width, w, nil)
	return t
}

//TableRowAttributes represents Table Row Family style fluent builder
type TableRowAttributes struct {
	named
	easy
}

func (t *TableRowAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*TableRowAttributes)
	if ok {
		ok = t.equal(&a.easy)
	}
	return
}

func (t *TableRowAttributes) Fit() model.LeafName { return table.TableRow }

func (t *TableRowAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyTableRow)
	wr.WritePos(New(style.TableRowProperties))
	t.apply(wr)
}

//UseOptimalRowHeight allows to auto-height rows when displayed
func (t *TableRowAttributes) UseOptimalRowHeight() *TableRowAttributes {
	t.put(style.UseOptimalRowHeight, true, triggerBoolAttr(style.UseOptimalRowHeight))
	return t
}

type TableColumnAttributes struct {
	named
	easy
}

//TableColumnAttributes represents Table Column Family style fluent builder
func (t *TableColumnAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*TableColumnAttributes)
	if ok {
		ok = t.equal(&a.easy)
	}
	return
}

func (t *TableColumnAttributes) Fit() model.LeafName { return table.TableColumn }

func (t *TableColumnAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyTableColumn)
	wr.WritePos(New(style.TableColumnProperties))
	t.apply(wr)
}

//UseOptimalColumnWidth allows to auto-width columns when displayed
func (t *TableColumnAttributes) UseOptimalColumnWidth() *TableColumnAttributes {
	t.put(style.UseOptimalColumnWidth, true, triggerBoolAttr(style.UseOptimalColumnWidth))
	return t
}

//TableCellAttributes represents Table Cell Family style fluent builder
type TableCellAttributes struct {
	named
	easy
}

func (t *TableCellAttributes) Equal(_a Attributes) (ok bool) {
	a, ok := _a.(*TableCellAttributes)
	if ok {
		ok = t.equal(&a.easy)
	}
	return
}

func (t *TableCellAttributes) Fit() model.LeafName { return table.TableCell }

func (t *TableCellAttributes) Write(wr model.Writer) {
	wr.Attr(style.Family, style.FamilyTableCell)
	wr.WritePos(New(style.TableCellProperties))
	t.apply(wr)
}

//Border sets attributes for all borders (left, right, top, bottom)
func (t *TableCellAttributes) Border(b Border) *TableCellAttributes {
	t.put(fo.BorderRight, b.String(), nil)
	t.put(fo.BorderLeft, b.String(), nil)
	t.put(fo.BorderTop, b.String(), nil)
	t.put(fo.BorderBottom, b.String(), nil)
	return t
}
