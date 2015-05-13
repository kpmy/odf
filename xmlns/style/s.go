package style

import (
	"odf/model"
	"odf/xmlns"
)

const (
	Style model.LeafName = "style:style"

	TextProperties        model.LeafName = "style:text-properties"
	TableProperties       model.LeafName = "style:table-properties"
	ParagraphProperties   model.LeafName = "style:paragraph-properties"
	TableRowProperties    model.LeafName = "style:table-row-properties"
	TableColumnProperties model.LeafName = "style:table-column-properties"
	TableCellProperties   model.LeafName = "style:table-cell-properties"

	DefaultStyle model.LeafName = "style:default-style"
	FontFace     model.LeafName = "style:font-face"
)

const (
	Family                model.AttrName = "style:family"
	Name                  model.AttrName = "style:name"
	FontName              model.AttrName = "style:font-name"
	Width                 model.AttrName = "style:width"
	UseOptimalRowHeight   model.AttrName = "style:use-optimal-row-height"
	UseOptimalColumnWidth model.AttrName = "style:use-optimal-column-width"
)

const (
	FamilyText        = "text"
	FamilyParagraph   = "paragraph"
	FamilyTable       = "table"
	FamilyTableRow    = "table-row"
	FamilyTableColumn = "table-column"
	FamilyTableCell   = "table-cell"
)

func init() {
	xmlns.Typed[Family] = xmlns.ENUM
	xmlns.Enums[Family] = []string{FamilyText, FamilyParagraph, FamilyTable, FamilyTableRow, FamilyTableColumn, FamilyTableCell}
	xmlns.Typed[Width] = xmlns.MEASURE
	xmlns.Typed[UseOptimalColumnWidth] = xmlns.BOOL
	xmlns.Typed[UseOptimalRowHeight] = xmlns.BOOL
}
