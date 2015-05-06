package style

import (
	"odf/model"
	"odf/xmlns"
)

const (
	Style          model.LeafName = "style:style"
	TextProperties model.LeafName = "style:text-properties"
	FontFace       model.LeafName = "style:font-face"
)

const (
	Family   model.AttrName = "style:family"
	Name     model.AttrName = "style:name"
	FontName model.AttrName = "style:font-name"
)

const (
	FamilyText = "text"
)

func init() {
	xmlns.Typed[Family] = xmlns.ENUM
	xmlns.Enums[Family] = []string{FamilyText}
}
