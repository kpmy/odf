package mappers

import (
	"odf/model"
	"odf/xmlns/table"
	"ypk/assert"
)

type Table struct {
	Rows                int
	Columns             int
	Root                model.Leaf
	rowCache, colsCache []model.Leaf
	cellCache           [][]model.Leaf
}

type TableMapper struct {
	List map[string]*Table
	fm   *Formatter
}

func (t *TableMapper) Ready() bool {
	return t.fm != nil && t.fm.ready
}

func (t *TableMapper) rider() model.Writer {
	return t.fm.rider
}

func (t *TableMapper) newWriter(old ...model.Writer) model.Writer {
	return t.fm.m.NewWriter(old...)
}

func (t *TableMapper) ConnectTo(fm *Formatter) {
	t.fm = fm
	t.List = make(map[string]*Table)
}

func (t *TableMapper) Write(name string, rows, cols int) {
	assert.For(t.Ready(), 20)
	assert.For(name != "" && t.List[name] == nil, 21)
	t.fm.writeAttr()
	this := &Table{Rows: rows, Columns: cols}
	t.List[name] = this
	wr := t.newWriter(t.rider())
	this.Root = wr.WritePos(New(table.Table))
	wr.Attr(table.Name, name)
	for i := 0; i < this.Columns; i++ {
		col := New(table.TableColumn)
		this.colsCache = append(this.colsCache, col)
		this.cellCache = append(this.cellCache, make([]model.Leaf, 0))
		wr.Write(col)
	}
	for i := 0; i < this.Rows; i++ {
		rwr := t.newWriter(wr)
		row := rwr.WritePos(New(table.TableRow))
		this.rowCache = append(this.rowCache, row)
		for j := 0; j < this.Columns; j++ {
			cell := New(table.TableCell)
			this.cellCache[j] = append(this.cellCache[j], cell)
			rwr.Write(cell)
		}
	}
}

func (t *TableMapper) WriteRows(this *Table, rows int) {
	assert.For(t.Ready(), 20)
	t.fm.writeAttr()
	wr := t.newWriter()
	for i := 0; i < rows; i++ {
		wr.Pos(this.Root)
		row := wr.WritePos(New(table.TableRow))
		this.rowCache = append(this.rowCache, row)
		for j := 0; j < this.Columns; j++ {
			cell := New(table.TableCell)
			this.cellCache[j] = append(this.cellCache[j], cell)
			wr.Write(cell)
		}
		this.Rows++
	}
}

func (t *TableMapper) WriteColumns(this *Table, cols int) {
	assert.For(t.Ready(), 20)
	t.fm.writeAttr()
	wr := t.newWriter()
	var last model.Leaf
	if this.Columns > 0 {
		last = this.colsCache[this.Columns-1]
	}
	for i := 0; i < cols; i++ {
		wr.Pos(this.Root)
		col := wr.WritePos(New(table.TableColumn), last)
		this.colsCache = append(this.colsCache, col)
		this.cellCache = append(this.cellCache, make([]model.Leaf, 0))
		this.Columns++
		for j := 0; j < this.Rows; j++ {
			t.WriteCells(this, j, 1)
		}
	}
}

func (t *TableMapper) WriteCells(this *Table, _row int, cells int) {
	assert.For(t.Ready(), 20)
	t.fm.writeAttr()
	wr := t.newWriter()
	row := this.rowCache[_row]
	wr.Pos(row)
	for i := 0; i < cells; i++ {
		cell := New(table.TableCell)
		this.cellCache[i] = append(this.cellCache[i], cell)
		wr.Write(cell)
	}
}
