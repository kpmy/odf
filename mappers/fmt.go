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

//New is shortcut for model.ModelFactory or wrapper
var New func(name model.LeafName) model.Leaf

//Formatter holds model of document and manages attributes and simple text mapper
//Main goal of Formatter conception is setting attributes before writing content. This allows to optimize the real attributes storage in document model
//Anyway, document model can be modified with any other tools, Formatter is my vision of this kind of tools.
//Represents Mapper in CRM pattern
type Formatter struct {
	m                 model.Model
	MimeType          xmlns.Mime
	attr              *Attr
	root              model.Node //root of document content, not model
	ready             bool
	defaultParaMapper *ParaMapper
}

//ConnectTo document model, for now it can be only newly created model, Formatter does not make existing content analise
func (f *Formatter) ConnectTo(m model.Model) {
	assert.For(m.Root().NofChild() == 0, 20, "only new documents for now")
	f.m = m
	f.attr = &Attr{}
	f.ready = false
}

//Init document model with empty content basing on MimeType
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

//WritePara writes text in new paragraph with the most latest text and paragraph attributes set
func (f *Formatter) WritePara(s string) {
	assert.For(f.ready, 20)
	f.defaultParaMapper.WritePara(s)
}

//WriteLn writes a line break
func (f *Formatter) WriteLn() {
	f.WriteString("\n")
}

//WriteString writes a text within existing paragraph or creates new paragraph if symbol \r met
func (f *Formatter) WriteString(s string) {
	assert.For(f.ready, 20)
	f.defaultParaMapper.WriteString(s)
}

//SetAttr sets any type of attributes to be used in future, only one instance of any typed attributes supported. Attributes are flushed only when real content is written. SetAttr can accept nil value for dropping all attributes
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

//RegisterFont sets the Font Face Declaration item, name later can be used in attr.TextAttributes
func (f *Formatter) RegisterFont(name, fontface string) {
	f.attr.RegisterFont(name, fontface)
}

//Sets the default attributes, that will be used by odf consumer to display non-attributed content (after SetAttr(nil))
func (f *Formatter) SetDefaults(a ...attr.Attributes) {
	f.attr.SetDefaults(a...)
}

func init() {
	New = model.LeafFactory
}
