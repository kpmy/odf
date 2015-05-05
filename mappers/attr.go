package mappers

import (
	"odf/model"
	"odf/xmlns/office"
)

type Attr struct {
	doc model.Model
	ds  model.Leaf //document styles
}

func (a *Attr) Init(m model.Model) {
	a.doc = m
	wr := a.doc.NewWriter()
	wr.Pos(a.doc.Root())
	a.ds = wr.WritePos(New(office.DocumentStyles))
	wr.Attr(office.Version, "1.0")
}
