package mappers

import (
	"odf/model"
	"odf/xmlns"
	"odf/xmlns/office"
	"ypk/assert"
)

var New func(name model.LeafName) model.Leaf

type Formatter struct {
	m        model.Model
	rider    model.Writer
	MimeType xmlns.Mime
	attr     *Attr
	text     model.Node
	ready    bool
}

func (f *Formatter) ConnectTo(m model.Model) {
	assert.For(m.Root().NofChild() == 0, 20, "only new documents for now")
	f.m = m
	f.rider = f.m.NewWriter()
	f.attr = &Attr{}
	f.ready = false
}

func (f *Formatter) Init() {
	assert.For(f.MimeType != "", 20)
	wr := f.m.NewWriter()
	wr.Pos(f.m.Root())
	wr.Write(New(office.DocumentMeta))
	f.attr.Init(f.m)
	wr.WritePos(New(office.DocumentContent))
	wr.Attr(office.Version, "1.0")
	wr.Write(f.attr.ffdc)
	wr.Write(f.attr.asc)
	wr.WritePos(New(office.Body))
	f.text = wr.WritePos(New(office.Text)).(model.Node)
	f.rider.Pos(f.text)
	f.ready = true
}

func init() {
	New = model.LeafFactory
}
