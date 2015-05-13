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
	BorderModel                         = "table:border-model"
	StyleName                           = "table:style-name"
	Align                               = "table:align"
)

const (
	BorderModelCollapsing = "collapsing"
	BorderModelSeparating = "separating"
	AlignLeft             = "left"
	AlignRight            = "right"
	AlignCenter           = "center"
)

func init() {
	xmlns.Typed[NumberColumnsSpanned] = xmlns.INT
	xmlns.Typed[NumberRowsSpanned] = xmlns.INT
	xmlns.Typed[BorderModel] = xmlns.ENUM
	xmlns.Enums[BorderModel] = []string{BorderModelCollapsing, BorderModelSeparating}
	xmlns.Typed[Align] = xmlns.ENUM
	xmlns.Enums[Align] = []string{AlignCenter, AlignLeft, AlignRight}
}
