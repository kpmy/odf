package fo

import (
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns"
)

const (
	FontSize     model.AttrName = "fo:font-size"
	TextAlign    model.AttrName = "fo:text-align"
	BreakBefore  model.AttrName = "fo:break-before"
	Color        model.AttrName = "fo:color"
	FontWeight   model.AttrName = "fo:font-weight"
	FontStyle    model.AttrName = "fo:font-style"
	BorderRight  model.AttrName = "fo:border-right"
	BorderLeft   model.AttrName = "fo:border-left"
	BorderTop    model.AttrName = "fo:border-top"
	BorderBottom model.AttrName = "fo:border-bottom"
)

const (
	Page   = "page"
	Bold   = "bold"
	Italic = "italic"
)

const (
	Start   = "start"
	End     = "end"
	Left    = "left"
	Right   = "right"
	Center  = "center"
	Justify = "justify"
)

const (
	None       = "none"
	Solid      = "solid"
	Dotted     = "dotted"
	Dash       = "dash"
	LongDash   = "long-dash"
	DotDash    = "dot-dash"
	DotDotDash = "dot-dot-dash"
	Wave       = "wave"
)

func init() {
	xmlns.Typed[FontSize] = xmlns.INT
	xmlns.Typed[TextAlign] = xmlns.ENUM
	xmlns.Enums[TextAlign] = []string{Start, End, Left, Right, Center, Justify}
	xmlns.Typed[BreakBefore] = xmlns.ENUM
	xmlns.Enums[BreakBefore] = []string{Page}
	xmlns.Typed[Color] = xmlns.COLOR
}
