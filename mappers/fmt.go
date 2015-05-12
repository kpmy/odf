package mappers

import (
	"github.com/kpmy/ypk/assert"
	"github.com/kpmy/ypk/halt"
	"odf/mappers/attr"
	"odf/model"
	"odf/xmlns"
	"odf/xmlns/office"
	"reflect"
)

var New func(name model.LeafName) model.Leaf

type Formatter struct {
	m                 model.Model
	MimeType          xmlns.Mime
	attr              *Attr
	root              model.Node //root of document content, not model
	ready             bool
	defaultParaMapper *ParaMapper
}

func (f *Formatter) ConnectTo(m model.Model) {
	assert.For(m.Root().NofChild() == 0, 20, "only new documents for now")
	f.m = m
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
	switch f.MimeType {
	case xmlns.MimeText:
		f.root = wr.WritePos(New(office.Text)).(model.Node)
	case xmlns.MimeSpreadsheet:
		f.root = wr.WritePos(New(office.Spreadsheet)).(model.Node)
	default:
		halt.As(100, f.MimeType)
	}
	f.defaultParaMapper = new(ParaMapper)
	f.defaultParaMapper.ConnectTo(f)
	f.ready = true
}

func (f *Formatter) WritePara(s string) {
	assert.For(f.ready, 20)
	f.defaultParaMapper.WritePara(s)
}

func (f *Formatter) WriteLn() {
	f.WriteString("\n")
}

func (f *Formatter) WriteString(s string) {
	assert.For(f.ready, 20)
	f.defaultParaMapper.WriteString(s)
}

func (f *Formatter) SetAttr(a attr.Attributes) *Formatter {
	assert.For(f.ready, 20)
	if a != nil {
		n := reflect.TypeOf(a).String()
		if old := f.attr.OldAttr(a); old != nil {
			f.attr.current[n] = old
		} else {
			c := f.attr.current[n]
			if (c == nil) || !c.Equal(a) {
				f.attr.stored = false
				f.attr.current[n] = a
			}
		}
	} else {
		f.attr.reset()
	}
	return f
}

func (f *Formatter) RegisterFont(name, fontface string) {
	f.attr.RegisterFont(name, fontface)
}

func (f *Formatter) SetDefaults(a ...attr.Attributes) {
	f.attr.SetDefaults(a...)
}

func init() {
	New = model.LeafFactory
}
