package mappers

import (
	"odf/model"
	"odf/xmlns/office"
)

type Attr struct {
	doc  model.Model
	ds   model.Leaf //document styles
	ffd  model.Leaf //font-face decls
	as   model.Leaf //automatic styles
	ms   model.Leaf //master styles
	asc  model.Leaf //automatic styles
	ffdc model.Leaf //font-face decls
}

func (a *Attr) Init(m model.Model) {
	a.doc = m
	wr := a.doc.NewWriter()
	wr.Pos(a.doc.Root())
	a.ds = wr.WritePos(New(office.DocumentStyles))
	wr.Attr(office.Version, "1.0")
	a.ffd = wr.WritePos(New(office.FontFaceDecls))
	a.as = wr.WritePos(New(office.AutomaticStyles))
	a.ms = wr.WritePos(New(office.MasterStyles))
	a.asc = wr.WritePos(New(office.AutomaticStyles))
	a.ffdc = wr.WritePos(New(office.FontFaceDecls))
}
