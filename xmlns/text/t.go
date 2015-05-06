package text

import (
	"odf/model"
	"odf/xmlns"
)

const (
	Paragraph model.LeafName = "text:p"
	S         model.LeafName = "text:s"
	LineBreak model.LeafName = "text:line-break"
	Tab       model.LeafName = "text:tab"
	Span      model.LeafName = "text:span"
)

const (
	C         model.AttrName = "text:c"
	StyleName model.AttrName = "text:style-name"
)

func init() {
	xmlns.Typed[C] = xmlns.INT
}
