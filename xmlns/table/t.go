package table

import (
	"odf/model"
	"odf/xmlns"
)

const (
	Table       model.LeafName = "table:table"
	TableColumn model.LeafName = "table:table-column"
	TableRow    model.LeafName = "table:table-row"
	TableCell   model.LeafName = "table:table-cell"
)

const (
	Name                 model.AttrName = "table:name"
	NumberColumnsSpanned                = "table:number-columns-spanned"
	NumberRowsSpanned                   = "table:number-rows-spanned"
)

func init() {
	xmlns.Typed[NumberColumnsSpanned] = xmlns.INT
	xmlns.Typed[NumberRowsSpanned] = xmlns.INT
}
