package fo

import (
	"odf/model"
	"odf/xmlns"
)

const (
	FontSize    model.AttrName = "fo:font-size"
	TextAlign   model.AttrName = "fo:text-align"
	BreakBefore model.AttrName = "fo:break-before"
)

const (
	Page = "page"
)

const (
	Start   = "start"
	End     = "end"
	Left    = "left"
	Right   = "right"
	Center  = "center"
	Justify = "justify"
)

func init() {
	xmlns.Typed[FontSize] = xmlns.INT
	xmlns.Typed[TextAlign] = xmlns.ENUM
	xmlns.Enums[TextAlign] = []string{Start, End, Left, Right, Center, Justify}
	xmlns.Typed[BreakBefore] = xmlns.ENUM
	xmlns.Enums[BreakBefore] = []string{Page}
}
