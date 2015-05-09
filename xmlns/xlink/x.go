package xlink

import (
	"odf/model"
	"odf/xmlns"
)

const (
	Href    model.AttrName = "xlink:href"
	Type    model.AttrName = "xlink:type"
	Show    model.AttrName = "xlink:show"
	Actuate model.AttrName = "xlink:actuate"
)

const Simple = "simple"
const Embed = "embed"
const OnLoad = "onload"

func init() {
	xmlns.Typed[Type] = xmlns.ENUM
	xmlns.Typed[Show] = xmlns.ENUM
	xmlns.Typed[Actuate] = xmlns.ENUM
	xmlns.Enums[Type] = []string{Simple}
	xmlns.Enums[Show] = []string{Embed}
	xmlns.Enums[Actuate] = []string{OnLoad}
}
