package svg

import (
	"github.com/kpmy/odf/model"
	"github.com/kpmy/odf/xmlns"
)

const (
	FontFamily model.AttrName = "svg:font-family"
	Width      model.AttrName = "svg:width"
	Height     model.AttrName = "svg:height"
)

func init() {
	xmlns.Typed[Height] = xmlns.MEASURE
	xmlns.Typed[Width] = xmlns.MEASURE
}
