package mappers

import (
	"odf/model"
	"odf/xmlns"
	"ypk/assert"
)

var New func(name model.LeafName) model.Leaf

type Formatter struct {
	m        model.Model
	inner    model.Writer
	MimeType xmlns.Mime
	attr     *Attr
}

func (f *Formatter) ConnectTo(m model.Model) {
	assert.For(m.Root().NofChild() == 0, 20, "only new documents for now")
	f.m = m
	f.inner = f.m.NewWriter()
	f.attr = &Attr{}
}

func (f *Formatter) Init() {
	assert.For(f.MimeType != "", 20)
	wr := f.m.NewWriter()
	wr.Pos(f.m.Root())
	f.attr.Init(f.m)
}

func init() {
	New = model.LeafFactory
}
