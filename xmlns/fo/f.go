package fo

import (
	"odf/model"
	"odf/xmlns"
)

const (
	FontSize model.AttrName = "fo:font-size"
)

func init() {
	xmlns.Typed[FontSize] = xmlns.INT
}
